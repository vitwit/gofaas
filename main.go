package main

import (
	"github.com/kataras/golog"
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

}
