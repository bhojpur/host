module github.com/bhojpur/host/pkg/client

go 1.17

replace k8s.io/client-go => k8s.io/client-go v0.23.3

require (
	github.com/sirupsen/logrus v1.6.0 // indirect
	k8s.io/apimachinery v0.21.0
)