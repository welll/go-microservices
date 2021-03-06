package service

import "log"

// Service describes a service that adds things together.
type Service interface {
	Sum(a, b int) (int, error)
	Concat(a, b string) (string, error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger *log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

func (s basicService) Sum(a, b int) (v int, err error) { return a + b, nil }

func (s basicService) Concat(a, b string) (v string, err error) { return a + b, nil }
