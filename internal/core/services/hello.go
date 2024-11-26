package services

type helloService struct{}

func NewHelloService() *helloService {
	return &helloService{}
}

func (s *helloService) SayHello() string {
	return "hello"
}
