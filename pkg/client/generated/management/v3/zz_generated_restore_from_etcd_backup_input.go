package client

const (
	RestoreFromEtcdBackupInputType                  = "restoreFromEtcdBackupInput"
	RestoreFromEtcdBackupInputFieldEtcdBackupID     = "etcdBackupId"
	RestoreFromEtcdBackupInputFieldRestoreBkeConfig = "restoreBkeConfig"
)

type RestoreFromEtcdBackupInput struct {
	EtcdBackupID     string `json:"etcdBackupId,omitempty" yaml:"etcdBackupId,omitempty"`
	RestoreBkeConfig string `json:"restoreBkeConfig,omitempty" yaml:"restoreBkeConfig,omitempty"`
}
