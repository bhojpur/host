package client

const (
	HTTPMatchRequestType           = "httpMatchRequest"
	HTTPMatchRequestFieldAuthority = "authority"
	HTTPMatchRequestFieldHeaders   = "headers"
	HTTPMatchRequestFieldMethod    = "method"
	HTTPMatchRequestFieldPort      = "port"
	HTTPMatchRequestFieldScheme    = "scheme"
	HTTPMatchRequestFieldUri       = "uri"
)

type HTTPMatchRequest struct {
	Authority *StringMatch           `json:"authority,omitempty" yaml:"authority,omitempty"`
	Headers   map[string]StringMatch `json:"headers,omitempty" yaml:"headers,omitempty"`
	Method    *StringMatch           `json:"method,omitempty" yaml:"method,omitempty"`
	Port      *int64                 `json:"port,omitempty" yaml:"port,omitempty"`
	Scheme    *StringMatch           `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Uri       *StringMatch           `json:"uri,omitempty" yaml:"uri,omitempty"`
}
