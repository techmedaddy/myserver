package service

import "fmt"

// Service defines your app's core logic layer.
// In real projects, you’d often inject dependencies here (like a DB or cache).
type Service struct{}

// New returns a new instance of Service.
func New() *Service {
	return &Service{}
}

// Hello returns a simple greeting.
func (s *Service) Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
