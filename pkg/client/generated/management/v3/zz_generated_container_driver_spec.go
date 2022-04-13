package client

const (
	ContainerDriverSpecType                  = "containerDriverSpec"
	ContainerDriverSpecFieldActive           = "active"
	ContainerDriverSpecFieldBuiltIn          = "builtIn"
	ContainerDriverSpecFieldChecksum         = "checksum"
	ContainerDriverSpecFieldUIURL            = "uiUrl"
	ContainerDriverSpecFieldURL              = "url"
	ContainerDriverSpecFieldWhitelistDomains = "whitelistDomains"
)

type ContainerDriverSpec struct {
	Active           bool     `json:"active,omitempty" yaml:"active,omitempty"`
	BuiltIn          bool     `json:"builtIn,omitempty" yaml:"builtIn,omitempty"`
	Checksum         string   `json:"checksum,omitempty" yaml:"checksum,omitempty"`
	UIURL            string   `json:"uiUrl,omitempty" yaml:"uiUrl,omitempty"`
	URL              string   `json:"url,omitempty" yaml:"url,omitempty"`
	WhitelistDomains []string `json:"whitelistDomains,omitempty" yaml:"whitelistDomains,omitempty"`
}
