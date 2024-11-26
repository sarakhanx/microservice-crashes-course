package ports

type HelloService interface {
	SayHello() string
}

type WorldService interface {
	SayWorld() string
}

type CalculatorService interface {
	Add(a, b int) int
	Subtract(a, b int) int
}
