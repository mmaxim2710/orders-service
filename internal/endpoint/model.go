package endpoint

type Service interface {
	HelloWorld() string
}

type Endpoint struct {
	s Service
}
