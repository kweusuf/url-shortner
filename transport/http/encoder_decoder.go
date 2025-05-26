package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kweusuf/url-shortner/pkg/constants"
	"github.com/kweusuf/url-shortner/pkg/model"
)

func decodeGetRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	query := req.URL.Query()
	request := model.GetRequest{
		Context:     ctx,
		QueryParams: query,
	}

	return request, nil
}

// func decodePostRequest(ctx context.Context, req *http.Request) (interface{}, error) {
// 	var request model.PostRequest
// 	err := json.NewDecoder(req.Body).Decode(&request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Context = ctx
// 	return request, nil
// }

// func decodePutRequest(ctx context.Context, req *http.Request) (interface{}, error) {
// 	var request model.PutRequest
// 	err := json.NewDecoder(req.Body).Decode(&request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Context = ctx
// 	return request, nil
// }

// func decodeDeleteRequest(ctx context.Context, req *http.Request) (interface{}, error) {
// 	request := model.DeleteRequest{
// 		Context: ctx,
// 		ID:      req.URL.Query().Get("id"),
// 	}
// 	return request, nil
// }

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(constants.ContentType, constants.ApplicationJSONContentType)
	w.WriteHeader(http.StatusAccepted)
	return json.NewEncoder(w).Encode(response)
}
