package streamdecoder

import (
	"encoding/binary"
	"fmt"
	"sync"
)

const maxMessageSize = 16 << 20 // 16 MiB
const headerSize = 4

type StreamDecoder struct {
	mu     sync.Mutex
	buf    []byte
	accept func([]byte)
}

func NewStreamDecoder(accept func([]byte)) *StreamDecoder {
	return &StreamDecoder{
		accept: accept,
	}
}

func (s *StreamDecoder) Receive(chunk []byte) error {
	var messagesToDeliver [][]byte
	s.mu.Lock()

	s.buf = append(s.buf, chunk...)

	for {
		if len(s.buf) < headerSize {
			break
		}

		messageLen := binary.BigEndian.Uint32(s.buf[:headerSize])

		if messageLen > maxMessageSize {
			s.mu.Unlock()
			return fmt.Errorf("message size %d exceeds limit", messageLen)
		}

		frameLen := headerSize + int(messageLen)

		// Header arrived but message is incomplete
		if len(s.buf) < frameLen {
			break
		}

		message := append([]byte(nil), s.buf[headerSize:frameLen]...)
		messagesToDeliver = append(messagesToDeliver, message)

		s.buf = s.buf[frameLen:]
	}

	s.mu.Unlock()

	// Separated the accept function since we do not control user provider call. If consuners are slow and we keep lock we will block receivers
	for _, message := range messagesToDeliver {
		s.accept(message)
	}

	return nil
}
