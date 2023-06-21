package request

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

type RequestPayload struct {
	URL         string
	Method      string      // Use http method constants
	Headers     http.Header // Use http header
	QueryParams url.Values  // Use url value
	Payload     []byte
}

// MakeHttpRequest returns response body as a byte slice.
func (rp *RequestPayload) MakeHttpRequest() (resp []byte, err error) {
	client := &http.Client{}
	// Create request - Add payload
	request, err := http.NewRequest(rp.Method, rp.URL, nil)
	if rp.Payload != nil {
		request, err = http.NewRequest(rp.Method, rp.URL, bytes.NewReader(rp.Payload))
	}
	if err != nil {
		return
	}
	// Headers
	request.Header = rp.Headers
	// Query Params
	request.URL.RawQuery = rp.QueryParams.Encode()
	// Request
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	resp, err = io.ReadAll(response.Body)
	return
}

// MakeHttpRequestResponse returns full response.
func (rp *RequestPayload) MakeHttpRequestResponse() (response *http.Response, err error) {
	client := &http.Client{}
	// Create request - Add payload
	request, err := http.NewRequest(rp.Method, rp.URL, nil)
	if rp.Payload != nil {
		request, err = http.NewRequest(rp.Method, rp.URL, bytes.NewReader(rp.Payload))
	}
	if err != nil {
		return
	}
	// Headers
	request.Header = rp.Headers
	// Query Params
	request.URL.RawQuery = rp.QueryParams.Encode()
	// Request
	response, err = client.Do(request)
	if err != nil {
		return
	}
	return
}
