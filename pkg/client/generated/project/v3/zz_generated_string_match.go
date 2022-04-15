package client

const (
	StringMatchType        = "stringMatch"
	StringMatchFieldExact  = "exact"
	StringMatchFieldRegex  = "regex"
	StringMatchFieldSuffix = "prefix"
)

type StringMatch struct {
	Exact  string `json:"exact,omitempty" yaml:"exact,omitempty"`
	Regex  string `json:"regex,omitempty" yaml:"regex,omitempty"`
	Suffix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
