package client

const (
	BKEConfigNodePlanType             = "bkeConfigNodePlan"
	BKEConfigNodePlanFieldAddress     = "address"
	BKEConfigNodePlanFieldAnnotations = "annotations"
	BKEConfigNodePlanFieldFiles       = "files"
	BKEConfigNodePlanFieldLabels      = "labels"
	BKEConfigNodePlanFieldPortChecks  = "portChecks"
	BKEConfigNodePlanFieldProcesses   = "processes"
	BKEConfigNodePlanFieldTaints      = "taints"
)

type BKEConfigNodePlan struct {
	Address     string             `json:"address,omitempty" yaml:"address,omitempty"`
	Annotations map[string]string  `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Files       []File             `json:"files,omitempty" yaml:"files,omitempty"`
	Labels      map[string]string  `json:"labels,omitempty" yaml:"labels,omitempty"`
	PortChecks  []PortCheck        `json:"portChecks,omitempty" yaml:"portChecks,omitempty"`
	Processes   map[string]Process `json:"processes,omitempty" yaml:"processes,omitempty"`
	Taints      []BKETaint         `json:"taints,omitempty" yaml:"taints,omitempty"`
}
