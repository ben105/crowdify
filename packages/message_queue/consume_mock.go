package message_queue

import (
	"github.com/stretchr/testify/mock"
)

type MockConsumer struct {
	mock.Mock
}

func (m *MockConsumer) ReadMessage() (*Message, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Message), args.Error(1)
}

func (m *MockConsumer) Initialize() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConsumer) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConsumer) CommitOffsets(offsets []TopicPartition) error {
	args := m.Called(offsets)
	return args.Error(0)
}
