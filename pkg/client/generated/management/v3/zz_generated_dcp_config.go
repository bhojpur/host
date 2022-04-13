package client

const (
	DcpConfigType                        = "dcpConfig"
	DcpConfigFieldClusterUpgradeStrategy = "dcpupgradeStrategy"
	DcpConfigFieldVersion                = "kubernetesVersion"
)

type DcpConfig struct {
	ClusterUpgradeStrategy *ClusterUpgradeStrategy `json:"dcpupgradeStrategy,omitempty" yaml:"dcpupgradeStrategy,omitempty"`
	Version                string                  `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
}
