package go_fass

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DefaultClient struct {
	HTTPClient *http.Client
}

func (dc *DefaultClient) New() *DefaultClient {
	return &DefaultClient{HTTPClient: http.DefaultClient}
}

// AddQueryParameters adds query parameters to the URL.
func (cl *OpenFaasClient) AddQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

// BuildRequestObject creates the HTTP request object.
func (cl *OpenFaasClient) BuildHTTPRequest(request *FaasRequestDefinition) (*http.Request, error) {
	// Add any query parameters to the URL.
	if len(request.QueryParams) != 0 {
		request.URL = cl.AddQueryParameters(request.URL, request.QueryParams)
	}

	// make new request
	req, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(request.Body))
	if err != nil {
		return nil, err
	}

	// set headers
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	_, exists := req.Header["Content-Type"]
	if len(request.Body) > 0 && !exists {
		req.Header.Set("Content-Type", "application/json")
	}

	// return
	return req, nil
}

// BuildResponse builds the response struct.
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

// Send the http req and return response
func (cl *OpenFaasClient) SendHTTPRequest(req *FaasRequestDefinition) (*HTTPResponse, error) {
	httpReq, err := cl.BuildHTTPRequest(req)
	if err != nil {
		return nil, err
	}

	var dc DefaultClient
	defaultClient := dc.New()
	resp, err := defaultClient.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	successResp, err := cl.BuildSuccessResponse(resp)
	if err != nil {
		return nil, err
	}
	return successResp, nil
}
