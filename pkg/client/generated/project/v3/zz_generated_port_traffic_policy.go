package client

const (
	PortTrafficPolicyType                  = "portTrafficPolicy"
	PortTrafficPolicyFieldConnectionPool   = "connectionPool"
	PortTrafficPolicyFieldLoadBalancer     = "loadBalancer"
	PortTrafficPolicyFieldOutlierDetection = "outlierDetection"
	PortTrafficPolicyFieldPort             = "port"
	PortTrafficPolicyFieldTls              = "tls"
)

type PortTrafficPolicy struct {
	ConnectionPool   *ConnectionPoolSettings `json:"connectionPool,omitempty" yaml:"connectionPool,omitempty"`
	LoadBalancer     *LoadBalancerSettings   `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`
	OutlierDetection *OutlierDetection       `json:"outlierDetection,omitempty" yaml:"outlierDetection,omitempty"`
	Port             *PortSelector           `json:"port,omitempty" yaml:"port,omitempty"`
	Tls              *TLSSettings            `json:"tls,omitempty" yaml:"tls,omitempty"`
}
