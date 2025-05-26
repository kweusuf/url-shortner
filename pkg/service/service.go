package service

import (
	helloservice "github.com/kweusuf/url-shortner/pkg/service/hello"
)

type Services struct {
	HelloService helloservice.HelloService
	// JobManagerService jobservice.JobManagerService
}
