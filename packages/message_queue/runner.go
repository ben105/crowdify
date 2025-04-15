package message_queue

import (
	"fmt"
	"log"
	"log/slog"
	"math"
	"sync"
	"time"

	"github.com/ben105/crowdify/packages/env"
)

type Runner interface {
	Start()
	Stop()
}

type Result struct {
	Message *Message
	Error   error
}

type PartitionRunLoop struct {
	consumer  Consumer
	processor MessageProcessor
	commit    CommitManager

	// results is used to send processed messages back to the main thread
	results chan Result
	// quit is used to signal the consumer to stop
	quit chan struct{}

	retryCount int8

	wg sync.WaitGroup
}

func NewRunner(consumer Consumer, processor MessageProcessor, commit CommitManager) Runner {
	return &PartitionRunLoop{
		consumer:  consumer,
		processor: processor,
		commit:    commit,
		results:   make(chan Result),
		quit:      make(chan struct{}),
	}
}

func (pr *PartitionRunLoop) Start() {
	err := pr.consumer.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize consumer: %v", err)
	}
	slog.Info("Initialized consumer\n")

	pr.wg.Add(1)
	go func() {
		defer pr.wg.Done()
		pr.processMessages()
	}()

	pr.wg.Add(1)
	go func() {
		defer pr.wg.Done()
		pr.handleResults()
	}()

	slog.Info("Partition run loop started\n")
}

func (pr *PartitionRunLoop) Stop() {
	close(pr.quit)
	err := pr.consumer.Close()
	if err != nil {
		panic(fmt.Sprintf("Failed to close consumer: %v", err))
	}
	pr.wg.Wait()
}

func (pr *PartitionRunLoop) processMessages() {
	for {
		slog.Info("Polling for messages...\n")
		select {
		case <-pr.quit:
			slog.Info("Stopping consumer...\n")
			return
		default:
			msg, err := pr.consumer.ReadMessage()
			slog.Info("Message read\n")
			if err == nil {
				go pr.processor.ProcessMessage(msg, pr.results)
			}
		}
	}
}

func (pr *PartitionRunLoop) handleResults() {
	offsetsToCommit := make(map[TopicPartition]int64)
	defer pr.commit.Stop()
	for {
		select {
		case <-pr.quit:
			slog.Debug("Stopping result handler...\n")
			pr.commitOffsets(offsetsToCommit)
			return
		case <-pr.commit.TickerChan:
			slog.Debug("Commit ticker ticked\n")
			if len(offsetsToCommit) > 0 {
				slog.Debug(fmt.Sprintf("Committing %d offsets\n", len(offsetsToCommit)))
				pr.commitOffsets(offsetsToCommit)
				offsetsToCommit = make(map[TopicPartition]int64)
			}
		case r := <-pr.results:
			tp := TopicPartition{
				Topic:     r.Message.TopicPartition.Topic,
				Partition: r.Message.TopicPartition.Partition,
			}
			if r.Error != nil {
				slog.Error(fmt.Sprintf("Error processing message: %v\n", r.Error))
				pr.handleError(r)
			} else {
				currentOffset, exists := offsetsToCommit[tp]
				if !exists || r.Message.TopicPartition.Offset > currentOffset {
					offsetsToCommit[tp] = r.Message.TopicPartition.Offset
				}
				slog.Debug(fmt.Sprintf("Processed message: %s\n", string(r.Message.Value)))
			}
		}
	}
}

func (pr *PartitionRunLoop) handleError(r Result) {
	if r.Error == nil {
		panic("handleError called with nil error")
	}
	if pr.retryCount >= 5 {
		slog.Info(fmt.Sprintf("Max retries reached for message: %s", string(r.Message.Value)))
		sendToDlq(r)
	} else {
		pr.retryCount++
		backoff := time.Duration(math.Pow(float64(pr.retryCount), 2)) * time.Second
		slog.Debug(fmt.Sprintf("Waiting %v before retrying message: %s, attempt %d\n", backoff, string(r.Message.Value), pr.retryCount))
		time.Sleep(backoff)
		slog.Debug(fmt.Sprintf("Retrying message: %s, attempt %d\n", string(r.Message.Value), pr.retryCount))
		go pr.processor.ProcessMessage(r.Message, pr.results)
	}
}

func (pr *PartitionRunLoop) commitOffsets(offsetsToCommit map[TopicPartition]int64) {
	if len(offsetsToCommit) == 0 {
		return
	}

	tps := make([]TopicPartition, 0, len(offsetsToCommit))

	for tp, offset := range offsetsToCommit {
		tps = append(tps, TopicPartition{
			Topic:     tp.Topic,
			Partition: tp.Partition,
			Offset:    offset + 1, // +1 because committed offset is the next expected one
		})
	}

	slog.Debug(fmt.Sprintf("Committing offsets for %d partitions\n", len(tps)))

	err := pr.consumer.CommitOffsets(tps)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to commit offsets: %v\n", err))
	} else {
		slog.Debug("Successfully committed offsets\n")
	}
}

func sendToDlq(r Result) {
	p := NewProducer(env.GetBroker(), env.GetDeadLetterQueueTopic())
	dl := NewDeadLetter(p)
	dl.SendToDlq(r)
}
