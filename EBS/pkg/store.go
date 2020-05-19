package pkg

type GetPutter interface {
	Put(name string, object []byte) error
	Get(name string) (object []byte, err error)
}
