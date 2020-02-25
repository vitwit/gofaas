package go_faas

import (
	"fmt"
	"net/http"
	"net/url"
)

// CreateSystemFunctions creates system function based on FunctionDefination. It returns a HTTPResponse when successful.
func (cl *OpenFaasClient) CreateSystemFunctions(def *FunctionDefintion) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, "/system/functions")
	request.Body = GetRequestBody(def)

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get system functions
func (cl *OpenFaasClient) GetSystemFunctions() (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodGet, "/system/functions")

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get system functions
func (cl *OpenFaasClient) UpdateSystemFunctions(data *FunctionDefintion) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPut, "/system/functions")
	request.Body = GetRequestBody(data)

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Delete a system function
func (cl *OpenFaasClient) DeleteSystemFunction(data *DeleteFunctionBodyOpts) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodDelete, "/system/functions")
	b := GetByteData(*data)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// System alert
func (cl *OpenFaasClient) SystemAlert(data *SystemAlertBodyOpts) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, "/system/alert")
	b := GetByteData(*data)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Invoke a function asynchronously in OpenFaaS
func (cl *OpenFaasClient) AsyncFunction(opts *AsyncInvocationOpts) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, fmt.Sprintf("/async-function/%s", opts.FunctionName))
	if opts.Body != nil {
		b := GetByteData(opts.Body)
		request.Body = b
	}

	request.Headers = map[string]string{
		"X-Callback-Url": opts.CallbackURL,
	}

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Invoke a function defined in OpenFaaS
func (cl *OpenFaasClient) InvokeFunction(opts *SyncInvocationOpts) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, fmt.Sprintf("/function/%s", opts.FunctionName))

	if opts.Body != nil {
		b := GetByteData(opts.Body)
		request.Body = b
	}

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Scale a function
func (cl *OpenFaasClient) ScaleFunction(opts *ScaleFunctionBodyOpts) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, fmt.Sprintf("/system/scale-function/%s", opts.Service))
	b := GetByteData(*opts)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get a summary of an OpenFaaS function
func (cl *OpenFaasClient) GetFunctionSummary(functionName string) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodGet, fmt.Sprintf("/system/function/%s", functionName))

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get a list of secret names and metadata from the provider
func (cl *OpenFaasClient) GetSecrets() (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodGet, "/system/secrets")

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Create a new secret.
func (cl *OpenFaasClient) CreateNewSecret(data *SecretBodyOpts) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, "/system/secrets")
	if data.Name != "" {
		b := GetByteData(*data)
		request.Body = b
	}

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Update a secret.
func (cl *OpenFaasClient) UpdateSecret(data *SecretBodyOpts) (*HTTPResponse, error) {
	if cl.ClusterType == "swarm" {
		return &HTTPResponse{}, fmt.Errorf("cannot update secret for swarm cluster")
	}

	request := GetRequestDefinition(cl, http.MethodPut, "/system/secrets")
	b := GetByteData(data)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Remove a secret.
func (cl *OpenFaasClient) DeleteSecret(data *SecretNameBodyOpts) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodDelete, "/system/secrets")
	b := GetByteData(*data)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get a stream of the logs for a specific function
func (cl *OpenFaasClient) GetSystemLogs(opts *SystemLogsQueryOpts) (*HTTPResponse, error) {
	path := fmt.Sprintf("/system/logs?name=%s", opts.Name)
	if opts.Since != "" {
		path += fmt.Sprintf("&since=%s", opts.Since)
	}
	if opts.Tail != 0 {
		path += fmt.Sprintf("&tail=%d", opts.Tail)
	}
	request := GetRequestDefinition(cl, http.MethodGet, path)

	u, err := url.Parse(request.URL)
	if err != nil {
		return &HTTPResponse{}, err
	}
	request.URL = u.String()

	fmt.Println(request.URL)

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get info such as provider version number and provider orchestrator
func (cl *OpenFaasClient) GetSystemInfo() (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodGet, "/system/info")

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Health Check
func (cl *OpenFaasClient) GetHealthz() (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodGet, "/healthz")

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}
