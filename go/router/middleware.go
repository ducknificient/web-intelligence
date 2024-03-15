package router

type Middleware interface {
}

type DefaultMiddleware struct {
}

func NewMiddleware() (default_middleware *DefaultMiddleware) {

	return &DefaultMiddleware{}

}

func (m *DefaultMiddleware) EnableCors() {

}
