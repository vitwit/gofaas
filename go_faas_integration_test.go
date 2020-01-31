// +build integration

package go_faas

import (
	"github.com/stretchr/testify/assert"
	"github.com/vitwit/go-faas/sdk"
)

func (suite *GoFaasTestSuite) TestFunctionsCRUD() {
	funcName := "integration_testnodeinfo"

	// get functions
	resp, err := suite.cli.GetSystemFunctions()
	if err != nil {
		suite.T().Fatalf("GetSystemFunctions/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 200, resp.StatusCode)

	// create a func
	def := &sdk.FunctionDefintion{
		Service:    funcName,
		Network:    "func_functions",
		Image:      "functions/nodeinfo:latest",
		EnvProcess: "node main.js",
		Constraints: []string{
			"node.platform.os == linux",
		},
		Labels: map[string]string{
			"labelkey": "labelval",
		},
		Annotations: sdk.Annotations{
			Topics: "awesome-kafka-topic",
			Foo:    "some",
		},
		RegistryAuth: "dXNlcjpwYXNzd29yZA==",
		Limits: sdk.Limits{
			Memory: "128M",
			CPU:    "0.01",
		},
		Requests: sdk.Requests{
			Memory: "128M",
			CPU:    "0.01",
		},
		ReadOnlyRootFilesystem: true,
	}
	resp, err = suite.cli.CreateSystemFunctions(def)
	if err != nil {
		suite.T().Fatalf("CreateSystemFunctions/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 202, resp.StatusCode)

	// get function summary
	resp, err = suite.cli.GetFunctionSummary(funcName)
	if err != nil {
		suite.T().Fatalf("CreateSystemFunctions/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 200, resp.StatusCode)

	// update the func
	update := &sdk.FunctionDefintion{
		Service: funcName,
		Image:   "functions/nodeinfo:latest",
		Limits: sdk.Limits{
			Memory: "130M",
			CPU:    "0.01",
		},
	}
	resp, err = suite.cli.UpdateSystemFunctions(update)
	if err != nil {
		suite.T().Fatalf("UpdateSystemFunctions/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 202, resp.StatusCode)

	// scale up the func
	resp, err = suite.cli.ScaleFunction(&sdk.ScaleFunctionBodyOpts{
		Service:  funcName,
		Replicas: 3,
	})
	if err != nil {
		suite.T().Fatalf("ScaleFunction/Up/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 202, resp.StatusCode)

	// scale down the func
	resp, err = suite.cli.ScaleFunction(&sdk.ScaleFunctionBodyOpts{
		Service:  funcName,
		Replicas: 1,
	})
	if err != nil {
		suite.T().Fatalf("ScaleFunction/Down/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 202, resp.StatusCode)

	// delete the func
	resp, err = suite.cli.DeleteSystemFunction(&sdk.DeleteFunctionBodyOpts{FunctionName: funcName})
	if err != nil {
		suite.T().Fatalf("DeleteSystemFunction/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 202, resp.StatusCode)
}

func (suite *GoFaasTestSuite) TestSecretsCRUD() {
	secretName := "integration_test_secret"
	secretVal := "integration_test_secret_val"

	//get secrets list
	resp, err := suite.cli.GetSecrets()
	if err != nil {
		suite.T().Errorf("GetSecrets/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), resp.StatusCode, 200)

	// create new secret
	resp, err = suite.cli.CreateNewSecret(&sdk.SecretBodyOpts{
		Name:  secretName,
		Value: secretVal,
	})
	if err != nil {
		suite.T().Fatalf("CreateNewSecret/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 201, resp.StatusCode)

	// update the secret if not a swarm cluster
	if suite.cli.ClusterType != "swarm" {
		resp, err = suite.cli.UpdateSecret(&sdk.SecretBodyOpts{
			Name:  secretName,
			Value: "updated_integration_test_secret_val",
		})
		if err != nil {
			suite.T().Fatalf("UpdateSecret/Error: %v", err)
		}
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 200, resp.StatusCode)
	}

	// delete the secret
	resp, err = suite.cli.DeleteSecret(&sdk.SecretNameBodyOpts{Name: secretName})
	if err != nil {
		suite.T().Fatalf("DeleteSecret/Error: %v", err)
	}
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 200, resp.StatusCode)
}
