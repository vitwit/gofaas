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

	//req := openfass_go.GetRequest( "/system/functions", "","admin",
	//	"9a651783248f8b61a92ae2fa02c7cb627a68e061f5cc0773822f72e0c7af077d")
	//req.Method = "GET"
	//
	//
	//golog.Info("reqqqqqq...........", req.Headers)
	//res, err := openfass_go.MakeRequest(req)
	if err != nil {
		golog.Error("here...", err)
	}

	golog.Info("res...........", res)

	//res1, err := cli.GetSectrets()
	//if err != nil {
	//	golog.Error("errrrrr......", err)
	//}
	//
	//golog.Info("respppp11111111.........", len(res1.Body))

	//url := "http://127.0.0.1:8080/system/functions"
	//
	//req, _ := http.NewRequest("GET", url, nil)
	//
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", "Basic admin:9a651783248f8b61a92ae2fa02c7cb627a68e061f5cc0773822f72e0c7af077d\"")
	//req.Header.Add("Cache-Control", "no-cache")
	//req.Header.Add("Postman-Token", "c1768a59-558a-4588-bb54-3b3c52ac74ee")
	//
	//res, _ := http.DefaultClient.Do(req)
	//
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(res)
	//fmt.Println(string(body))
}
