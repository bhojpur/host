package client

const (
	BKETaintType           = "bkeTaint"
	BKETaintFieldEffect    = "effect"
	BKETaintFieldKey       = "key"
	BKETaintFieldTimeAdded = "timeAdded"
	BKETaintFieldValue     = "value"
)

type BKETaint struct {
	Effect    string `json:"effect,omitempty" yaml:"effect,omitempty"`
	Key       string `json:"key,omitempty" yaml:"key,omitempty"`
	TimeAdded string `json:"timeAdded,omitempty" yaml:"timeAdded,omitempty"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
}
