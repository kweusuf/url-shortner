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
	router.Methods(http.MethodGet).Path(constants.API_V1 + "/hello").Handler(
		kithttp.NewServer(endpoints.HelloEndpoint,
			decodeGetRequest,
			encodeResponse,
			kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// URL Shorten Endpoint
	// /api/v1/shorten
	router.Methods(http.MethodPost).Path(constants.API_V1 + "/shorten").Handler(
		kithttp.NewServer(endpoints.URLShortenEndpoint,
			decodePostRequest,
			encodeResponse,
			kithttp.ServerErrorEncoder(httpErrorEncoder)))

	// URL Expand Endpoint
	router.Methods(http.MethodGet).Path(constants.API_V1 + "/s/{shortCode}").Handler(
		kithttp.NewServer(endpoints.URLExpandEndpoint,
			decodeGetRequest,
			encodeResponse,
			kithttp.ServerErrorEncoder(httpErrorEncoder)))

}
