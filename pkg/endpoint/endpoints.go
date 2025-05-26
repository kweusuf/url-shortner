package endpoint

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/kweusuf/url-shortner/pkg/service"
)

type AppEndpoints struct {
	HelloEndpoint endpoint.Endpoint

	// GetAllJobsEndpoint       endpoint.Endpoint
	// GetAllActiveJobsEndpoint endpoint.Endpoint
	// GetJobEndpoint           endpoint.Endpoint
	// CreateJobEndpoint        endpoint.Endpoint
	// UpdateJobEndpoint        endpoint.Endpoint
	// DeleteJobEndpoint        endpoint.Endpoint
}

func MakeEndpoints(services service.Services) AppEndpoints {
	return AppEndpoints{
		HelloEndpoint: makeHelloEndpoint(services.HelloService),

		// GetAllJobsEndpoint:       makeGetAllJobsEndpoint(services.JobManagerService),
		// GetAllActiveJobsEndpoint: makeGetAllActiveJobsEndpoint(services.JobManagerService),
		// GetJobEndpoint:           makeGetJobEndpoint(services.JobManagerService),
		// CreateJobEndpoint:        makeCreateJobEndpoint(services.JobManagerService),
		// UpdateJobEndpoint:        makeUpdateJobEndpoint(services.JobManagerService),
		// DeleteJobEndpoint:        makeDeleteJobEndpoint(services.JobManagerService),
	}
}
