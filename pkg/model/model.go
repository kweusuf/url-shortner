package model

import "context"

type GetRequest struct {
	context.Context
	QueryParams map[string][]string
}

type PostRequest struct {
	context.Context
	Body interface{}
}

type PutRequest struct {
	context.Context
	Body interface{}
}

type DeleteRequest struct {
	context.Context
	ID string
}

type AppURI struct {
	Host       string
	Port       string
	HttpScheme string
}

type GetConfResponse struct {
	HttpStatus int
	Body       interface{}
}
