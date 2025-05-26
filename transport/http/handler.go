package http

import (
	"crypto/tls"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/kweusuf/url-shortner/pkg/constants"
	"github.com/kweusuf/url-shortner/pkg/endpoint"
)

var httpErrorEncoder kithttp.ErrorEncoder

func NewHttpHandler(endpoints endpoint.AppEndpoints) http.Handler {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	// var options []kithttp.ServerOption
	makeHandler(router, endpoints /* , options */)
	return router
}

func handler(writer http.ResponseWriter, request *http.Request) {}

func makeHandler(router *mux.Router, endpoints endpoint.AppEndpoints /* , options []kithttp.ServerOption */) {
	http.DefaultTransport.(*http.Transport).TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)

	router.HandleFunc("/", handler)

	// TODO: Add Middleware here for auth

	// Hello Endpoint
	// /api/v1/hello
	router.Methods(http.MethodGet).Path(constants.API + "/hello").Handler(
		kithttp.NewServer(endpoints.HelloEndpoint,
			decodeGetRequest,
			encodeResponse,
			kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// // /api/v1/getJobs
	// router.Methods(http.MethodGet).Path(constants.API + "/getJobs").Handler(
	// 	kithttp.NewServer(endpoints.GetAllJobsEndpoint,
	// 		decodeGetRequest,
	// 		encodeResponse,
	// 		kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// // /api/v1/getActiveJobs
	// router.Methods(http.MethodGet).Path(constants.API + "/getActiveJobs").Handler(
	// 	kithttp.NewServer(endpoints.GetAllActiveJobsEndpoint,
	// 		decodeGetRequest,
	// 		encodeResponse,
	// 		kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// // /api/v1/getJob
	// router.Methods(http.MethodGet).Path(constants.API + "/getJob").Handler(
	// 	kithttp.NewServer(endpoints.GetJobEndpoint,
	// 		decodeGetRequest,
	// 		encodeResponse,
	// 		kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// // /api/v1/createJob
	// router.Methods(http.MethodPost).Path(constants.API + "/createJob").Handler(
	// 	kithttp.NewServer(endpoints.CreateJobEndpoint,
	// 		decodePostRequest,
	// 		encodeResponse,
	// 		kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// // /api/v1/updateJob
	// router.Methods(http.MethodPut).Path(constants.API + "/updateJob").Handler(
	// 	kithttp.NewServer(endpoints.UpdateJobEndpoint,
	// 		decodePutRequest,
	// 		encodeResponse,
	// 		kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// // /api/v1/deleteJob
	// router.Methods(http.MethodDelete).Path(constants.API + "/deleteJob").Handler(
	// 	kithttp.NewServer(endpoints.DeleteJobEndpoint,
	// 		decodeDeleteRequest,
	// 		encodeResponse,
	// 		kithttp.ServerErrorEncoder(httpErrorEncoder)))

}
