package bucket

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"path"
	"time"

	"EBS/m/v2/pkg/awsutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

const MaxKeys = 1000

var svc = s3.New(awsutil.Session)

type Bucket struct {
	Name string
}

// New returns a pointer to a bucket
func New(name string) *Bucket {
	return &Bucket{Name: name}
}

// Returns the object stored in the bucket under the specified key
// Docs: https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.GetObject
func (b *Bucket) Get(key string) ([]byte, error) {
	res, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(b.Name),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// Stores the object in the bucket under the specified key
// Docs: https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.PutObject
func (b *Bucket) Put(key string, body []byte) error {
	input := &s3.PutObjectInput{
		Body:   bytes.NewReader(body),
		Bucket: aws.String(b.Name),
		Key:    aws.String(key),
	}

	_, err := svc.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}

// EmitKeys returns all keys after given start key filtered by the given
// prefix
func (b *Bucket) EmitKeys(
	prefix,
	startKey string,
) (
	keys chan interface{},
) {

	keys = make(chan interface{}, MaxKeys)
	input := &s3.ListObjectsV2Input{}
	input.SetBucket(b.Name)
	if prefix != "" {
		input.SetPrefix(prefix)
	}
	if startKey != "" {
		input.SetStartAfter(startKey)
	}

	go func() {
		err := svc.ListObjectsV2Pages(
			input,
			func(res *s3.ListObjectsV2Output, lastPage bool) bool {
				for _, obj := range res.Contents {
					keys <- *obj.Key
				}
				return true
			},
		)
		if err != nil {
			keys <- err
		}
		close(keys)
	}()

	return
}

// TailKeys every interval
func (b *Bucket) TailKeys(
	marker, prefix string,
	interval time.Duration,
	keys chan interface{},
) {
	input := &s3.ListObjectsV2Input{}
	input.SetBucket(b.Name)
	input.SetPrefix(prefix)
	input.SetStartAfter(marker)

	for {
		res, err := svc.ListObjectsV2(input)
		if err != nil {
			keys <- err
			time.Sleep(time.Second)
			continue
		}

		next := ""

		for _, obj := range res.Contents {
			next = *obj.Key
			keys <- next
		}

		if next != "" {
			input.SetStartAfter(next)
		}

		time.Sleep(interval)
	}
}

// List returns a list of flat prefixes found within a given subdirectory
// or root of the bucket.
func (b *Bucket) List(
	prefix string,
) (
	prefixes []string,
	err error,
) {
	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(b.Name),
		Delimiter: aws.String("/"),
	}
	if prefix != "" {
		input.SetPrefix(prefix)
	}

	err = svc.ListObjectsV2Pages(
		input,
		func(res *s3.ListObjectsV2Output, lastPage bool) bool {
			for _, prefix := range res.CommonPrefixes {
				prefixes = append(prefixes, path.Base(*prefix.Prefix))
			}
			return true
		},
	)

	return
}

func (b *Bucket) DeleteKey(key string) (err error) {
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(b.Name),
		Key:    aws.String(key),
	})

	return
}

func (b *Bucket) KeyExists(key string) (bool, error) {
	_, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(b.Name),
		Key:    aws.String(key),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			// currently this doesn't return NoSuchKey as stated in the documentation,
			// for precaution we use both
			case "NotFound", s3.ErrCodeNoSuchKey:
				return false, nil
			default:
				return false, aerr
			}
		}

		return false, err
	}

	return true, nil
}
