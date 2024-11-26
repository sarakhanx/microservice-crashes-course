package services

import "testing"

func TestHelloService_SayHello(t *testing.T) {
    // Arrange
    service := NewHelloService()

    // Act
    result := service.SayHello()

    // Assert
    expected := "hello"
    if result != expected {
        t.Errorf("SayHello() = %v, want %v", result, expected)
    }
}