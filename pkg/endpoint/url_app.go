package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kweusuf/url-shortner/pkg/model"
	service "github.com/kweusuf/url-shortner/pkg/service/url"
	"github.com/kweusuf/url-shortner/pkg/utils"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

func makeURLShortenEndpoint(svc service.URLService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		log.Debug("In makeURLShortenEndpoint method")
		url := request.(model.URLRequest).URL
		result, err := svc.ShortenURLService(ctx, url)
		return utils.ConstructResponse(result, err)
	}
}
