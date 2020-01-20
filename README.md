
# go-fass
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
	go get github.com/vitwit/go-fass
```


## Setup Environment Variables

### Initial Setup

```bash
cp .env.example .env
Update HOST_URL, USER and PASSWORD
```

### Environment Variable

Update the development environment with your keys, for example:

```bash
"export OPENFAAS_HOST_URL='openFasS URL'"
"export OPENFAAS_USER=''"
"export OPENFAAS_PASSWORD=''"

```
<a name="quick-start"></a>
# Quick Start

### Create a client

```
	import fass "go get github.com/vitwit/go-fass";

    client := fass.NewClient(os.Getenv("OPENFAAS_USER"), os.Getenv("OPENFAAS_PASSWORD"),"")
```

<details>

<summary>getFunctions</summary>

getFunctions
---
 **Example**

 ```
 data, error := client.getFunctions()
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
 reqBody := fass.FunctionDefintion{
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

    data, err := client.CreateSystemFunctions(reqBody)
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

 ```js
 data,err = client.updateFunction({
  /** FunctionDefintion modal, description-Function to update,required-true */
})
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
 data,error := client.deleteFunction({
  /** DeleteFunctionRequest modal, description-Function to delete,required-true */
})
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
 data,error := client.handleAlert({
  /** undefined modal,type - object, description-Incoming alert */
})
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

 ```js
 data,error := client.invokeFunctionAsync({
 input:undefined, /** description-(Optional) data to pass to function,required-false */
  _pathParams: {
   functionName:string, /** description-Function name,required-true */
  }
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
 data,error := client.invokeFunction({
 input:undefined, /** description-(Optional) data to pass to function,required-false */
  _pathParams: {
   functionName:string, /** description-Function name,required-true */
  }
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
  data, error := client.scaleFunction({
 input:undefined, /** description-Function to scale plus replica count,required-false */
  _pathParams: {
   functionName:string, /** description-Function name,required-true */
  }
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
 data, error := await client.getFunctionSummary({
  _pathParams: {
   functionName:string, /** description-Function name,required-true */
  }
})
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
    data, error := client.getSecrets()
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

 ```js
 const  { data, error } = await jsFass.createSecret({
  /** Secret modal, description-A new secret to create,required-true */
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

 ```js
 const  { data, error } = await jsFass.updateSecret({
  /** Secret modal, description-Secret to update,required-true */
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
 data, error := client.deleteSecret({
  /** SecretName modal, description-Secret to delete,required-true */
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
 data, error := client.getLogsOfAFunction({
  _params: {
   name:string, /** description-Function name,required-true */
   since:string, /** description-Only return logs after a specific date (RFC3339),required-false */
   tail:integer, /** description-Sets the maximum number of log messages to return, <=0 means unlimited,required-false */
   follow:boolean, /** description-When true, the request will stream logs until the request timeout,required-false */
  }
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
 data, error := client.getInfo()
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
 data, error := client.checkHealth()
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

