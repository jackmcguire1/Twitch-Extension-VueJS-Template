package streamer

import (
	"io"

	"EBS/m/v2/pkg"
	"EBS/m/v2/pkg/awsutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var svc = s3.New(awsutil.Session)

var _ pkg.Streamer = &Streamer{}

type Streamer struct {
	InputBucket  string
	OutputBucket string
	InputKey     string
	OutputKey    string
}

func New(inputBucket, outputBucket, inputKey, outputKey string) *Streamer {
	return &Streamer{
		InputBucket:  inputBucket,
		OutputBucket: outputBucket,
		InputKey:     inputKey,
		OutputKey:    outputKey,
	}
}

func (s *Streamer) Stream(processor func(io.Reader, io.Writer) error) (err error) {
	r, w := io.Pipe()

	done := make(chan error)
	defer close(done)
	go upload(r, s.OutputBucket, s.OutputKey, done)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.InputBucket),
		Key:    aws.String(s.InputKey),
	})
	if err != nil {
		return
	}

	err = processor(obj.Body, w)
	if err != nil {
		return
	}

	obj.Body.Close()
	w.Close()
	err = <-done
	return
}

func upload(body io.Reader, bucket, key string, done chan error) {
	uploader := s3manager.NewUploaderWithClient(svc)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	done <- err
}
