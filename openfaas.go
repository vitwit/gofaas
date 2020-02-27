package gofaas

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// GetRequestDefinition will return request definition with given method, path and URL
func GetRequestDefinition(cli *OpenFaasClient, method, path string) *FaasRequestDefinition {
	cli.Method = method
	cli.Path = path // path expects a trailing slash
	cli.URL = cli.GatewayAddress + cli.Path
	return &cli.FaasRequestDefinition
}

// GetGatewayAddress will return gateway address
func GetGatewayAddress(gateway string) string {
	if gateway == "" {
		gateway = os.Getenv("OPENFAAS_GATEWAY_ADDR")
	}
	// remove leading slash if any
	if gateway[len(gateway)-1:] == "/" {
		gateway = strings.TrimRight(gateway, "/")
	}

	return gateway
}

// setClientRequestOpts will return request definition with given credentials and request headers
func setClientRequestOpts(creds *FaasGatewayCredentials) FaasRequestDefinition {
	userAndPassword := creds.Username + ":" + creds.Password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(userAndPassword))
	requestHeaders := map[string]string{
		"Accept":        "application/json",
		"Authorization": "Basic " + encodedAuth,
	}

	return FaasRequestDefinition{
		GatewayAddress: GetGatewayAddress(creds.GatewayAddress),
		Headers:        requestHeaders,
		ClusterType:    creds.ClusterType,
	}
}

// NewClient will return a new client with given credentials
func NewClient(creds *FaasGatewayCredentials) (*OpenFaasClient, error) {
	if creds.ClusterType != "swarm" && creds.ClusterType != "kubernetes" {
		return &OpenFaasClient{}, fmt.Errorf("invalid cluster type %v", creds.ClusterType)
	}
	return &OpenFaasClient{setClientRequestOpts(creds)}, nil
}
