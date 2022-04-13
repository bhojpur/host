package client

const (
	ContainerDriverStatusType                = "containerDriverStatus"
	ContainerDriverStatusFieldActualURL      = "actualUrl"
	ContainerDriverStatusFieldConditions     = "conditions"
	ContainerDriverStatusFieldDisplayName    = "displayName"
	ContainerDriverStatusFieldExecutablePath = "executablePath"
)

type ContainerDriverStatus struct {
	ActualURL      string      `json:"actualUrl,omitempty" yaml:"actualUrl,omitempty"`
	Conditions     []Condition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	DisplayName    string      `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	ExecutablePath string      `json:"executablePath,omitempty" yaml:"executablePath,omitempty"`
}
