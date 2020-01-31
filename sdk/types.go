package sdk

import (
	"time"
)

// HTTP Request
type QueryParams map[string]string

// Response holds the response from an API call.
type HTTPResponse struct {
	StatusCode int                 // e.g. 200
	Body       string              // e.g. {"result: success"}
	Headers    map[string][]string // e.g. Authorization:"Basic qwertyuiop"
}

// client
type OpenFaasClient struct {
	*FaasRequestDefinition
}

type FaasRequestDefinition struct {
	// Method can be GET/PUT/POST/PATCH/DELETE
	Method string
	// This is the address of OpenFaas gateway you get e.g. http://127.0.0.1:8080
	GatewayAddress string
	// Path of the API being called e.g. /system/functions (Trailing slash must be specified)
	Path string
	// URL is the full path of the API being called which is basically gateway address + path (E.g. http://127.0.0.1:8080/system/functions)
	URL string
	// http headers
	Headers map[string]string
	// query params
	QueryParams QueryParams
	// request body to send. Must be byte array
	Body []byte
	// Cluster Type valid e.g => swarm/kubernetes
	ClusterType string
}

type FaasGatewayCredentials struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	GatewayAddress string `json:"gatewayAddress"`
	ClusterType    string `json:"clusterType"`
}

type FunctionListEntry []struct {
	Name              string            `json:"name"`
	Image             string            `json:"image"`
	InvocationCount   int64             `json:"invocationCount"`
	Replicas          int64             `json:"replicas"`
	AvailableReplicas int64             `json:"availableReplicas"`
	EnvProcess        string            `json:"envProcess"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

type Annotations struct {
	Topics string `json:"topics"`
	Foo    string `json:"foo"`
}

type EnvVars struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type Limits struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

type Requests struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

type SyncInvocationOpts struct {
	Body         interface{}
	FunctionName string
}

type AsyncInvocationOpts struct {
	Body         interface{}
	FunctionName string
	CallbackURL  string
}

type FunctionDefintion struct {
	Service                string            `json:"service"`
	Network                string            `json:"network"`
	Image                  string            `json:"image"`
	EnvProcess             string            `json:"envProcess"`
	EnvVars                EnvVars           `json:"envVars"`
	Constraints            []string          `json:"constraints"`
	Labels                 map[string]string `json:"labels"`
	Annotations            Annotations       `json:"annotations"`
	Secrets                []string          `json:"secrets"`
	RegistryAuth           string            `json:"registryAuth"`
	Limits                 Limits            `json:"limits"`
	Requests               Requests          `json:"requests"`
	ReadOnlyRootFilesystem bool              `json:"readOnlyRootFilesystem"`
}

type DeleteFunctionBodyOpts struct {
	FunctionName string `json:"functionName"`
}

type SystemAlertLables struct {
	Alertname    string `json:"alertname"`
	Code         string `json:"code"`
	FunctionName string `json:"function_name"`
	Instance     string `json:"instance"`
	Job          string `json:"job"`
	Monitor      string `json:"monitor"`
	Service      string `json:"service"`
	Severity     string `json:"severity"`
	Value        string `json:"value"`
}

type SystemAlertAnnotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

type SystemAlertsStruct struct {
	Status       string                 `json:"status"`
	Labels       SystemAlertLables      `json:"labels"`
	Annotations  SystemAlertAnnotations `json:"annotations"`
	StartsAt     time.Time              `json:"startsAt"`
	EndsAt       time.Time              `json:"endsAt"`
	GeneratorURL string                 `json:"generatorURL"`
}

type GroupLabels struct {
	Alertname string `json:"alertname"`
	Service   string `json:"service"`
}

type CommonLabels struct {
	Alertname    string `json:"alertname"`
	Code         string `json:"code"`
	FunctionName string `json:"function_name"`
	Instance     string `json:"instance"`
	Job          string `json:"job"`
	Monitor      string `json:"monitor"`
	Service      string `json:"service"`
	Severity     string `json:"severity"`
	Value        string `json:"value"`
}

type CommonAnnotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

type SystemAlertBodyOpts struct {
	Receiver          string               `json:"receiver"`
	Status            string               `json:"status"`
	Alerts            []SystemAlertsStruct `json:"alerts"`
	GroupLabels       GroupLabels          `json:"groupLabels"`
	CommonLabels      CommonLabels         `json:"commonLabels"`
	CommonAnnotations CommonAnnotations    `json:"commonAnnotations"`
	ExternalURL       string               `json:"externalURL"`
	Version           string               `json:"version"`
	GroupKey          int64                `json:"groupKey"`
}

type ScaleFunctionBodyOpts struct {
	Service  string `json:"service"`
	Replicas int    `json:"replicas"`
}

type SecretBodyOpts struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type LogEntry struct {
	Name      string    `json:"name"`
	Instance  string    `json:"instance"`
	Timestamp time.Time `json:"timestamp"`
	Text      string    `json:"text"`
}

type SecretNameBodyOpts struct {
	Name string `json:"name"`
}

type SystemLogsQueryOpts struct {
	Name   string `json:"name"`
	Since  string `json:"since"`
	Tail   int    `json:"tail"`
	Follow bool   `json:"follow"`
}
