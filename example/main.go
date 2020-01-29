package main

import (
	"github.com/kataras/golog"
	faas "github.com/vitwit/go-faas"
	"os"
)

func main() {
	cli, err := faas.NewClient(&faas.FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"), // example: http://127.0.0.1:8080
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		golog.Error(err)
		return
	}

	_, err = cli.GetSystemFunctions()
	if err != nil {
		golog.Error("Error from system functions:  ", err)
	}

	data := &faas.FunctionDefintion{
		Service: "nodeinfo",
		Image:   "functions/nodeinfo:latest",
		Limits: faas.Limits{
			Memory: "130M",
			CPU:    "0.01",
		},
	}

	resp, err := cli.UpdateSystemFunctions(data)
	if err != nil {
		golog.Error("Error while creating system:  ", err)
	}
	golog.Info(resp)
}
