package go_faas

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestCreateSystemFunctions(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		FunctionDef   *FunctionDefintion
		StatusCode    int
		ErrorExpected bool
	}{
		{
			Name: "Function/nodeinfo123456",
			FunctionDef: &FunctionDefintion{
				Service:    "nodeinfo123456",
				Network:    "func_functions",
				Image:      "functions/nodeinfo:latest",
				EnvProcess: "node main.js",
				EnvVars: EnvVars{
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
				Annotations: Annotations{
					Topics: "awesome-kafka-topic",
					Foo:    "some",
				},
				RegistryAuth: "dXNlcjpwYXNzd29yZA==",
				Limits: Limits{
					Memory: "128M",
					CPU:    "0.01",
				},
				Requests: Requests{
					Memory: "128M",
					CPU:    "0.01",
				},
				ReadOnlyRootFilesystem: true,
			},
			StatusCode:    202,
			ErrorExpected: false,
		},
		{
			Name: "Function/yetanothernodeinfo",
			FunctionDef: &FunctionDefintion{
				Service:    "yetanothernodeinfo",
				Network:    "func_functions",
				Image:      "functions/nodeinfo:latest",
				EnvProcess: "node main.js",
				Constraints: []string{
					"node.platform.os == linux",
				},
				Labels: map[string]string{
					"label": "val",
				},
				RegistryAuth: "dXNlcjpwYXNzd29yZA==",
				Limits: Limits{
					Memory: "128M",
					CPU:    "0.01",
				},
				Requests: Requests{
					Memory: "128M",
					CPU:    "0.01",
				},
			},
			StatusCode:    202,
			ErrorExpected: false,
		},
		{
			Name: "Function/InvalidService",
			FunctionDef: &FunctionDefintion{
				Service:      "InvalidService",
				Network:      "func_functions",
				Image:        "",
				EnvProcess:   "node main.js",
				RegistryAuth: "dXNlcjpwYXNzd29yZA==",
				Limits: Limits{
					Memory: "128M",
					CPU:    "0.01",
				},
				Requests: Requests{
					Memory: "128M",
					CPU:    "0.01",
				},
			},
			StatusCode:    400,
			ErrorExpected: true,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.CreateSystemFunctions(c.FunctionDef)
			if err != nil {
				t.Errorf("error while calling %s/system/functions: %v", os.Getenv("OPENFAAS_GATEWAY_ADDR"), err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestGetSystemFunctions(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("GetSystemFunctions", func(t *testing.T) {
		resp, err := cli.GetSystemFunctions()
		if err != nil {
			t.Errorf("error while getting system functions: %v", err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	})
}

func TestUpdateSystemFunctions(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		FunctionDef   *FunctionDefintion
		ExpectedError bool
		StatusCode    int
	}{
		{
			Name: "nodeinfo123456/UpdateLabel",
			FunctionDef: &FunctionDefintion{
				Service: "nodeinfo123456",
				Image:   "functions/nodeinfo:latest",
				Labels: map[string]string{
					"changedlabelkey": "changedlabelval",
				},
			},
			ExpectedError: false,
			StatusCode:    202,
		},
		//{
		//	Name: "nodeinfo123456/UpdateLimits",
		//	FunctionDef: &FunctionDefintion{
		//		Service:    "nodeinfo123456",
		//		Image:      "functions/nodeinfo:latest",
		//		Limits: Limits{
		//			Memory: "130M",
		//			CPU:    "0.01",
		//		},
		//	},
		//	ExpectedError: false,
		//	StatusCode: 202,
		//},
		{
			Name: "nodeinfo123456/NoImage",
			FunctionDef: &FunctionDefintion{
				Service: "nodeinfo123456",
				Labels: map[string]string{
					"changedlabelkey": "changedlabelval",
				},
			},
			ExpectedError: true,
			StatusCode:    400,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.UpdateSystemFunctions(c.FunctionDef)
			if err != nil {
				t.Fatalf("error while calling [PUT] %s/system/functions: %v", os.Getenv("OPENFAAS_GATEWAY_ADDR"), err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestDeleteSystemFunctions(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		DeleteRequest *DeleteFunctionBodyOpts
		ErrorExpected bool
		StatusCode    int
	}{
		{
			Name:          "yetanothernodeinfo",
			DeleteRequest: &DeleteFunctionBodyOpts{FunctionName: "yetanothernodeinfo"},
			ErrorExpected: false,
			StatusCode:    202,
		},
		{
			Name:          "InvalidFunc",
			DeleteRequest: &DeleteFunctionBodyOpts{FunctionName: "notexists"},
			ErrorExpected: true,
			StatusCode:    404,
		},
		{
			Name:          "NoFunc",
			DeleteRequest: &DeleteFunctionBodyOpts{},
			ErrorExpected: true,
			StatusCode:    400,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.DeleteSystemFunction(c.DeleteRequest)
			if err != nil {
				t.Errorf("error while calling [DELETE] %s/system/functions: %v", os.Getenv("OPENFAAS_GATEWAY_ADDR"), err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestSystemAlert(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		AlertDef      *SystemAlertBodyOpts
		ExpectedError bool
		StatusCode    int
	}{
		{
			Name: "APIHighInvocationRate",
			AlertDef: &SystemAlertBodyOpts{
				Receiver: "scale-up",
				Status:   "firing",
				Alerts: []SystemAlertsStruct{
					{
						Status: "firing",
						Labels: SystemAlertLables{
							Alertname:    "APIHighInvocationRate",
							Code:         "200",
							FunctionName: "func_nodeinfo",
							Instance:     os.Getenv("OPENFAAS_GATEWAY_ADDR"),
							Job:          "gateway",
							Monitor:      "faas-monitor",
							Service:      "gateway",
							Severity:     "major",
							Value:        "8.998200359928017",
						},
						Annotations: SystemAlertAnnotations{
							Description: "High invocation total on gateway:8080",
							Summary:     "High invocation total on gateway:8080",
						},
						StartsAt: time.Now(),
						EndsAt:   time.Now().Add(time.Hour * 24),
					},
				},
				GroupLabels: GroupLabels{
					Alertname: "APIHighInvocationRate",
					Service:   "gateway",
				},
				CommonLabels: CommonLabels{
					Alertname:    "APIHighInvocationRate",
					Code:         "200",
					FunctionName: "func_nodeinfo",
					Instance:     os.Getenv("OPENFAAS_GATEWAY_ADDR"),
					Job:          "gateway",
					Monitor:      "faas-monitor",
					Service:      "gateway",
					Severity:     "major",
					Value:        "8.998200359928017",
				},
				CommonAnnotations: CommonAnnotations{
					Description: "High invocation total on gateway:8080",
					Summary:     "High invocation total on gateway:8080",
				},
				ExternalURL: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
				Version:     "3",
			},
			ExpectedError: false,
			StatusCode:    200,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.SystemAlert(c.AlertDef)
			if err != nil {
				t.Errorf("error while setting system alert: %v", err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestAsyncFunction(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		Data          interface{}
		ErrorExpected bool
		StatusCode    int
		FuncName      string
		CallbackURL   string
		Desc          string
	}{
		{
			Name: "AsyncInvokeWithJSONBody",
			Data: map[string]string{
				"hello": "world",
			},
			ErrorExpected: false,
			StatusCode:    202,
			FuncName:      "nodeinfo",
			CallbackURL:   "",
			Desc:          "Send json body for nodeinfo func async invocation",
		},
		{
			Name:          "AsyncInvokeWithStringBody",
			Data:          "Testing func_nodeinfo",
			ErrorExpected: false,
			StatusCode:    202,
			FuncName:      "nodeinfo",
			CallbackURL:   "",
			Desc:          "Send string body for nodeinfo func async invocation",
		},
		{
			Name:          "AsyncInvokeInvalidFunc",
			Data:          "Testing invalid func",
			ErrorExpected: false,
			StatusCode:    202,
			FuncName:      "somethingThatDoesNotExist",
			CallbackURL:   "",
			Desc:          "Invoke func which does not exists asynchronously",
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.AsyncFunction(&AsyncInvocationOpts{
				Body:         c.Data,
				FunctionName: c.FuncName,
				CallbackURL:  c.CallbackURL,
			})
			if err != nil {
				t.Errorf("error while asynchronously invoking func %s: %v", c.FuncName, err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestInvokeFunction(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		Data          interface{}
		ErrorExpected bool
		StatusCode    int
		FuncName      string
		Desc          string
	}{
		{
			Name: "InvokeWithJSONBody",
			Data: map[string]string{
				"hello": "world",
			},
			ErrorExpected: false,
			StatusCode:    200,
			FuncName:      "nodeinfo123456",
			Desc:          "Send json for nodeinfo func invocation",
		},
		{
			Name:          "InvokeWithStringBody",
			Data:          "Testing func_nodeinfo",
			ErrorExpected: false,
			StatusCode:    200,
			FuncName:      "nodeinfo123456",
			Desc:          "Send string for nodeinfo func invocation",
		},
		{
			Name:          "InvokeInvalidFunc",
			Data:          "Testing invalid func",
			ErrorExpected: true,
			StatusCode:    404,
			FuncName:      "somethingThatDoesNotExist",
			Desc:          "Invoke func which does not exists",
		},
	}

	for _, c := range testcases {
		resp, err := cli.InvokeFunction(&SyncInvocationOpts{
			Body:         c.Data,
			FunctionName: c.FuncName,
		})
		if err != nil {
			t.Errorf("error while invoking func %s: %v", c.FuncName, err)
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, c.StatusCode, resp.StatusCode)
	}
}

func TestScaleFunction(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		Body          *ScaleFunctionBodyOpts
		ErrorExpected bool
		StatusCode    int
	}{
		{
			Name: "nodeinfo/ScaleUp/2",
			Body: &ScaleFunctionBodyOpts{
				Service:  "nodeinfo",
				Replicas: 2,
			},
			ErrorExpected: false,
			StatusCode:    202,
		},
		{
			Name: "nodeinfo/ScaleDown/1",
			Body: &ScaleFunctionBodyOpts{
				Service:  "nodeinfo",
				Replicas: 1,
			},
			ErrorExpected: false,
			StatusCode:    202,
		},
		{
			Name: "InvalidFunc/ScaleUp/2",
			Body: &ScaleFunctionBodyOpts{
				Service:  "InvalidFunc",
				Replicas: 2,
			},
			ErrorExpected: true,
			StatusCode:    500,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.ScaleFunction(c.Body)
			if err != nil {
				t.Errorf("error while scaling func %s: %v", c.Body.Service, err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestGetFunctionSummary(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		FuncName      string
		ExpectedError bool
		StatusCode    int
	}{
		{
			Name:          "nodeinfo",
			FuncName:      "nodeinfo",
			ExpectedError: false,
			StatusCode:    200,
		},
		{
			Name:          "InvalidFunc",
			FuncName:      "InvalidFunc",
			ExpectedError: true,
			StatusCode:    404,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.GetFunctionSummary(c.FuncName)
			if err != nil {
				t.Errorf("error while getting summary for func %s: %v", c.FuncName, err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestCreateNewSecret(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		Body          *SecretBodyOpts
		ExpectedError bool
		StatusCode    int
	}{
		{
			Name: "secretkey101",
			Body: &SecretBodyOpts{
				Name:  "secretkey101",
				Value: "secretval101",
			},
			ExpectedError: false,
			StatusCode:    201,
		},
		{
			Name: "secretkey102",
			Body: &SecretBodyOpts{
				Name:  "secretkey102",
				Value: "secretval102",
			},
			ExpectedError: false,
			StatusCode:    201,
		},
		{
			Name: "secretkey103",
			Body: &SecretBodyOpts{
				Name:  "secretkey103",
				Value: "secretval103",
			},
			ExpectedError: false,
			StatusCode:    201,
		},
		{
			Name: "<NoSecretName>",
			Body: &SecretBodyOpts{
				Value: "secretval101",
			},
			ExpectedError: true,
			StatusCode:    500,
		},
		{
			Name: "<NoSecretVal>",
			Body: &SecretBodyOpts{
				Name: "secretkey101",
			},
			ExpectedError: true,
			StatusCode:    500,
		},
		{
			Name:          "<NoSecretBody>",
			Body:          &SecretBodyOpts{},
			ExpectedError: true,
			StatusCode:    500,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.CreateNewSecret(c.Body)
			if err != nil {
				t.Errorf("error creating secret: %v", err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestGetSecrets(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("List", func(t *testing.T) {
		resp, err := cli.GetSecrets()
		if err != nil {
			t.Errorf("error getting secrets list: %v", err)
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, resp.StatusCode, 200)
	})
}

func TestUpdateSecret(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	if cli.ClusterType == "swarm" {
		t.Skip("Cluster swarm does not have an update secret method")
	}

	testcases := []struct {
		Name          string
		Body          *SecretBodyOpts
		ExpectedError bool
		StatusCode    int
	}{
		{
			Name: "secretkey102",
			Body: &SecretBodyOpts{
				Name:  "secretkey102",
				Value: "updatedsecretval102",
			},
			ExpectedError: false,
			StatusCode:    200,
		},
		{
			Name: "secretkey103",
			Body: &SecretBodyOpts{
				Name: "secretkey103",
			},
			ExpectedError: false,
			StatusCode:    200,
		},
		{
			Name:          "<NoSecretBody>",
			Body:          &SecretBodyOpts{},
			ExpectedError: true,
			StatusCode:    404,
		},
		{
			Name: "<NoSecretName>",
			Body: &SecretBodyOpts{
				Value: "updatedsecretval101",
			},
			ExpectedError: true,
			StatusCode:    404,
		},
		{
			Name: "<InvalidSecret>",
			Body: &SecretBodyOpts{
				Name:  "InvalidSecret",
				Value: "updatedsecretval101",
			},
			ExpectedError: true,
			StatusCode:    404,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.UpdateSecret(c.Body)
			if err != nil {
				t.Errorf("error updating secret: %v", err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestDeleteSecret(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		Body          *SecretNameBodyOpts
		ExpectedError bool
		StatusCode    int
	}{
		{
			Name: "secretkey101",
			Body: &SecretNameBodyOpts{
				Name: "secretkey101",
			},
			ExpectedError: false,
			StatusCode:    200,
		},
		{
			Name:          "<NoSecretBody>",
			Body:          &SecretNameBodyOpts{},
			ExpectedError: true,
			StatusCode:    404,
		},
		{
			Name: "<InvalidSecret>",
			Body: &SecretNameBodyOpts{
				Name: "DoesNotExist",
			},
			ExpectedError: true,
			StatusCode:    404,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.DeleteSecret(c.Body)
			if err != nil {
				t.Errorf("error deleting secret: %v", err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestGetSystemLogs(t *testing.T) {
	t.Skip() // unable to parse "Since" TODO: Check parsing query string
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name          string
		Query         *SystemLogsQueryOpts
		ExpectedError bool
		StatusCode    int
	}{
		{
			Name: "nodeinfo/Tail2/FollowFalse",
			Query: &SystemLogsQueryOpts{
				Name:   "nodeinfo",
				Tail:   2,
				Follow: false,
			},
			ExpectedError: false,
			StatusCode:    200,
		},
		{
			Name: "nodeinfo/Tail5/Since2020-01-22T07:48:18+00:00",
			Query: &SystemLogsQueryOpts{
				Name:  "nodeinfo",
				Tail:  5,
				Since: "2020-01-22T07:48:18+00:00",
			},
			ExpectedError: false,
			StatusCode:    200,
		},
		{
			Name: "InvalidFunc/Tail2/FollowFalse",
			Query: &SystemLogsQueryOpts{
				Name:   "InvalidFunc",
				Tail:   2,
				Follow: false,
			},
			ExpectedError: true,
			StatusCode:    500,
		},
	}

	for _, c := range testcases {
		t.Run(c.Name, func(t *testing.T) {
			resp, err := cli.GetSystemLogs(c.Query)
			if err != nil {
				t.Errorf("error getting system logs: %v", err)
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, c.StatusCode, resp.StatusCode)
		})
	}
}

func TestGetSystemInfo(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("GetSystemInfo", func(t *testing.T) {
		resp, err := cli.GetSystemInfo()
		if err != nil {
			t.Errorf("error getting system logs: %v", err)
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

func TestGetHealthz(t *testing.T) {
	cli, err := NewClient(&FaasGatewayCredentials{
		Username:       os.Getenv("OPENFAAS_USERNAME"),
		Password:       os.Getenv("OPENFAAS_PASSWORD"),
		GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
		ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("GetHealthz", func(t *testing.T) {
		resp, err := cli.GetHealthz()
		if err != nil {
			t.Errorf("error getting system logs: %v", err)
		}
		assert.Equal(t, nil, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
