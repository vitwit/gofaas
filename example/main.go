package main

import (
	"os"

	"github.com/kataras/golog"
	fass "github.com/vitwit/go-fass"
)

func main() {
	cli := fass.NewClient(&fass.FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"), // example: http://127.0.0.1:8080
	})

	golog.Info("HOST: ", cli.URL)

	res, err := cli.GetSystemFunctions()
	if err != nil {
		golog.Error("Error from system functions:  ", err)
	}

	golog.Info("response of get system:  ", res.Body)

	data := fass.FunctionDefintion{
		Service:    "nodeinfo12345",
		Network:    "func_functions",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
		EnvVars: fass.EnvVars{
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
		Annotations: fass.Annotations{
			Topics: "awesome-kafka-topic",
			Foo:    "some",
		},
		RegistryAuth: "dXNlcjpwYXNzd29yZA==",
		Limits: fass.Limits{
			Memory: "128M",
			CPU:    "0.01",
		},
		Requests: fass.Requests{
			Memory: "128M",
			CPU:    "0.01",
		},
		ReadOnlyRootFilesystem: true,
	}

	resp, err := cli.CreateSystemFunctions(data)
	if err != nil {
		golog.Error("Error while creating system:  ", err)
	}

	golog.Info("response of create system: ", resp)

	dd := map[string]interface{}{
		"receiver": "scale-up",
		"status":   "firing",
		"check": map[string]string{
			"nested": "true",
		},
	}

	alertInfo, err := cli.SystemAlert(dd)
	if err != nil {
		golog.Error("Error while getting alert info: ", err)
	}
	golog.Info("alert info: ", alertInfo)
}
