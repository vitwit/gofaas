package go_fass

import "time"

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

type DeleteFunctionRequest struct {
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

type SystemAlert struct {
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

type Secret struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type LogEntry struct {
	Name      string    `json:"name"`
	Instance  string    `json:"instance"`
	Timestamp time.Time `json:"timestamp"`
	Text      string    `json:"text"`
}

type SecretName struct {
	Name string `json:"name"`
}
