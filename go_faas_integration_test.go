// +build integration

package go_faas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func (suite *GoFaasTestSuite) TestFunctionsCRUD() {
	funcName := "integration_testnodeinfo"

	// get functions
	suite.T().Run("IntTests/TestFunctionsCRUD/GetSystemFunctions", func(t *testing.T) {
		resp, err := suite.cli.GetSystemFunctions()
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 200, resp.StatusCode)
	})

	// create a func
	def := &FunctionDefintion{
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
	}
	suite.T().Run("IntTests/TestFunctionsCRUD/CreateSystemFunctions", func(t *testing.T) {
		resp, err = suite.cli.CreateSystemFunctions(def)
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 202, resp.StatusCode)
	})

	// get function summary
	suite.T().Run("IntTests/TestFunctionsCRUD/GetFunctionSummary", func(t *testing.T) {
		resp, err = suite.cli.GetFunctionSummary(funcName)
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 200, resp.StatusCode)
	})

	// update the func
	update := &FunctionDefintion{
		Service: funcName,
		Image:   "functions/nodeinfo:latest",
		Limits: Limits{
			Memory: "130M",
			CPU:    "0.01",
		},
	}
	suite.T().Run("IntTests/TestFunctionsCRUD/UpdateSystemFunctions", func(t *testing.T) {
		resp, err = suite.cli.UpdateSystemFunctions(update)
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 202, resp.StatusCode)
	})

	// scale up the func
	suite.T().Run("IntTests/TestFunctionsCRUD/ScaleFunction/Up", func(t *testing.T) {
		resp, err = suite.cli.ScaleFunction(&ScaleFunctionBodyOpts{
			Service:  funcName,
			Replicas: 3,
		})
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 202, resp.StatusCode)
	})

	// scale down the func
	suite.T().Run("IntTests/TestFunctionsCRUD/ScaleFunction/Down", func(t *testing.T) {
		resp, err = suite.cli.ScaleFunction(&ScaleFunctionBodyOpts{
			Service:  funcName,
			Replicas: 1,
		})
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 202, resp.StatusCode)
	})

	// delete the func
	suite.T().Run("IntTests/TestFunctionsCRUD/DeleteSystemFunction", func(t *testing.T) {
		resp, err = suite.cli.DeleteSystemFunction(&DeleteFunctionBodyOpts{FunctionName: funcName})
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 202, resp.StatusCode)
	})
}

func (suite *GoFaasTestSuite) TestSecretsCRUD() {
	secretName := "integration_test_secret"
	secretVal := "integration_test_secret_val"

	//get secrets list
	suite.T().Run("IntTests/TestSecretsCRUD/GetSecrets", func(t *testing.T) {
		resp, err := suite.cli.GetSecrets()
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), resp.StatusCode, 200)
	})

	// create new secret
	suite.T().Run("IntTests/TestSecretsCRUD/CreateNewSecret", func(t *testing.T) {
		resp, err = suite.cli.CreateNewSecret(&SecretBodyOpts{
			Name:  secretName,
			Value: secretVal,
		})
	})
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 201, resp.StatusCode)

	// update the secret if not a swarm cluster
	if suite.cli.ClusterType != "swarm" {
		suite.T().Run("IntTests/TestSecretsCRUD/UpdateSecret", func(t *testing.T) {
			resp, err = suite.cli.UpdateSecret(&SecretBodyOpts{
				Name:  secretName,
				Value: "updated_integration_test_secret_val",
			})
		})
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 200, resp.StatusCode)
	}

	// delete the secret
	suite.T().Run("IntTests/TestSecretsCRUD/DeleteSecret", func(t *testing.T) {
		resp, err = suite.cli.DeleteSecret(&SecretNameBodyOpts{Name: secretName})
		assert.Equal(suite.T(), nil, err)
		assert.Equal(suite.T(), 200, resp.StatusCode)
	})
}
