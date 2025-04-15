package message_queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCommitsOffsets(t *testing.T) {
	// Arrange
	testTopic := "test_topic"
	mockConsumer := new(MockConsumer)
	mockProcessor := new(MockMessageProcessor)
	queue := NewSimulatedMessageQueue(mockConsumer, mockProcessor)
	msg := &Message{
		Key:            []byte("key"),
		Value:          []byte("value"),
		TopicPartition: TopicPartition{Topic: testTopic, Partition: 0, Offset: 1},
	}

	mockConsumer.On("ReadMessage").Return(msg, nil).Once().Run(func(args mock.Arguments) {
		queue.SignalRead()
	})
	mockProcessor.On("ProcessMessage", msg, mock.AnythingOfType("chan<- message_queue.Result")).Return(nil).Once().Run(func(args mock.Arguments) {
		args[1].(chan<- Result) <- Result{
			Message: msg,
			Error:   nil,
		}
	})
	expectedCommitOffset := msg.TopicPartition.Offset + 1
	tps := []TopicPartition{{Topic: testTopic, Partition: 0, Offset: expectedCommitOffset}}
	mockConsumer.On("CommitOffsets", tps).Return(nil).Once().Run(func(args mock.Arguments) {
		queue.SignalCommit()
	})
	mockConsumer.On("ReadMessage").Return(nil, errors.New("closed")).Run(func(args mock.Arguments) {
		queue.HangReads()
	})

	// Act
	queue.Start()

	queue.WaitForRead()
	queue.Tick()
	queue.WaitForCommit()

	queue.Stop()

	// Assert
	queue.AssertExpectations(t)
}
