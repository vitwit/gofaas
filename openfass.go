package go_fass

import (
	"encoding/base64"
	"os"
	"strings"
)

func GetRequestDefinition(cli *OpenFaasClient, method, path string) *FaasRequestDefinition {
	cli.Method = method
	cli.Path = path // path expects a trailing slash
	cli.URL = cli.GatewayAddress + cli.Path
	return &cli.FaasRequestDefinition
}

func getGatewayAddress(gateway string) string {
	var addr string
	if gateway == "" {
		addr = os.Getenv("HOST_URL")
	} else {
		addr = gateway
	}
	// remove leading slash if any
	if string(addr[len(addr)-1]) == "/" {
		addr = strings.TrimRight(addr, "/")
	}
	return addr
}

func SetClientRequestOpts(creds *FaasGatewayCredentials) FaasRequestDefinition {
	userAndPassword := creds.Username + ":" + creds.Password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(userAndPassword))
	requestHeaders := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + encodedAuth,
	}

	return FaasRequestDefinition{
		GatewayAddress: getGatewayAddress(creds.GatewayAddress),
		Headers:        requestHeaders,
	}
}

func NewClient(creds *FaasGatewayCredentials) *OpenFaasClient {
	request := SetClientRequestOpts(creds)
	return &OpenFaasClient{request}
}
