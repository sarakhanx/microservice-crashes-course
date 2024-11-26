package services

type calculatorService struct{}

// สร้างตัวอย่างของตัวคำนวณ
func NewCalculatorService() *calculatorService {
	return &calculatorService{}
}

// การบวก
func (s *calculatorService) Add(a, b int) int {
	return a + b
}

// การลบ
func (s *calculatorService) Subtract(a, b int) int {
	return a - b
}
