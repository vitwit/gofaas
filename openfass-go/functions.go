package openfass_go

import (
	"fmt"
	"github.com/vitwit/go-fass/rest"
)

const host = "http://127.0.0.1:8080"

// Create a system function
func (cl *Client) CreateSystemFunctions(data FunctionDefintion) (*rest.Response, error) {
	request := GetRequestObject(cl, "POST", host+"/system/functions")
	request.Body = GetRequestBody(data)

	return MakeRequest(request)
}

// Get system functions
func (cl *Client) GetSystemFunctions(user, password string) (*rest.Response, error) {
	request := GetRequestObject(cl, "GET", host+"/system/functions")

	return MakeRequest(request)
}

// Get system functions
func (cl *Client) UpdateSystemFunctions(data FunctionDefintion) (*rest.Response, error) {
	request := GetRequestObject(cl, "PUT", host+"/system/functions")
	request.Body = GetRequestBody(data)

	return MakeRequest(request)
}

// Delete a system function

func (cl *Client) DeleteSystemFunction(data DeleteFunctionRequest) (*rest.Response, error) {
	request := GetRequestObject(cl, "DELETE", host+"/system/functions")
	b := GetByteData(data)
	request.Body = b

	return MakeRequest(request)
}

// System alert
func (cl *Client) SystemAlert(data SystemAlert) (*rest.Response, error) {
	request := GetRequestObject(cl, "POST", host+"/system/alert")
	b := GetByteData(data)
	request.Body = b

	return MakeRequest(request)
}

//Invoke a function asynchronously in OpenFaaS
func (cl *Client) AsyncFunction(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/async-function/%s", functionName)
	request := GetRequestObject(cl, "POST", host+s)

	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

	return MakeRequest(request)
}

//Invoke a function defined in OpenFaaS
func (cl *Client) InvokeFunction(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/function/%s", functionName)
	request := GetRequestObject(cl, "POST", host+s)

	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

	return MakeRequest(request)
}

//Scale a function
func (cl *Client) ScaleFunction(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/system/scale-function/%s", functionName)
	request := GetRequestObject(cl, "POST", host+s)

	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

	return MakeRequest(request)
}

//Get a summary of an OpenFaaS function
func (cl *Client) GetFunctionSummary(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/system/function/%s", functionName)
	request := GetRequestObject(cl, "GET", host+s)

	if data != nil {
		b := GetByteData(data)
		request.Body = b
	}

	return MakeRequest(request)
}

//Get a list of secret names and metadata from the provider
func (cl *Client) GetSectrets() (*rest.Response, error) {
	request := GetRequestObject(cl, "GET", host+"/system/secrets")

	return MakeRequest(request)
}

// Create a new secret.
func (cl *Client) CreateNewSecret(data Secret) (*rest.Response, error) {
	request := GetRequestObject(cl, "POST", host+"/system/secrets")

	if data.Name != "" {
		b := GetByteData(data)
		request.Body = b
	}

	return MakeRequest(cl.Request)
}

// Update a secret.
func (cl *Client) UpdateSecret(data Secret) (*rest.Response, error) {
	request := GetRequestObject(cl, "PUT", host+"/system/secrets")

	b := GetByteData(data)
	request.Body = b

	return MakeRequest(cl.Request)
}

// Remove a secret.
func (cl *Client) DeleteSecret(data SecretName) (*rest.Response, error) {
	request := GetRequestObject(cl, "DELETE", host+"/system/secrets")

	b := GetByteData(data)
	request.Body = b

	return MakeRequest(cl.Request)
}

// Get a stream of the logs for a specific function
func (cl *Client) GetSystemLogs(functionName, since string, tail int64, follow bool) (*rest.Response, error) {
	s := fmt.Sprintf("/system/logs?name=%s&since=%s&tail=%d&follow=%t", functionName, since, tail, follow)
	request := GetRequestObject(cl, "GET", host+s)

	return MakeRequest(request)
}

// Get info such as provider version number and provider orchestrator
func (cl *Client) GetSystemInfo() (*rest.Response, error) {
	request := GetRequestObject(cl, "GET", host+"/system/info")

	return MakeRequest(request)
}

// Healthcheck
func (cl *Client) GetHealthz() (*rest.Response, error) {
	request := GetRequestObject(cl, "GET", host+"/healthz")

	return MakeRequest(request)
}
