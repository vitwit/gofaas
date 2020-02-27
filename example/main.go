package main

import (
	"os"

	"github.com/kataras/golog"
	gofaas "github.com/vitwit/gofaas"
)

func main() {
	cli, err := gofaas.NewClient(&gofaas.FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USER"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"), // example: http://127.0.0.1:8080
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		golog.Error(err)
		return
	}

	createData := &gofaas.FunctionDefintion{
		Service:    "nodeinfo",
		Network:    "func_functions",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
		EnvVars: gofaas.EnvVars{
			AdditionalProp1: "string",
			AdditionalProp2: "string",
			AdditionalProp3: "string",
		},
		Constraints: []string{
			"node.platform.os == linux",
		},
		Labels: map[string]string{
			"example": "func1",
		},
		Annotations: gofaas.Annotations{
			Topics: "awesome-kafka-topic",
			Foo:    "some",
		},
		RegistryAuth: "dXNlcjpwYXNzd29yZA==",
		Limits: gofaas.Limits{
			Memory: "128M",
			CPU:    "0.01",
		},
		Requests: gofaas.Requests{
			Memory: "128M",
			CPU:    "0.01",
		},
		ReadOnlyRootFilesystem: true,
	}
	_, createErr := cli.CreateSystemFunctions(createData)
	if createErr != nil {
		golog.Error("Error in CreateSystemFunctions ", createErr)
	}

	data := &gofaas.FunctionDefintion{
		Service: "nodeinfo",
		Image:   "functions/nodeinfo:latest",
		Labels: map[string]string{
			"changedlabelkey": "changedlabelval",
		},
	}
	_, updateErr := cli.UpdateSystemFunctions(data)
	if updateErr != nil {
		golog.Error("Error while creating system:  ", updateErr)
	}

	_, summaryErr := cli.GetFunctionSummary("nodeinfo")
	if summaryErr != nil {
		golog.Error("Error in GetFunctionSummary ", summaryErr)
	}

	_, invokeErr := cli.InvokeFunction(&gofaas.SyncInvocationOpts{
		Body:         "Testing func_nodeinfo",
		FunctionName: "nodeinfo",
	})
	if invokeErr != nil {
		golog.Error("Error in InvokeFunction ", invokeErr)
	}

	_, scaleErr := cli.ScaleFunction(&gofaas.ScaleFunctionBodyOpts{
		Service:  "nodeinfo",
		Replicas: 1,
	})
	if scaleErr != nil {
		golog.Error("Error in ScaleFunction ", scaleErr)
	}
}
