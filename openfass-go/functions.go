package openfass_go

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/vitwit/go-fass/rest"

)

const host = "http://127.0.0.1:8080"

// Create a system function
func (cl *Client) CreateSystemFunctions(data FunctionDefintion) (*rest.Response, error) {
	//request := GetRequest("/system/functions", "", user, password)
	cl.Method = "POST"
	cl.BaseURL = host + "/system/functions"
	cl.Body = GetRequestBody(data)
	return MakeRequest(cl.Request)
}

// Get system functions
func (cl *Client) GetSystemFunctions(user, password string) (*rest.Response, error) {
	//request := GetRequest("/system/functions", "", user, password)
	cl.Method = "GET"
	cl.BaseURL = host + "/system/functions"
	return MakeRequest(cl.Request)
}

// Get system functions
func (cl *Client) UpdateSystemFunctions(data FunctionDefintion) (*rest.Response, error) {
	//request := GetRequest("/system/functions", "", user, password)
	cl.Method = "PUT"
	cl.BaseURL = host + "/system/functions"
	cl.Body = GetRequestBody(data)
	return MakeRequest(cl.Request)
}

// Delete a system function

func (cl *Client) DeleteSystemFunction(data DeleteFunctionRequest) (*rest.Response, error) {
	//request := GetRequest("/system/functions", "", user, password)
	cl.Method = "DELETE"
	cl.BaseURL = host + "/system/functions"

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	cl.Body = b
	return MakeRequest(cl.Request)
}

// System alert
func (cl *Client) SystemAlert(data SystemAlert) (*rest.Response, error) {
	//request := GetRequest("/system/alert", "", user, password)
	cl.Method = "POST"
	//cl.BaseURL = host + "/system/functions"

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	cl.Body = b
	return MakeRequest(cl.Request)
}

//Invoke a function asynchronously in OpenFaaS
func (cl *Client) AsyncFunction(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/async-function/%s", functionName)
	//request := GetRequest(s, "", user, password)
	cl.Method = "POST"
	cl.BaseURL = host + s

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		cl.Body = b
	}

	return MakeRequest(cl.Request)
}

//Invoke a function defined in OpenFaaS
func (cl *Client) InvokeFunction(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/function/%s", functionName)
	//request := GetRequest(s, "", user, password)
	cl.Method = "POST"
	cl.BaseURL = host + s

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		cl.Body = b
	}

	return MakeRequest(cl.Request)
}

//Scale a function
func (cl *Client) ScaleFunction(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/system/scale-function/%s", functionName)
	//request := GetRequest(s, "", user, password)
	cl.Method = "POST"
	cl.BaseURL = host + s

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		cl.Body = b
	}

	return MakeRequest(cl.Request)
}

//Get a summary of an OpenFaaS function
func (cl *Client) GetFunctionSummary(data map[string]string, functionName string) (*rest.Response, error) {
	s := fmt.Sprintf("/system/function/%s", functionName)
	//request := GetRequest(s, "", user, password)
	cl.Method = "GET"
	cl.BaseURL = host + s

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		cl.Body = b
	}

	return MakeRequest(cl.Request)
}

//Get a list of secret names and metadata from the provider
func (cl *Client) GetSectrets() (*rest.Response, error) {
	//request := GetRequest("/system/secrets", "", user, password)
	cl.Method = "GET"
	cl.BaseURL = host + "/system/secrets"

	return MakeRequest(cl.Request)
}

// Create a new secret.
func (cl *Client) CreateNewSecret(data Secret) (*rest.Response, error) {
	//request := GetRequest("/system/secrets", "", user, password)
	cl.Method = "POST"
	cl.BaseURL = host + "/system/secrets"

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	cl.Body = b

	return MakeRequest(cl.Request)
}

// Update a secret.
func (cl *Client) UpdateSecret(data Secret) (*rest.Response, error) {
	//request := GetRequest("/system/secrets", "", user, password)
	cl.Method = "PUT"
	cl.BaseURL = host + "/system/secrets"

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	cl.Body = b

	return MakeRequest(cl.Request)
}

// Remove a secret.
func (cl *Client) DeleteSecret(data SecretName) (*rest.Response, error) {
	//request := GetRequest("/system/secrets", "", user, password)
	cl.Method = "DELETE"

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	cl.Body = b

	return MakeRequest(cl.Request)
}

// Get a stream of the logs for a specific function
func (cl *Client) GetSystemLogs(functionName, since string, tail int64, follow bool) (*rest.Response, error) {
	s := fmt.Sprintf("/system/logs?name=%s&since=%s&tail=%d&follow=%t", functionName, since, tail, follow)

	//request := GetRequest(s, "", user, password)
	cl.Method = "GET"
	cl.BaseURL = host + s

	return MakeRequest(cl.Request)
}

// Get info such as provider version number and provider orchestrator
func (cl *Client) GetSystemInfo() (*rest.Response, error) {
	//request := GetRequest("/system/info", "", user, password)
	cl.Method = "GET"
	cl.BaseURL = host + "/system/info"

	return MakeRequest(cl.Request)
}

// Healthcheck
func (cl *Client) GetHealthz() (*rest.Response, error) {
	//request := GetRequest("/healthz", "", user, password)
	cl.Method = "GET"
	cl.BaseURL = host + "/healthz"

	return MakeRequest(cl.Request)
}
