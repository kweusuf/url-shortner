package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kweusuf/url-shortner/pkg/constants"
)

type URLRequest struct {
	context.Context
	URL string `json:"url"`
}
type URLResponse struct {
	ShortenedURL string `json:"shortened_url"`
}
type URLShortenerError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *URLShortenerError) Error() string {
	return e.Message
}

type URLShortenerResponse struct {
	ShortenedURL string `json:"shortened_url"`
	OriginalURL  string `json:"original_url"`
	StatusCode   int    `json:"status_code"`
	Message      string `json:"message"`
}

func NewURLShortenerError(message string, code int) *URLShortenerError {
	return &URLShortenerError{
		Message: message,
		Code:    code,
	}
}
func NewURLShortenerResponse(shortenedURL, originalURL string, statusCode int, message string) *URLShortenerResponse {
	return &URLShortenerResponse{
		ShortenedURL: shortenedURL,
		OriginalURL:  originalURL,
		StatusCode:   statusCode,
		Message:      message,
	}
}
func (r *URLShortenerResponse) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}
func (r *URLShortenerResponse) IsError() bool {
	return r.StatusCode < 200 || r.StatusCode >= 300
}
func (r *URLShortenerResponse) Error() string {
	if r.IsError() {
		return r.Message
	}
	return ""
}
func (r *URLShortenerResponse) String() string {
	return "URLShortenerResponse{" +
		"ShortenedURL='" + r.ShortenedURL + "'" +
		", OriginalURL='" + r.OriginalURL + "'" +
		", StatusCode=" + fmt.Sprintf("%d", r.StatusCode) +
		", Message='" + r.Message + "'" +
		"}"
}
func (r *URLShortenerResponse) ToJSON() string {
	data, err := json.Marshal(r)
	if err != nil {
		return "{}"
	}
	return string(data)
}
func ParseURLRequest(r *http.Request) (*URLRequest, error) {
	if r.Method != http.MethodPost {
		return nil, fmt.Errorf("invalid method: %s, expected POST", r.Method)
	}
	if r.Header.Get(constants.ContentType) != constants.ApplicationJSONContentType {
		return nil, fmt.Errorf("invalid content type: %s, expected %s", r.Header.Get(constants.ContentType), constants.ApplicationJSONContentType)
	}
	var request URLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}
	if strings.TrimSpace(request.URL) == "" {
		return nil, fmt.Errorf("url cannot be empty")
	}
	return &request, nil
}
func ParseURLResponse(resp *http.Response) (*URLShortenerResponse, error) {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp URLShortenerError
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, fmt.Errorf("failed to decode error response: %w", err)
		}
		return nil, &errResp
	}
	var urlResp URLShortenerResponse
	if err := json.NewDecoder(resp.Body).Decode(&urlResp); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	urlResp.StatusCode = resp.StatusCode
	return &urlResp, nil
}
func EncodeURLRequest(req *URLRequest, apiURL string) (*http.Request, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request body: %w", err)
	}
	httpReq, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	httpReq.Header.Set(constants.ContentType, constants.ApplicationJSONContentType)
	return httpReq, nil
}
func EncodeURLResponse(resp *URLShortenerResponse) (*http.Response, error) {
	body, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to encode response body: %w", err)
	}
	httpResp := &http.Response{
		StatusCode: resp.StatusCode,
		Body:       ioutil.NopCloser(strings.NewReader(string(body))),
	}
	httpResp.Header.Set(constants.ContentType, constants.ApplicationJSONContentType)
	return httpResp, nil
}
func DecodeURLRequest(r *http.Request) (*URLRequest, error) {
	if r.Method != http.MethodPost {
		return nil, &URLShortenerError{
			Message: "Invalid method, expected POST",
			Code:    http.StatusMethodNotAllowed,
		}
	}
	if r.Header.Get(constants.ContentType) != constants.ApplicationJSONContentType {
		return nil, &URLShortenerError{
			Message: "Invalid content type, expected application/json",
			Code:    http.StatusUnsupportedMediaType,
		}
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, &URLShortenerError{
			Message: "Failed to read request body",
			Code:    http.StatusBadRequest,
		}
	}
	var request URLRequest
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, &URLShortenerError{
			Message: "Failed to decode request body",
			Code:    http.StatusBadRequest,
		}
	}
	if request.URL == "" {
		return nil, &URLShortenerError{
			Message: "URL cannot be empty",
			Code:    http.StatusBadRequest,
		}
	}
	return &request, nil
}

func DecodeURLResponse(resp *http.Response) (*URLShortenerResponse, error) {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp URLShortenerError
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read error response body: %w", err)
		}
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("failed to decode error response: %w", err)
		}
		return nil, &errResp
	}
	var urlResp URLShortenerResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if err := json.Unmarshal(body, &urlResp); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	urlResp.StatusCode = resp.StatusCode
	return &urlResp, nil
}
