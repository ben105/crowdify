package message_queue

import (
	"fmt"
)

type DeadLettering interface {
	SendToDlq(r Result)
}

type DeadLetter struct {
	p Producer
}

func NewDeadLetter(p Producer) *DeadLetter {
	return &DeadLetter{p: p}
}

func (dl *DeadLetter) SendToDlq(r Result) {
	headers := []Header{
		{
			Key:   "error_message",
			Value: []byte(r.Error.Error()),
		},
		{
			Key:   "original_topic",
			Value: []byte(r.Message.TopicPartition.Topic),
		},
		{
			Key:   "original_partition",
			Value: []byte(fmt.Sprintf("%d", r.Message.TopicPartition.Partition)),
		},
		{
			Key:   "original_offset",
			Value: []byte(fmt.Sprintf("%d", r.Message.TopicPartition.Offset)),
		},
	}
	if r.Message.Headers != nil {
		headers = append(headers, r.Message.Headers...)
	}
	dl.p.Produce(r.Message.Value)
}
