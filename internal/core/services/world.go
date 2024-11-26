package services

type worldService struct{}

func NewWorldService() *worldService {
    return &worldService{}
}

func (s *worldService) SayWorld() string {
    return "world"
}