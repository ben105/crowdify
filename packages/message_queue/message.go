package message_queue

type Header struct {
	Key   string
	Value []byte
}

type TopicPartition struct {
	Topic     string
	Partition int32
	Offset    int64
}

type Message struct {
	Key            []byte
	Value          []byte
	TopicPartition TopicPartition
	Headers        []Header
}

func NewMessage(key []byte, value []byte, topicPartition TopicPartition, headers ...Header) *Message {
	return &Message{
		Key:            key,
		Value:          value,
		TopicPartition: topicPartition,
		Headers:        headers,
	}
}

type MessageProcessor interface {
	ProcessMessage(msg *Message, results chan<- Result)
}
