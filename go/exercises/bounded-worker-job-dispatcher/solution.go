package boundedworkerjobdispatcher

import (
	"context"
	"fmt"
	"sync"
)

// Job is one delivery request submitted to Dispatch.
type Job struct {
	ID          string
	Destination string
	Payload     string
}

// Result is the outcome for one input Job.
type Result struct {
	Job       Job
	Completed bool
	Err       error
}

// Sender delivers a job and must respect context cancellation.
type Sender interface {
	Send(ctx context.Context, job Job) error
}

// indexedJob keeps a job's original position in the input slice.
type indexedJob struct {
	index int
	job   Job
}

// indexedResult keeps a result's original position in the input slice.
type indexedResult struct {
	index  int
	result Result
}

// Dispatch sends jobs using exactly workers sender goroutines.
//
// It returns results in the same order as jobs. Sender failures are stored in
// individual Result values. Context cancellation returns partial results and
// ctx.Err().
func Dispatch(
	ctx context.Context,
	sender Sender,
	jobs []Job,
	workers int,
) ([]Result, error) {
	if workers <= 0 {
		return nil, fmt.Errorf("dispatch: workers must be greater than zero")
	}

	for i, job := range jobs {
		if job.ID == "" {
			return nil, fmt.Errorf("dispatch: job at index %d has an empty ID", i)
		}
	}

	// Allocate exactly one result slot per input job.
	results := make([]Result, len(jobs))

	// Store input jobs up front so unsubmitted jobs still appear in partial results.
	for i, job := range jobs {
		results[i].Job = job
	}

	// jobCh is bounded to workers, providing producer backpressure.
	jobCh := make(chan indexedJob, workers)

	// resultCh carries worker outcomes to the single collector.
	resultCh := make(chan indexedResult)

	var senderWG sync.WaitGroup

	senderWG.Add(workers)

	// Start a fixed-size worker pool rather than one goroutine per job.
	// Process jobs
	for i := 0; i < workers; i++ {
		go func() {
			defer senderWG.Done()

			// Receive and process jobs until cancellation or channel closure.
			for {
				// item holds the next indexed job.
				var item indexedJob

				// ok reports whether jobCh is still open.
				var ok bool

				// Wait for either cancellation or the next job.
				select {
				case <-ctx.Done():
					return
				case item, ok = <-jobCh:
					// Exit after the producer closes jobCh.
					if !ok {
						return
					}
				}

				// Send the job using the same cancellation context.
				err := sender.Send(ctx, item.job)

				// Keep the input index so the collector can restore ordering.
				result := indexedResult{
					index: item.index,
					result: Result{
						Job:       item.job,
						Completed: err == nil,
						Err:       err,
					},
				}

				// Send the result unless cancellation tells the worker to exit.
				select {
				case <-ctx.Done():
					return
				case resultCh <- result:
				}
			}
		}()
	}

	// Start the producer that owns and closes jobCh.
	go func() {
		// Only the producer closes its outbound jobs channel.
		defer close(jobCh)

		// Submit jobs in input order until all are sent or ctx is cancelled.
		for i, job := range jobs {
			// Block only until a worker accepts the job or cancellation occurs.
			select {
			case <-ctx.Done():
				return
			case jobCh <- indexedJob{
				index: i,
				job:   job,
			}:
			}
		}
	}()

	// Start the result-channel closer after every sender worker exits.
	go func() {
		senderWG.Wait()
		close(resultCh)
	}()

	// Drain results until all sender workers finish and close resultCh.
	for result := range resultCh {
		results[result.index] = result.result
	}

	// Return partial results plus the context error after cancellation.
	if err := ctx.Err(); err != nil {
		return results, err
	}

	// Return ordered results when dispatch completed normally.
	return results, nil
}
