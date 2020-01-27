
# go-faas
Go SDK for OpenFaaS
This documentation is under (WIP)

**This library allows you to quickly and easily use the openFasS API**

# How to use

<a name="installation"></a>
# Installation

## Prerequisites

- go
- Installed openFasS

### Install Package

```sh
	go get github.com/vitwit/go-faas
```


## Setup Environment Variables

### Initial Setup

```bash
cp .env.example .env
```
Update OPENFAAS_GATEWAY_ADDR, OPENFAAS_USER, OPENFAAS_PASSWORD, OPENFAAS_CLUSTER_TYPE

### Environment Variable

Update the development environment with your keys, for example:

```bash
$ export OPENFAAS_GATEWAY_ADDR=http://127.0.0.1:8080
$ export OPENFAAS_USER='admin'
$ export OPENFAAS_PASSWORD='password'
$ export OPENFAAS_CLUSTER_TYPE='swarm/kubernetes'
```
<a name="quick-start"></a>
# Quick Start

### Create a client

```
import faas "github.com/vitwit/go-faas";

client := faas.NewClient(&faas.FaasGatewayCredentials{
    Username:       os.Getenv("OPENFAAS_USERNAME"),
    Password:       os.Getenv("OPENFAAS_PASSWORD"),
    GatewayAddress: os.Getenv("OPENFAAS_GATEWAY_ADDR"),
    ClusterType:    os.Getenv("OPENFAAS_CLUSTER_TYPE"),
})
```

<details>

<summary>getFunctions</summary>

getFunctions
---
 **Example**

 ```
 resp, err := cli.GetSystemFunctions()
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "List of deployed functions.",
    "schema": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/FunctionListEntry"
      }
    }
  }
}
```

######  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>createFunction</summary>

createFunction
---
 **Example**

 ```
funcdef := &faas.FunctionDefintion{
    Service:    "nodeinfo12345",
    Network:    "func_functions",
    Image:      "functions/nodeinfo:latest",
    EnvProcess: "node main.js",
    EnvVars: faas.EnvVars{
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
    Annotations: faas.Annotations{
        Topics: "awesome-kafka-topic",
        Foo:    "some",
    },
    RegistryAuth: "dXNlcjpwYXNzd29yZA==",
    Limits: faas.Limits{
        Memory: "128M",
        CPU:    "0.01",
    },
    Requests: faas.Requests{
        Memory: "128M",
        CPU:    "0.01",
    },
    ReadOnlyRootFilesystem: true,
}

resp, err := cli.CreateSystemFunctions(funcdef)
```
**Responses**


> Success 2XX
```json
{
  "202": {
    "description": "Accepted"
  }
}
```

> Error 4XX
```json
{
  "400": {
    "description": "Bad Request"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>updateFunction</summary>

updateFunction
---
 **Example**

 ```
update := &faas.FunctionDefintion{
    Service: "nodeinfo",
    Image:   "functions/nodeinfo:latest",
    Labels: map[string]string{
        "changedlabelkey": "changedlabelval",
    },
},
resp, err := cli.UpdateSystemFunctions(update)
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Accepted"
  }
}
```

> Error 4XX
```json
{
  "400": {
    "description": "Bad Request"
  },
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>deleteFunction</summary>

deleteFunction
---
 **Example**

 ```
resp, err := cli.DeleteSystemFunction(&faas.DeleteFunctionBodyOpts{FunctionName: "nodeinfo"})
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "OK"
  }
}
```

> Error 4XX
```json
{
  "400": {
    "description": "Bad Request"
  },
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>handleAlert</summary>

handleAlert
---
 **Example**

 ```
sysalert := &faas.SystemAlertBodyOpts{
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
}
resp, err := cli.SystemAlert(syslert)
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Alert handled successfully"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal error with swarm or request JSON invalid"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>invokeFunctionAsync</summary>

invokeFunctionAsync
---
 **Example**

```
resp, err := cli.AsyncFunction(&faas.AsyncInvocationOpts{
    Body:         "Hey there!",
    FunctionName: "nodeinfo",
    CallbackURL:  "http://localhost:5000",
})
```
**Responses**


> Success 2XX
```json
{
  "202": {
    "description": "Request accepted and queued"
  }
}
```

> Error 4XX
```json
{
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>invokeFunction</summary>

invokeFunction
---
 **Example**

 ```
resp, err := cli.InvokeFunction(&faas.SyncInvocationOpts{
    Body:         "Hi there!",
    FunctionName: "nodeinfo",
})
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Value returned from function"
  }
}
```

> Error 4XX
```json
{
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal server error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>scaleFunction</summary>

scaleFunction
---
 **Example**

```
resp, err := cli.ScaleFunction(&faas.ScaleFunctionBodyOpts{
    Service:  "nodeinfo",
    Replicas: 2,
})
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Scaling OK"
  },
  "202": {
    "description": "Scaling OK"
  }
}
```

> Error 4XX
```json
{
  "404": {
    "description": "Function not found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Error scaling function"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>getFunctionSummary</summary>

getFunctionSummary
---
 **Example**

 ```
resp, err := cli.GetFunctionSummary("nodeinfo")
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Function Summary",
    "schema": {
      "$ref": "#/definitions/FunctionListEntry"
    }
  }
}
```

> Error 4XX
```json
{
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)
</details>

