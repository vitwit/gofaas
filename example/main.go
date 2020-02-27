package main

import (
	"fmt"
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

	fmt.Println("client ", cli)
	if err != nil {
		golog.Error(err)
		return
	}

	//_, err = cli.GetSystemFunctions()
	//if err != nil {
	//	golog.Error("Error from system functions:  ", err)
	//}

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
	createResp, createErr := cli.CreateSystemFunctions(createData)
	// if createErr != nil {
	// golog.Error("error while calling %s/system/functions: %v", cli.GatewayAddress, createErr)
	// }
	fmt.Println("create response==> ", createResp, createErr)

	data := &gofaas.FunctionDefintion{
		Service: "nodeinfo",
		Image:   "functions/nodeinfo:latest",
		Labels: map[string]string{
			"changedlabelkey": "changedlabelval",
		},
	}
	resp, updateErr := cli.UpdateSystemFunctions(data)
	if updateErr != nil {
		golog.Error("Error while creating system:  ", updateErr)
	}
	fmt.Println("UpdateSystemFunctions resp ==> ", resp, updateErr)

	res, summaryErr := cli.GetFunctionSummary("nodeinfo")
	// if summaryErr != nil {
	// golog.Error("error while getting summary for func %s: %v", "nodeinfo", summaryErr)
	// }
	fmt.Println("GetFunctionSummary==> ", res, summaryErr)

	response, invokeErr := cli.InvokeFunction(&gofaas.SyncInvocationOpts{
		Body:         "Testing func_nodeinfo",
		FunctionName: "nodeinfo",
	})
	// if invokeErr != nil {
	// golog.Error("error while invoking func %s: %v", "nodeinfo", invokeErr)
	// }
	fmt.Println("InvokeFunction resp==> ", response, invokeErr)

	response1, scaleErr := cli.ScaleFunction(&gofaas.ScaleFunctionBodyOpts{
		Service:  "nodeinfo",
		Replicas: 1,
	})
	// if scaleErr != nil {
	// golog.Error("error while scaling func %s: %v", "nodeinfo", scaleErr)
	// }
	fmt.Println("ScaleFunction resp==> ", response1, scaleErr)
}
