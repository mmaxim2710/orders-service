package service

func New() *Service {
	return &Service{}
}

func (s *Service) HelloWorld() string {
	return "Hello world!"
}
