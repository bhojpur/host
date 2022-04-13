package client

const (
	UkeConfigType                        = "ukeConfig"
	UkeConfigFieldClusterUpgradeStrategy = "ukeupgradeStrategy"
	UkeConfigFieldVersion                = "kubernetesVersion"
)

type UkeConfig struct {
	ClusterUpgradeStrategy *ClusterUpgradeStrategy `json:"ukeupgradeStrategy,omitempty" yaml:"ukeupgradeStrategy,omitempty"`
	Version                string                  `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
}
