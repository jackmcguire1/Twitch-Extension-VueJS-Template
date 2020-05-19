package topic

import (
	"EBS/m/v2/pkg"
	"EBS/m/v2/pkg/awsutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

var svc = sns.New(ebs)

var _ pkg.Publisher = &Topic{}

type Topic struct {
	url string
}

func New(name string) (t *Topic) {
	t = &Topic{name}
	return
}

func (t *Topic) Publish(message []byte) (err error) {

	_, err = svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(message)),
		TopicArn: &t.url,
	})

	return
}
