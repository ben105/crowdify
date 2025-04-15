package message_queue

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type MockMessageProcessor struct {
	mock.Mock

	// If true, the ProcessMessage method will signal errors
	Fail bool
}

func (m *MockMessageProcessor) ProcessMessage(msg *Message, results chan<- Result) {
	_ = m.Called(msg, results)
	if m.Fail {
		results <- Result{
			Message: nil,
			Error:   errors.New("processing error"),
		}
	} else {
		results <- Result{
			Message: msg,
			Error:   nil,
		}
	}
}
