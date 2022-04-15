package client

const (
	L4MatchAttributesType                   = "l4MatchAttributes"
	L4MatchAttributesFieldDestinationSubnet = "destinationSubnet"
	L4MatchAttributesFieldGateways          = "gateways"
	L4MatchAttributesFieldPort              = "port"
	L4MatchAttributesFieldSourceLabel       = "sourceLabel"
	L4MatchAttributesFieldSourceSubnet      = "sourceSubnet"
)

type L4MatchAttributes struct {
	DestinationSubnet string            `json:"destinationSubnet,omitempty" yaml:"destinationSubnet,omitempty"`
	Gateways          []string          `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	Port              int64             `json:"port,omitempty" yaml:"port,omitempty"`
	SourceLabel       map[string]string `json:"sourceLabel,omitempty" yaml:"sourceLabel,omitempty"`
	SourceSubnet      string            `json:"sourceSubnet,omitempty" yaml:"sourceSubnet,omitempty"`
}
