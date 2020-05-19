package s3

import (
	"EBS/m/v2/pkg/awsutil"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var svc = s3.New(awsutil.Session)

func NewReader(bucket, key string) (io.ReadCloser, error) {

	inp := s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	out, err := svc.GetObject(&inp)
	return out.Body, err
}
