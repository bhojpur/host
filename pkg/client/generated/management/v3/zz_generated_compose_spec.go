package client

const (
	ComposeSpecType                = "composeSpec"
	ComposeSpecFieldBhojpurCompose = "bhojpurCompose"
)

type ComposeSpec struct {
	BhojpurCompose string `json:"bhojpurCompose,omitempty" yaml:"bhojpurCompose,omitempty"`
}
