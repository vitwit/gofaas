package main

import (
	"github.com/kataras/golog"
	"github.com/vitwit/go-fass/openfass-go"
	"os"
)

func main() {
	cli := openfass_go.NewClient(os.Getenv("USER"), os.Getenv("PASSWORD"),
		"")

	golog.Info("host: ", cli.BaseURL)

	res, err := cli.GetSystemFunctions()

	if err != nil {
		golog.Error("Error from system fucntions:  ", err)
	}

	golog.Info("response of get system:  ", res)

	data := openfass_go.FunctionDefintion{
		Service:    "nodeinfo123",
		Network:    "func_functions",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
		EnvVars: openfass_go.EnvVars{
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
		Annotations: openfass_go.Annotations{
			Topics: "awesome-kafka-topic",
			Foo:    "some",
		},
		RegistryAuth: "dXNlcjpwYXNzd29yZA==",
		Limits: openfass_go.Limits{
			Memory: "128M",
			CPU:    "0.01",
		},
		Requests: openfass_go.Requests{
			Memory: "128M",
			CPU:    "0.01",
		},
		ReadOnlyRootFilesystem: true,
	}

	clires, err := cli.CreateSystemFunctions(data)
	if err != nil {
		golog.Error("Error while creating system:  ", err)
	}

	golog.Info("response of create system: ", clires)

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
