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
	_, err = cli.CreateSystemFunctions(createData)
	if err != nil {
		golog.Error("Error in CreateSystemFunctions ", err)
	}

	data := &gofaas.FunctionDefintion{
		Service: "nodeinfo",
		Image:   "functions/nodeinfo:latest",
		Labels: map[string]string{
			"changedlabelkey": "changedlabelval",
		},
	}
	_, err = cli.UpdateSystemFunctions(data)
	if err != nil {
		golog.Error("Error while creating system:  ", err)
	}

	_, err = cli.GetFunctionSummary("nodeinfo")
	if err != nil {
		golog.Error("Error in GetFunctionSummary ", err)
	}

	_, err = cli.InvokeFunction(&gofaas.SyncInvocationOpts{
		Body:         "Testing func_nodeinfo",
		FunctionName: "nodeinfo",
	})
	if err != nil {
		golog.Error("Error in InvokeFunction ", err)
	}

	_, err = cli.ScaleFunction(&gofaas.ScaleFunctionBodyOpts{
		Service:  "nodeinfo",
		Replicas: 1,
	})
	if err != nil {
		golog.Error("Error in ScaleFunction ", err)
	}
}
