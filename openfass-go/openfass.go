package openfass_go

import (
	"encoding/base64"
	"github.com/vitwit/go-fass/rest"
	"os"
)

// client
type Client struct {
	// rest.Request
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
	cli.BaseURL = os.Getenv("hostUrl") + endPoint

	return cli.Request
}

// GetRequest
// @return [Request] a default request object
func GetRequest(endpoint, host, user, password string) rest.Request {
	return requestNew(options{endpoint, host, user, password})
}

// requestNew create Request
// @return [Request] a default request object
func requestNew(options options) rest.Request {
	if options.Host == "" {
		options.Host = "http://127.0.0.1:8080"
	}

	userAndPassword := options.User + ":" + options.Password
	encode := base64.StdEncoding.EncodeToString([]byte(userAndPassword))

	requestHeaders := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + encode,
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

// MakeRequest attempts a Twilio SendGrid request synchronously.
func MakeRequest(request rest.Request) (*rest.Response, error) {
	return DefaultClient.Send(request)

}

// NewSendClient constructs a new Twilio SendGrid client given an API key
func NewSendClient(user, password, endpoint string) *Client {
	request := GetRequest(endpoint, "", user, password)
	//request.Method = "POST"
	return &Client{request}
}
