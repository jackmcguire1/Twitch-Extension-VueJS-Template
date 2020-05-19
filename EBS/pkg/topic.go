package pkg

type Publisher interface {
	Publish([]byte) error
}
