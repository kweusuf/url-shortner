package service

import (
	helloservice "github.com/kweusuf/url-shortner/pkg/service/hello"
	urlservice "github.com/kweusuf/url-shortner/pkg/service/url"
)

type Services struct {
	HelloService helloservice.HelloService
	URLService   urlservice.URLService
	// JobManagerService jobservice.JobManagerService
}
