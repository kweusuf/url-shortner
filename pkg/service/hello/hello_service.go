package service

import (
	"context"

	client "github.com/kweusuf/url-shortner/pkg/client/hello"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

type HelloService interface {
	HelloFromAppService(ctx context.Context) (interface{}, error)
}

type helloService struct {
	client client.HelloClient
}

func MakeHelloService(client client.HelloClient) HelloService {
	return &helloService{
		client: client,
	}
}

// HelloFromApp implements AppService.
func (appService *helloService) HelloFromAppService(ctx context.Context) (interface{}, error) {
	log.Debug("In HelloFromAppService method")
	return appService.client.HelloFromAppClient(ctx)
}
