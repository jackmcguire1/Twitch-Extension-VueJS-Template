package s3

import (
	"EBS/m/v2/pkg/awsutil"
	"archive/zip"
	"compress/gzip"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var uploader = s3manager.NewUploader(awsutil.Session)

// Writer uploads content of unknown length (vs put object)
type Writer struct {
	open  bool
	err   error
	pw    *io.PipeWriter
	outC  chan interface{}
	write func([]byte) (int, error)
	close func() error
	inp   *s3manager.UploadInput
	out   *s3manager.UploadOutput
}

func NewWriter(bucket, key string) *Writer {

	pr, pw := io.Pipe()
	outC := make(chan interface{})

	return &Writer{
		pw:    pw,
		outC:  outC,
		write: func(p []byte) (int, error) { return pw.Write(p) },
		close: func() error { return pw.Close() },
		inp: &s3manager.UploadInput{
			Body:   pr,
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
	}
}

func NewGzipWriter(bucket, key string) (w *Writer) {

	w = NewWriter(bucket, key)
	w.SetContentEncoding("gzip")

	gzw := gzip.NewWriter(w.pw)

	w.write = func(p []byte) (int, error) {
		return gzw.Write(p)
	}

	w.close = func() error {
		gzerr := gzw.Close()
		perr := w.pw.Close()
		if gzerr == nil {
			return perr
		}
		return gzerr
	}

	return
}

func NewZipWriter(bucket, key, filename string) (w *Writer) {

	w = NewWriter(bucket, key)
	w.SetContentEncoding("zip")

	zw := zip.NewWriter(w.pw)
	zwio, err := zw.Create(filename)
	if err != nil {
		return
	}
	w.write = func(p []byte) (int, error) {
		return zwio.Write(p)
	}

	w.close = func() error {
		ziperr := zw.Close()
		perr := w.pw.Close()
		if ziperr == nil {
			return perr
		}
		return ziperr
	}

	return
}

func (w *Writer) SetContentEncoding(enc string) {
	w.inp.ContentEncoding = aws.String(enc)
}

func (w *Writer) Write(p []byte) (n int, err error) {

	if !w.open {
		w.open = true
		go upload(w.inp, w.outC)
	}

	return w.write(p)
}

func (w *Writer) Close() (err error) {

	err = w.close()

	switch x := (<-w.outC).(type) {
	case *s3manager.UploadOutput:
		w.out = x
	case error:
		w.err = x
	}

	return
}

func upload(inp *s3manager.UploadInput, outC chan interface{}) {

	out, err := uploader.Upload(inp)

	if err == nil {
		outC <- out
	} else {
		outC <- err
	}
}
