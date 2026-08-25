package streamdecoderchannels

import (
	"encoding/binary"
	"fmt"
	"sync"
)

const maxMessageSize = 16 << 20 // 16 MiB
const headerSize = 4

type StreamReceiver struct {
	chunks   chan []byte
	messages chan []byte
	accept   func([]byte)
	mu       sync.RWMutex
	closed   bool
	wg       sync.WaitGroup
}

func NewStreamReceiver(accept func([]byte)) *StreamReceiver {
	s := &StreamReceiver{
		chunks:   make(chan []byte, 100),
		messages: make(chan []byte, 100),
		accept:   accept,
	}

	s.wg.Add(2)

	go s.parseLoop()
	go s.deliveryLoop()

	return s
}

// Receive is safe to call concurrently.
// The RLock prevents a send from racing with Close.
func (s *StreamReceiver) Receive(chunk []byte) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		return fmt.Errorf("stream receiver is closed")
	}

	// Copy to avoid external callers editing chunk value
	// Make a copy to avoid external changes
	chunkCopy := append([]byte(nil), chunk...)
	s.chunks <- chunkCopy

	return nil
}

func (s *StreamReceiver) Wait() {
	s.wg.Wait()
}

func (s *StreamReceiver) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}

	s.closed = true
	close(s.chunks)
}

func (s *StreamReceiver) parseLoop() {
	defer s.wg.Done()
	defer close(s.messages)

	var buffer []byte

	for chunk := range s.chunks {
		buffer = append(buffer, chunk...)

		for {
			if len(buffer) < headerSize {
				break
			}

			messageLen := binary.BigEndian.Uint32(buffer[:headerSize])

			if messageLen > maxMessageSize {
				fmt.Println("invalid message size:", messageLen)
				return
			}

			frameLen := headerSize + int(messageLen)

			// Header arrived but message is incomplete
			if len(buffer) < frameLen {
				break
			}

			// Make a copy to avoid external changes
			message := append([]byte(nil), buffer[headerSize:frameLen]...)

			// Drop what we already read
			buffer = buffer[frameLen:]

			s.messages <- message
		}
	}
}

func (s *StreamReceiver) deliveryLoop() {
	defer s.wg.Done()

	for message := range s.messages {
		s.accept(message)
	}
}
