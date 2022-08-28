package handler

type Handler interface {
	ObjectCreated(interface{}) error
	ObjectDeleted(interface{}) error
}

// todo 通过泛型简化handler
type Handle[T any] struct {
	Handler
}

func (Handle[T]) ObjectCreated(obj interface{}) (err error) {
	return
}

func (Handle[T]) ObjectDeleted(obj interface{}) (err error) {
	return
}
