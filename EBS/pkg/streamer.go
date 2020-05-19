package pkg

import (
	"io"
)

type Streamer interface {
	Stream(func(io.Reader, io.Writer) error) error
}
