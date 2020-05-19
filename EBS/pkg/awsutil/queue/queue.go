package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	"EBS/m/v2/pkg"
	"EBS/m/v2/pkg/awsutil"
)

var svc = sqs.New(awsutil.Session)

var _ pkg.PushPuller = &Queue{}

type Queue struct {
	url string
}

func New(url string) (queue *Queue) {
	queue = &Queue{
		url: url,
	}

	return
}

func (q *Queue) Push(message []byte) (err error) {

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(message)),
		QueueUrl:    &q.url,
	})

	return
}

func (q *Queue) PushWithDelay(message []byte, sec int64) (err error) {

	_, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(sec),
		MessageBody:  aws.String(string(message)),
		QueueUrl:     &q.url,
	})

	return
}

// PushBatch pushes a batch of messages to the queue. Please be aware that
// AWS only supports batches of 10 for the time being.
// https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html
func (q *Queue) PushBatch(entries []*sqs.SendMessageBatchRequestEntry) (err error) {
	_, err = svc.SendMessageBatch(&sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(q.url),
	})
	return
}

// Pull will poll the queue for up to a second and may, or may not return a
// token.  Once it has been verifiably processed the token must be Release()'d
// to finally remove the message from the queue.
func (q *Queue) Pull() (token pkg.MessageReleaser, err error) {

	rcv, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		VisibilityTimeout:   aws.Int64(5),
		MaxNumberOfMessages: aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(1),
		QueueUrl:            &q.url,
	})

	if err != nil {
		return
	}

	if len(rcv.Messages) == 0 {
		return
	}

	msg := rcv.Messages[0]
	token = &Token{
		url:     q.url,
		handle:  *msg.ReceiptHandle,
		message: *msg.Body,
	}

	return
}

// Token implements the com.MessageReleaser interface, see Queue.Pull()
type Token struct {
	url     string
	handle  string
	message string
}

func (t *Token) Message() []byte {
	return []byte(t.message)
}

func (t *Token) Release() (err error) {

	_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
		ReceiptHandle: &t.handle,
		QueueUrl:      &t.url,
	})

	return
}