<details>

<summary>getSecrets</summary>

getSecrets
---
 **Example**

 ```
resp, err := client.getSecrets()
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "List of submitted secrets.",
    "schema": {
      "$ref": "#/definitions/SecretName"
    }
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [SecretName](###SecretName-modal)
</details>

<details>

<summary>createSecret</summary>

createSecret
---
 **Example**

```
resp, err := cli.CreateNewSecret(&faas.SecretBodyOpts{
    Name:  "key",
    Value: "val",
})
```
**Responses**


> Success 2XX
```json
{
  "201": {
    "description": "Created"
  }
}
```

> Error 4XX
```json
{
  "400": {
    "description": "Bad Request"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [Secret](###Secret-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [SecretName](###SecretName-modal)
</details>

<details>

<summary>updateSecret</summary>

updateSecret
---
 **Example**

```
resp, err := cli.UpdateSecret(&faas.SecretBodyOpts{
    Name:  "key",
    Value: "updatedval",
})
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Ok"
  }
}
```

> Error 4XX
```json
{
  "400": {
    "description": "Bad Request"
  },
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [Secret](###Secret-modal)  [Secret](###Secret-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [SecretName](###SecretName-modal)
</details>

<details>

<summary>deleteSecret</summary>

deleteSecret
---
 **Example**

 ```
resp, err := cli.DeleteSecret(&faas.SecretNameBodyOpts{
    Name: "key",
})
```
**Responses**


> Success 2XX
```json
{
  "204": {
    "description": "OK"
  }
}
```

> Error 4XX
```json
{
  "400": {
    "description": "Bad Request"
  },
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [Secret](###Secret-modal)  [Secret](###Secret-modal)  [SecretName](###SecretName-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [SecretName](###SecretName-modal)
</details>

<details>

<summary>getLogsOfAFunction</summary>

getLogsOfAFunction
---
 **Example**

 ```
resp, err := cli.GetSystemLogs(&faas.SystemLogsQueryOpts{
    Name:   "nodeinfo",
    Tail:   10,
    Follow: false,
    Since: "2020-01-22T07:48:18+00:00"
})
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Newline delimited stream of log messages",
    "schema": {
      "$ref": "#/definitions/LogEntry"
    }
  }
}
```

> Error 4XX
```json
{
  "404": {
    "description": "Not Found"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [Secret](###Secret-modal)  [Secret](###Secret-modal)  [SecretName](###SecretName-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [SecretName](###SecretName-modal)  [LogEntry](###LogEntry-modal)
</details>

<details>

<summary>getInfo</summary>

getInfo
---
 **Example**

 ```
resp, err := cli.GetSystemInfo()
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Info result",
    "schema": {
      "$ref": "#/definitions/Info"
    }
  }
}
```

> Error 4XX
```json
{
  "404": {
    "description": "Provider does not support info endpoint"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Internal Server Error"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [Secret](###Secret-modal)  [Secret](###Secret-modal)  [SecretName](###SecretName-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [SecretName](###SecretName-modal)  [LogEntry](###LogEntry-modal)  [Info](###Info-modal)
</details>

<details>

<summary>checkHealth</summary>

checkHealth
---
 **Example**

 ```
resp, err := cli.GetHealthz()
```
**Responses**


> Success 2XX
```json
{
  "200": {
    "description": "Healthy"
  }
}
```

> Error 5XX
```json
{
  "500": {
    "description": "Not healthy"
  }
}
```

######  [FunctionDefintion](###FunctionDefintion-modal)  [FunctionDefintion](###FunctionDefintion-modal)  [DeleteFunctionRequest](###DeleteFunctionRequest-modal)  [undefined](###undefined-modal)  [Secret](###Secret-modal)  [Secret](###Secret-modal)  [SecretName](###SecretName-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [FunctionListEntry](###FunctionListEntry-modal)  [SecretName](###SecretName-modal)  [LogEntry](###LogEntry-modal)  [Info](###Info-modal)
</details>

# Modal Definations

 ### Info-modal
 ```json
{
  "type": "object",
  "properties": {
    "provider": {
      "type": "object",
      "description": "The OpenFaaS Provider",
      "properties": {
        "provider": {
          "type": "string",
          "example": "faas-swarm"
        },
        "orchestration": {
          "type": "string",
          "example": "swarm"
        },
        "version": {
          "type": "object",
          "description": "Version of the OpenFaaS Provider",
          "properties": {
            "commit_message": {
              "type": "string",
              "example": "Sample Message"
            },
            "sha": {
              "type": "string",
              "example": "7108418d9dd6b329ddff40e7393b3166f8160a88"
            },
            "release": {
              "type": "string",
              "format": "semver",
              "example": "0.2.6"
            }
          }
        }
      }
    },
    "version": {
      "type": "object",
      "description": "Version of the Gateway",
      "properties": {
        "commit_message": {
          "type": "string",
          "example": "Sample Message"
        },
        "sha": {
          "type": "string",
          "example": "7108418d9dd6b329ddff40e7393b3166f8160a88"
        },
        "release": {
          "type": "string",
          "format": "semver",
          "example": "0.8.9"
        }
      }
    },
    "arch": {
      "type": "string",
      "description": "Platform architecture",
      "example": "x86_64"
    }
  },
  "required": [
    "provider",
    "version"
  ]
}
```

 ### DeleteFunctionRequest-modal
 ```json
{
  "type": "object",
  "properties": {
    "functionName": {
      "type": "string",
      "description": "Name of deployed function",
      "example": "nodeinfo"
    }
  },
  "required": [
    "functionName"
  ]
}
```

 ### FunctionDefintion-modal
 ```json
{
  "type": "object",
  "properties": {
    "service": {
      "type": "string",
      "description": "Name of deployed function",
      "example": "nodeinfo"
    },
    "network": {
      "type": "string",
      "description": "Docker swarm network, usually func_functions",
      "example": "func_functions"
    },
    "image": {
      "type": "string",
      "description": "Docker image in accessible registry",
      "example": "functions/nodeinfo:latest"
    },
    "envProcess": {
      "type": "string",
      "description": "Process for watchdog to fork",
      "example": "node main.js"
    },
    "envVars": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      },
      "description": "Overrides to environmental variables"
    },
    "constraints": {
      "type": "array",
      "items": {
        "type": "string",
        "description": "Constraints are specific to OpenFaaS Provider",
        "example": "node.platform.os == linux"
      }
    },
    "labels": {
      "description": "A map of labels for making scheduling or routing decisions",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      },
      "example": {
        "foo": "bar"
      }
    },
    "annotations": {
      "description": "A map of annotations for management, orchestration, events and build tasks",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      },
      "example": {
        "topics": "awesome-kafka-topic",
        "foo": "bar"
      }
    },
    "secrets": {
      "type": "array",
      "items": {
        "type": "string",
        "description": "An array of names of secrets that are required to be loaded from the Docker Swarm.",
        "example": "secret-name-1"
      }
    },
    "registryAuth": {
      "type": "string",
      "description": "Private registry base64-encoded basic auth (as present in ~/.docker/config.json)",
      "example": "dXNlcjpwYXNzd29yZA=="
    },
    "limits": {
      "type": "object",
      "properties": {
        "memory": {
          "type": "string",
          "example": "128M"
        },
        "cpu": {
          "type": "string",
          "example": "0.01"
        }
      }
    },
    "requests": {
      "type": "object",
      "properties": {
        "memory": {
          "type": "string",
          "example": "128M"
        },
        "cpu": {
          "type": "string",
          "example": "0.01"
        }
      }
    },
    "readOnlyRootFilesystem": {
      "type": "boolean",
      "description": "Make the root filesystem of the function read-only"
    }
  },
  "required": [
    "service",
    "image",
    "envProcess"
  ]
}
```

 ### FunctionListEntry-modal
 ```json
{
  "type": "object",
  "properties": {
    "name": {
      "description": "The name of the function",
      "type": "string",
      "example": "nodeinfo"
    },
    "image": {
      "description": "The fully qualified docker image name of the function",
      "type": "string",
      "example": "functions/nodeinfo:latest"
    },
    "invocationCount": {
      "description": "The amount of invocations for the specified function",
      "type": "number",
      "format": "integer",
      "example": 1337
    },
    "replicas": {
      "description": "The current minimal ammount of replicas",
      "type": "number",
      "format": "integer",
      "example": 2
    },
    "availableReplicas": {
      "description": "The current available amount of replicas",
      "type": "number",
      "format": "integer",
      "example": 2
    },
    "envProcess": {
      "description": "Process for watchdog to fork",
      "type": "string",
      "example": "node main.js"
    },
    "labels": {
      "description": "A map of labels for making scheduling or routing decisions",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      },
      "example": {
        "foo": "bar"
      }
    },
    "annotations": {
      "description": "A map of annotations for management, orchestration, events and build tasks",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      },
      "example": {
        "topics": "awesome-kafka-topic",
        "foo": "bar"
      }
    }
  },
  "required": [
    "name",
    "image",
    "invocationCount",
    "replicas",
    "availableReplicas",
    "envProcess",
    "labels"
  ]
}
```

 ### Secret-modal
 ```json
{
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "description": "Name of secret",
      "example": "aws-key"
    },
    "value": {
      "type": "string",
      "description": "Value of secret in plain-text",
      "example": "changeme"
    }
  },
  "required": [
    "name"
  ]
}
```

 ### LogEntry-modal
 ```json
{
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "description": "the function name"
    },
    "instance": {
      "type": "string",
      "description": "the name/id of the specific function instance"
    },
    "timestamp": {
      "type": "string",
      "format": "date-time",
      "description": "the timestamp of when the log message was recorded"
    },
    "text": {
      "type": "string",
      "description": "raw log message content"
    }
  }
}
```

 ### SecretName-modal
 ```json
{
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "description": "Name of secret",
      "example": "aws-key"
    }
  }
}
```

