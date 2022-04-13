package client

const (
	AlertmanagerConfigurationType      = "alertmanagerConfiguration"
	AlertmanagerConfigurationFieldName = "name"
)

type AlertmanagerConfiguration struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
