package client

import (
	"context"

	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

type HelloClient interface {
	HelloFromAppClient(ctx context.Context) (interface{}, error)
}

type helloClient struct {
}

func MakeHelloClient() HelloClient {
	return &helloClient{}
}

// HelloFromAppClient implements HelloClient.
func (client *helloClient) HelloFromAppClient(ctx context.Context) (interface{}, error) {
	log.Debug("In HelloFromAppClient method")
	resp := "Hello!"
	return resp, nil
}
