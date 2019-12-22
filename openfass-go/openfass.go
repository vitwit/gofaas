package openfass_go

import (
"encoding/base64"
"github.com/vitwit/go-fass/rest"
)

// Version is this client library's current version
const (
	Version        = "3.1.0"
	rateLimitRetry = 5
	rateLimitSleep = 1100
	user           = "admin"
	password       = "9a651783248f8b61a92ae2fa02c7cb627a68e061f5cc0773822f72e0c7af077d"
)

// Client is the Twilio SendGrid Go client
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
		//"Authorization": "Basic YWRtaW46OWE2NTE3ODMyNDhmOGI2MWE5MmFlMmZhMDJjN2NiNjI3YTY4ZTA2MWY1Y2MwNzczODIyZjcyZTBjN2FmMDc3ZA==",
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
