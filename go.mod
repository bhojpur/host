module github.com/bhojpur/host

go 1.16

require (
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/bhojpur/errors v0.0.3
	github.com/bhojpur/units v0.0.2
	github.com/blang/semver v3.5.1+incompatible
	github.com/coreos/go-semver v0.3.0
	github.com/docker/go-connections v0.4.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.20.2
	github.com/go-openapi/loads v0.21.1
	github.com/go-openapi/runtime v0.23.3
	github.com/go-openapi/spec v0.20.5
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.21.1
	github.com/go-openapi/validate v0.21.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.5.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/maruel/panicparse v1.6.2
	github.com/matryer/moq v0.2.7
	github.com/mattn/go-colorable v0.1.12
	github.com/moby/locker v1.0.1
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.55.1
	github.com/prometheus/client_golang v1.12.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad
	golang.org/x/tools v0.1.10
	google.golang.org/grpc v1.45.0
	gotest.tools/v3 v3.0.3
	k8s.io/apiextensions-apiserver v0.23.6
	k8s.io/apimachinery v0.23.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.23.6
	k8s.io/gengo v0.0.0-20220307231824-4627b89bbf1b
	k8s.io/klog v1.0.0
	k8s.io/kube-aggregator v0.23.6
	knative.dev/pkg v0.0.0-20220418171127-12be06090b51
	sigs.k8s.io/cli-utils v0.29.4
	sigs.k8s.io/cluster-api v1.1.3
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/containerd/containerd v1.6.2 // indirect
	github.com/fsnotify/fsnotify v1.5.2 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-openapi/analysis v0.21.3 // indirect
	github.com/google/gnostic v0.6.8 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/sys/mount v0.3.2 // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	go.mongodb.org/mongo-driver v1.9.0 // indirect
	go.starlark.net v0.0.0-20220328144851-d1966c6b9fcd // indirect
	golang.org/x/term v0.0.0-20220411215600-e5f449aeb171 // indirect
	sigs.k8s.io/controller-runtime v0.11.2 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/kustomize/api v0.11.4 // indirect
)

require (
	github.com/Azure/azure-sdk-for-go v63.3.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.26
	github.com/Azure/go-autorest/autorest/adal v0.9.18
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/aws/aws-sdk-go v1.43.43
	github.com/bugsnag/bugsnag-go v2.1.2+incompatible
	github.com/bugsnag/panicwrap v1.3.4 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/digitalocean/godo v1.78.0
	github.com/docker/distribution v2.8.1+incompatible
	github.com/docker/docker v20.10.14+incompatible
	github.com/evanphx/json-patch v5.6.0+incompatible
	github.com/exoscale/egoscale v1.19.0
	github.com/exponent-io/jsonpath v0.0.0-20210407135951-1de76d718b3f // indirect
	github.com/go-ini/ini v1.66.4
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gophercloud/gophercloud v0.24.0
	github.com/gophercloud/utils v0.0.0-20220307143606-8e7800759d16
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/heptio/authenticator v0.0.0-20180409043135-d282f87a1972
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/mcuadros/go-version v0.0.0-20190830083331-035f6764e8d2
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.1.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.34.0 // indirect
	github.com/racker/perigee v0.1.0 // indirect
	github.com/rackspace/gophercloud v1.0.0
	github.com/russross/blackfriday v1.6.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/skarademir/naturalsort v0.0.0-20150715044055-69a5d87bef62
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.1
	github.com/urfave/cli v1.22.5
	github.com/vmware/govcloudair v0.0.2
	github.com/vmware/govmomi v0.27.4
	go.etcd.io/etcd/client/v2 v2.305.3
	go.etcd.io/etcd/client/v3 v3.5.3
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
	golang.org/x/net v0.0.0-20220420153159-1850ba15e1be
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	google.golang.org/api v0.75.0
	google.golang.org/genproto v0.0.0-20220420195807-44278fea765b // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.23.6
	k8s.io/apiserver v0.23.6
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/kubectl v0.23.6
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
)

replace (
	github.com/bhojpur/host/pkg/client => ./pkg/client
	github.com/docker/docker => github.com/docker/docker v20.10.14+incompatible // oras dep requires a replace is set
	k8s.io/client-go => k8s.io/client-go v0.23.6
	knative.dev/pkg => github.com/bhojpur/knative-pkg v0.0.3
	sigs.k8s.io/json => sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2
)

replace k8s.io/api => k8s.io/api v0.23.6

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.6

replace k8s.io/apimachinery => k8s.io/apimachinery v0.23.7-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.23.6

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.6

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.6

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.6

replace k8s.io/code-generator => k8s.io/code-generator v0.23.7-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.23.6

replace k8s.io/cri-api => k8s.io/cri-api v0.23.7-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.6

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.6

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.6

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.6

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.6

replace k8s.io/kubelet => k8s.io/kubelet v0.23.6

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.6

replace k8s.io/metrics => k8s.io/metrics v0.23.6

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.6

replace k8s.io/component-helpers => k8s.io/component-helpers v0.23.6

replace k8s.io/controller-manager => k8s.io/controller-manager v0.23.6

replace k8s.io/kubectl => k8s.io/kubectl v0.23.6

replace k8s.io/mount-utils => k8s.io/mount-utils v0.23.7-rc.0

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.6

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.23.6

replace k8s.io/sample-controller => k8s.io/sample-controller v0.23.6
