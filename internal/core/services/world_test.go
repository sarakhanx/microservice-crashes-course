package services

import "testing"

func TestWorldService_SayWorld(t *testing.T) {
    // Arrange
    service := NewWorldService()

    // Act
    result := service.SayWorld()

    // Assert
    expected := "world"
    if result != expected {
        t.Errorf("SayWorld() = %v, want %v", result, expected)
    }
}