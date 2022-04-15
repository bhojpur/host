package client

const (
	SubsetType               = "subset"
	SubsetFieldLabels        = "labels"
	SubsetFieldName          = "port"
	SubsetFieldTrafficPolicy = "trafficPolicy"
)

type Subset struct {
	Labels        map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name          string            `json:"port,omitempty" yaml:"port,omitempty"`
	TrafficPolicy *TrafficPolicy    `json:"trafficPolicy,omitempty" yaml:"trafficPolicy,omitempty"`
}
