package main

import (
	"github.com/kataras/golog"
	"github.com/vitwit/go-fass/openfass-go/models"
	"github.com/vitwit/go-fass/openfass-go"
)

func main() {
	cli := openfass_go.NewSendClient("admin", "9a651783248f8b61a92ae2fa02c7cb627a68e061f5cc0773822f72e0c7af077d",
		"")

	golog.Info("hostt.......", cli.BaseURL, cli.Method)

	res, err := cli.GetSystemFunctions("admin", "9a651783248f8b61a92ae2fa02c7cb627a68e061f5cc0773822f72e0c7af077d")

	if err != nil {
		golog.Error("here...", err)
	}

	golog.Info("response og get system...........", res)

	data := models.FunctionDefintion{
		Service:    "nodeinfo123",
		Network:    "func_functions",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
		EnvVars: models.EnvVars{
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
		Annotations: models.Annotations{
			Topics: "awesome-kafka-topic",
			Foo:    "some",
		},
		RegistryAuth: "dXNlcjpwYXNzd29yZA==",
		Limits: models.Limits{
			Memory: "128M",
			CPU:    "0.01",
		},
		Requests: models.Requests{
			Memory: "128M",
			CPU:    "0.01",
		},
		ReadOnlyRootFilesystem: true,
	}

	clires, err := cli.CreateSystemFunctions(data)
	if err != nil {
		golog.Error("Error while creating system ", err)
	}

	golog.Info("response of create system....", clires)

	dd := map[string]string{
		"receiver": "scale-up",
		"status": "firing",
		"alertName":"check",
	}

	rr, _ := cli.SystemAlert(dd)
	golog.Info("alert info.......",rr)

	//corsRouter := cors.New(corsRouter11).Handler(router)
	//log.Fatal(http.ListenAndServe(":3000", nil))
}
