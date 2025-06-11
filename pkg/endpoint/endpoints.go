package endpoint

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/kweusuf/url-shortner/pkg/service"
)

type AppEndpoints struct {
	HelloEndpoint      endpoint.Endpoint
	URLShortenEndpoint endpoint.Endpoint
	URLExpandEndpoint  endpoint.Endpoint
}

func MakeEndpoints(services service.Services) AppEndpoints {
	return AppEndpoints{
		HelloEndpoint:      makeHelloEndpoint(services.HelloService),
		URLShortenEndpoint: makeURLShortenEndpoint(services.URLService),
		URLExpandEndpoint:  makeURLExpandEndpoint(services.URLService),
	}
}
