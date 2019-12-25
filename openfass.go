package go_fass

import (
	"encoding/base64"
	"github.com/vitwit/go-fass/rest"
	"os"
)

// client
type Client struct {
	rest.Request
}

// options for requestNew
type options struct {
	Endpoint string
	Host     string
	User     string
	Password string
}

func (o *options) baseURL() string {
	return o.Host + o.Endpoint
}

func GetRequestObject(cli *Client, method rest.Method, endPoint string) rest.Request {
	cli.Method = method
	cli.BaseURL = os.Getenv("HOST_URL") + endPoint

	return cli.Request
}

// GetRequest
// @return [Request] a default request object
func GetRequest(endpoint, host, user, password string) rest.Request {
	var options options

	if host == "" {
		options.Host = os.Getenv("HOST_URL")
	}

	userAndPassword := user + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(userAndPassword))

	requestHeaders := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + encodedAuth,
	}

	return rest.Request{
		BaseURL: options.baseURL(),
		Headers: requestHeaders,
	}

}

// DefaultClient is used if no custom HTTP client is defined
var DefaultClient = rest.DefaultClient

func API(request rest.Request) (*rest.Response, error) {
	return MakeRequest(request)
}

func MakeRequest(request rest.Request) (*rest.Response, error) {
	return DefaultClient.Send(request)

}

// Create a client for openFass
func NewClient(user, password, endpoint string) *Client {
	request := GetRequest(endpoint, "", user, password)
	return &Client{request}
}
