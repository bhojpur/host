package client

const (
	VirtualServiceSpecType          = "virtualServiceSpec"
	VirtualServiceSpecFieldGateways = "gateways"
	VirtualServiceSpecFieldHosts    = "hosts"
	VirtualServiceSpecFieldHttp     = "http"
	VirtualServiceSpecFieldTcp      = "tcp"
)

type VirtualServiceSpec struct {
	Gateways []string    `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	Hosts    []string    `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Http     []HTTPRoute `json:"http,omitempty" yaml:"http,omitempty"`
	Tcp      []TCPRoute  `json:"tcp,omitempty" yaml:"tcp,omitempty"`
}
