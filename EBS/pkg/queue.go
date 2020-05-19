package pkg

import "github.com/aws/aws-sdk-go/service/sqs"

type PushPuller interface {
	Push([]byte) error
	Pull() (MessageReleaser, error)
	PushBatch([]*sqs.SendMessageBatchRequestEntry) error
	PushWithDelay(message []byte, delay int64) error
}

type MessageReleaser interface {
	Message() []byte
	Release() error
}

type Worker interface {
	Work(done []byte) (todo []byte, err error)
}
