package go_fass

import (
	"fmt"
	"net/http"
)

// Create a system function
func (cl *OpenFaasClient) CreateSystemFunctions(data FunctionDefintion) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, "/system/functions")
	request.Body = GetRequestBody(data)

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
func (cl *OpenFaasClient) UpdateSystemFunctions(data FunctionDefintion) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPut, "/system/functions")
	request.Body = GetRequestBody(data)

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Delete a system function
func (cl *OpenFaasClient) DeleteSystemFunction(data DeleteFunctionRequest) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodDelete, "/system/functions")
	b := GetByteData(data)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// System alert
func (cl *OpenFaasClient) SystemAlert(data map[string]interface{}) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, "/system/alert")
	b := GetByteData(data)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Invoke a function asynchronously in OpenFaaS
func (cl *OpenFaasClient) AsyncFunction(data map[string]string, functionName string) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, fmt.Sprintf("/async-function/%s", functionName))
	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Invoke a function defined in OpenFaaS
func (cl *OpenFaasClient) InvokeFunction(data map[string]string, functionName string) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, fmt.Sprintf("/function/%s", functionName))

	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Scale a function
func (cl *OpenFaasClient) ScaleFunction(data map[string]string, functionName string) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, fmt.Sprintf("/system/scale-function/%s", functionName))
	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get a summary of an OpenFaaS function
func (cl *OpenFaasClient) GetFunctionSummary(data map[string]string, functionName string) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodGet, fmt.Sprintf("/system/function/%s", functionName))
	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

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
func (cl *OpenFaasClient) CreateNewSecret(data Secret) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodPost, "/system/secrets")
	if data.Name != "" {
		b := GetByteData(data)
		request.Body = b
	}

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Update a secret.
func (cl *OpenFaasClient) UpdateSecret(data Secret) (*HTTPResponse, error) {
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
func (cl *OpenFaasClient) DeleteSecret(data SecretName) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodDelete, "/system/secrets")
	b := GetByteData(data)
	request.Body = b

	resp, err := cl.SendHTTPRequest(request)
	if err != nil {
		return &HTTPResponse{}, err
	}
	return resp, nil
}

// Get a stream of the logs for a specific function
func (cl *OpenFaasClient) GetSystemLogs(functionName, since string, tail int64, follow bool) (*HTTPResponse, error) {
	request := GetRequestDefinition(cl, http.MethodGet, fmt.Sprintf("/system/logs?name=%s&since=%s&tail=%d&follow=%t", functionName, since, tail, follow))

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
