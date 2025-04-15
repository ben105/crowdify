package message_queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type SimulatedMessageQueue struct {
	mockConsumer  *MockConsumer
	mockProcessor *MockMessageProcessor

	// Synchronization Channels
	readCalledChan   chan struct{} // Signal first read
	commitCalledChan chan struct{} // Signal commit
	hangChan         chan struct{} // Channel to block subsequent reads

	// Mocked Ticker channel to simulate periodic commits
	mockTickerChan chan<- time.Time

	runner *Runner
}

func NewSimulatedMessageQueue(mockConsumer *MockConsumer, mockProcessor *MockMessageProcessor) *SimulatedMessageQueue {
	// Synchronization Channels
	readCalledChan := make(chan struct{}, 1)   // Signal first read
	commitCalledChan := make(chan struct{}, 1) // Signal commit
	hangChan := make(chan struct{})            // Channel to block subsequent reads

	mockConsumer.On("Initialize").Return(nil).Once()
	mockConsumer.On("Close").Return(nil).Once().Run(func(args mock.Arguments) {
		close(hangChan)
	})

	mockTickerChan := make(chan time.Time)
	runner := NewRunner(mockConsumer, mockProcessor, *NewCommitManagerForTest(mockTickerChan))

	return &SimulatedMessageQueue{
		mockConsumer:     mockConsumer,
		mockProcessor:    mockProcessor,
		readCalledChan:   readCalledChan,
		commitCalledChan: commitCalledChan,
		hangChan:         hangChan,
		mockTickerChan:   mockTickerChan,
		runner:           &runner,
	}
}

func (s *SimulatedMessageQueue) SignalRead() {
	s.readCalledChan <- struct{}{}
}

func (s *SimulatedMessageQueue) WaitForRead() {
	select {
	case <-s.readCalledChan:
	case <-time.After(3 * time.Second):
		panic("Timed out waiting for ReadMessage call.")
	}
}

func (s *SimulatedMessageQueue) HangReads() {
	<-s.hangChan
}

func (s *SimulatedMessageQueue) SignalCommit() {
	s.commitCalledChan <- struct{}{}
}

func (s *SimulatedMessageQueue) WaitForCommit() {
	select {
	case <-s.commitCalledChan:
	case <-time.After(3 * time.Second):
		panic("Timed out waiting for CommitOffsets call.")
	}
}

func (s *SimulatedMessageQueue) Tick() {
	time.Sleep(100 * time.Millisecond)
	s.mockTickerChan <- time.Time{}
}

func (s *SimulatedMessageQueue) Start() {
	(*s.runner).Start()
}

func (s *SimulatedMessageQueue) Stop() {
	(*s.runner).Stop()
}

func (s *SimulatedMessageQueue) AssertExpectations(t *testing.T) {
	s.mockConsumer.AssertExpectations(t)
	s.mockProcessor.AssertExpectations(t)
}
