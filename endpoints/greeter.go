package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"web-template/middleware"
	"web-template/service"
)

type (
	Endpoints struct {
		HealthEndpoint   endpoint.Endpoint
		GreetingEndpoint endpoint.Endpoint
	}
	HealthRequest struct {
	}
	HealthResponse struct {
		Healthy bool  `json:"healthy,omitempty"`
		Err     error `json:"err,omitempty"`
	}
	GreetingRequest struct {
		Name string `json:"name,omitempty"`
	}
	GreetingResponse struct {
		Greeting string `json:"greeting,omitempty"`
		Err      error  `json:"err,omitempty"`
	}
	Failer interface {
		Failed() error
	}
)

func MakeServiceEndpoints(s service.Service, logger log.Logger) Endpoints {
	var healthEndpoint endpoint.Endpoint
	{
		healthEndpoint = MakeGreetingEndpoints(s)
		healthEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Health"))(healthEndpoint)
	}
	var greetingEndpoint endpoint.Endpoint
	{
		greetingEndpoint = MakeGreetingEndpoints(s)
		greetingEndpoint = middleware.LoggingMiddleware(log.With(logger, "method", "Greeting"))(greetingEndpoint)
	}
	return Endpoints{
		HealthEndpoint:   healthEndpoint,
		GreetingEndpoint: greetingEndpoint,
	}
}

func MakeHealthEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		healthy := s.Health()
		return HealthResponse{Healthy: healthy}, nil
	}
}

func MakeGreetingEndpoints(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GreetingRequest)
		greeting := s.Greeting(req.Name)
		return GreetingResponse{
			Greeting: greeting,
			Err:      nil,
		}, nil
	}
}

func (r HealthResponse) Failed() error {
	return r.Err
}

func (r GreetingResponse) Failed() error {
	return r.Err
}
