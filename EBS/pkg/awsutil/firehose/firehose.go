package firehose

import (
	"EBS/m/v2/pkg"
	"EBS/m/v2/pkg/awsutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
)

var svc = firehose.New(awsutil.Session)

var _ pkg.Putter = &Firehose{}

type Firehose struct {
	name string
}

func New(name string) (firehose *Firehose) {
	firehose = &Firehose{name}
	return
}

func (f *Firehose) Put(message []byte) (err error) {
	input := &firehose.PutRecordInput{
		DeliveryStreamName: aws.String(f.name),
		Record: &firehose.Record{
			Data: message,
		},
	}

	_, err = svc.PutRecord(input)
	return
}
