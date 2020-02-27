package gofaas

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type defaultClient struct {
	HTTPClient *http.Client
}

func (dc *defaultClient) new() *defaultClient {
	return &defaultClient{HTTPClient: http.DefaultClient}
}

// AddQueryParameters adds query parameters to the URL.
func (cl *OpenFaasClient) AddQueryParameters(req *http.Request, queryParams QueryParams) {
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	req.URL.RawQuery = params.Encode()
}

// BuildHTTPRequest creates the HTTP request object.
func (cl *OpenFaasClient) BuildHTTPRequest(reqDef *FaasRequestDefinition) (*http.Request, error) {
	// make new request
	req, err := http.NewRequest(reqDef.Method, reqDef.URL, bytes.NewBuffer(reqDef.Body))
	if err != nil {
		return nil, err
	}

	// Add any query parameters to the URL.
	if len(reqDef.QueryParams) != 0 {
		cl.AddQueryParameters(req, reqDef.QueryParams)
	}

	// set headers
	for key, value := range reqDef.Headers {
		req.Header.Set(key, value)
	}
	_, exists := req.Header["Content-Type"]
	if len(reqDef.Body) > 0 && !exists {
		req.Header.Set("Content-Type", "application/json")
	}

	// return
	return req, nil
}

// BuildSuccessResponse will builds the response struct.
func (cl *OpenFaasClient) BuildSuccessResponse(res *http.Response) (*HTTPResponse, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &HTTPResponse{}, err
	}
	response := &HTTPResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
		Headers:    res.Header,
	}
	_ = res.Body.Close()
	return response, nil
}

// SendHTTPRequest will send the http req and return http response
func (cl *OpenFaasClient) SendHTTPRequest(req *FaasRequestDefinition) (*HTTPResponse, error) {
	httpReq, err := cl.BuildHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	var dc defaultClient
	newDefaultClient := dc.new()
	resp, err := newDefaultClient.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	successResp, err := cl.BuildSuccessResponse(resp)
	if err != nil {
		return nil, err
	}
	return successResp, nil
}
