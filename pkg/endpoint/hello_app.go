package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	service "github.com/kweusuf/url-shortner/pkg/service/hello"
	"github.com/kweusuf/url-shortner/pkg/utils"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

func makeHelloEndpoint(svc service.HelloService) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (response interface{}, err error) {
		log.Debug("In makeHelloEndpoint method")
		result, err := svc.HelloFromAppService(context)
		return utils.ConstructResponse(result, err)
	}
}
