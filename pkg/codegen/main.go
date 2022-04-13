package main

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"os"

	fleet "github.com/bhojpur/host/pkg/apis/fleet.bhojpur.net/v1alpha1"
	v3 "github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	planv1 "github.com/bhojpur/host/pkg/apis/upgrade.bhojpur.net/v1"
	"github.com/bhojpur/host/pkg/codegen/generator"
	controllergen "github.com/bhojpur/host/pkg/common/controller-gen"
	"github.com/bhojpur/host/pkg/common/controller-gen/args"
	"github.com/bhojpur/host/pkg/core/types"
	clusterSchema "github.com/bhojpur/host/pkg/schemas/cluster.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/schemas/factory"
	managementSchema "github.com/bhojpur/host/pkg/schemas/management.bhojpur.net/v3"
	publicSchema "github.com/bhojpur/host/pkg/schemas/management.bhojpur.net/v3public"
	projectSchema "github.com/bhojpur/host/pkg/schemas/project.bhojpur.net/v3"
	istiov1alpha3 "github.com/knative/pkg/apis/istio/v1alpha3"
	"github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	scalingv2beta2 "k8s.io/api/autoscaling/v2beta2"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	coordinationv1 "k8s.io/api/coordination/v1"
	v1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	knetworkingv1 "k8s.io/api/networking/v1"
	networkingv1 "k8s.io/api/networking/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	apiv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
)

func main() {
	os.Unsetenv("GOPATH")

	controllergen.Run(args.Options{
		OutputPackage: "github.com/bhojpur/host/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			v1.GroupName: {
				Types: []interface{}{
					v1.Event{},
					v1.Node{},
					v1.Namespace{},
					v1.Secret{},
					v1.Service{},
					v1.ServiceAccount{},
					v1.Endpoints{},
					v1.ConfigMap{},
					v1.PersistentVolume{},
					v1.PersistentVolumeClaim{},
					v1.Pod{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			extensionsv1beta1.GroupName: {
				Types: []interface{}{
					extensionsv1beta1.Ingress{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			rbacv1.GroupName: {
				Types: []interface{}{
					rbacv1.Role{},
					rbacv1.RoleBinding{},
					rbacv1.ClusterRole{},
					rbacv1.ClusterRoleBinding{},
				},
				OutputControllerPackageName: "rbac",
				InformersPackage:            "k8s.io/client-go/informers",
				ClientSetPackage:            "k8s.io/client-go/kubernetes",
				ListersPackage:              "k8s.io/client-go/listers",
			},
			appsv1.GroupName: {
				Types: []interface{}{
					appsv1.Deployment{},
					appsv1.DaemonSet{},
					appsv1.StatefulSet{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			storagev1.GroupName: {
				OutputControllerPackageName: "storage",
				Types: []interface{}{
					storagev1.StorageClass{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			apiextv1.GroupName: {
				Types: []interface{}{
					apiextv1.CustomResourceDefinition{},
				},
				ClientSetPackage: "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset",
				InformersPackage: "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions",
				ListersPackage:   "k8s.io/apiextensions-apiserver/pkg/client/listers",
			},
			apiv1.GroupName: {
				Types: []interface{}{
					apiv1.APIService{},
				},
				ClientSetPackage: "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset",
				InformersPackage: "k8s.io/kube-aggregator/pkg/client/informers/externalversions",
				ListersPackage:   "k8s.io/kube-aggregator/pkg/client/listers",
			},
			batchv1.GroupName: {
				Types: []interface{}{
					batchv1.Job{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			networkingv1.GroupName: {
				Types: []interface{}{
					networkingv1.NetworkPolicy{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			admissionregistrationv1.GroupName: {
				Types: []interface{}{
					admissionregistrationv1.ValidatingWebhookConfiguration{},
					admissionregistrationv1.MutatingWebhookConfiguration{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			coordinationv1.GroupName: {
				Types: []interface{}{
					coordinationv1.Lease{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},

			"management.bhojpur.net": {
				PackageName: "management.bhojpur.net",
				Types: []interface{}{
					// All structs with an embedded ObjectMeta field will be picked up
					"./pkg/apis/management.bhojpur.net/v3",
					v3.ProjectCatalog{},
					v3.ClusterCatalog{},
				},
				GenerateTypes: true,
			},
			"ui.bhojpur.net": {
				PackageName: "ui.bhojpur.net",
				Types: []interface{}{
					"./pkg/apis/ui.bhojpur.net/v1",
				},
				GenerateTypes: true,
			},
			"cluster.bhojpur.net": {
				PackageName: "cluster.bhojpur.net",
				Types: []interface{}{
					// All structs with an embedded ObjectMeta field will be picked up
					"./pkg/apis/cluster.bhojpur.net/v3",
				},
				GenerateTypes: true,
			},
			"project.bhojpur.net": {
				PackageName: "project.bhojpur.net",
				Types: []interface{}{
					// All structs with an embedded ObjectMeta field will be picked up
					"./pkg/apis/project.bhojpur.net/v3",
				},
				GenerateTypes: true,
			},
			"catalog.bhojpur.net": {
				PackageName: "catalog.bhojpur.net",
				Types: []interface{}{
					// All structs with an embedded ObjectMeta field will be picked up
					"./pkg/apis/catalog.bhojpur.net/v1",
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
			"upgrade.bhojpur.net": {
				PackageName: "upgrade.bhojpur.net",
				Types: []interface{}{
					planv1.Plan{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
			"provisioning.bhojpur.net": {
				Types: []interface{}{
					"./pkg/apis/provisioning.bhojpur.net/v1",
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
			"fleet.bhojpur.net": {
				Types: []interface{}{
					fleet.Bundle{},
					fleet.Cluster{},
				},
				GenerateTypes: true,
			},
			"aks.bhojpur.net": {
				Types: []interface{}{
					"./pkg/apis/aks.bhojpur.net/v1",
				},
				GenerateTypes: true,
			},
			"eks.bhojpur.net": {
				Types: []interface{}{
					"./pkg/apis/eks.bhojpur.net/v1",
				},
				GenerateTypes: true,
			},
			"gke.bhojpur.net": {
				Types: []interface{}{
					"./pkg/apis/gke.bhojpur.net/v1",
				},
				GenerateTypes: true,
			},
			"bke.bhojpur.net": {
				Types: []interface{}{
					"./pkg/apis/bke.bhojpur.net/v1",
				},
				GenerateTypes: true,
			},
			"cluster.x-k8s.io": {
				Types: []interface{}{
					capi.Machine{},
					capi.MachineDeployment{},
					capi.Cluster{},
				},
			},
		},
	})

	clusterAPIVersion := &types.APIVersion{Group: capi.GroupVersion.Group, Version: capi.GroupVersion.Version, Path: "/v1"}
	generator.GenerateClient(factory.Schemas(clusterAPIVersion).Init(func(schemas *types.Schemas) *types.Schemas {
		return schemas.MustImportAndCustomize(clusterAPIVersion, capi.Machine{}, func(schema *types.Schema) {
			schema.ID = "cluster.x-k8s.io.machine"
		})
	}), nil)

	generator.GenerateComposeType(projectSchema.Schemas, managementSchema.Schemas, clusterSchema.Schemas)
	generator.Generate(managementSchema.Schemas, map[string]bool{
		"userAttribute": true,
	})
	generator.GenerateClient(publicSchema.PublicSchemas, nil)
	generator.Generate(clusterSchema.Schemas, map[string]bool{
		"clusterUserAttribute": true,
		"clusterAuthToken":     true,
	})
	generator.Generate(projectSchema.Schemas, nil)
	generator.GenerateNativeTypes(v1.SchemeGroupVersion, []interface{}{
		v1.Endpoints{},
		v1.PersistentVolumeClaim{},
		v1.Pod{},
		v1.Service{},
		v1.Secret{},
		v1.ConfigMap{},
		v1.ServiceAccount{},
		v1.ReplicationController{},
		v1.ResourceQuota{},
		v1.LimitRange{},
	}, []interface{}{
		v1.Node{},
		v1.ComponentStatus{},
		v1.Namespace{},
		v1.Event{},
	})
	generator.GenerateNativeTypes(appsv1.SchemeGroupVersion, []interface{}{
		appsv1.Deployment{},
		appsv1.DaemonSet{},
		appsv1.StatefulSet{},
		appsv1.ReplicaSet{},
	}, nil)
	generator.GenerateNativeTypes(rbacv1.SchemeGroupVersion, []interface{}{
		rbacv1.RoleBinding{},
		rbacv1.Role{},
	}, []interface{}{
		rbacv1.ClusterRoleBinding{},
		rbacv1.ClusterRole{},
	})
	generator.GenerateNativeTypes(knetworkingv1.SchemeGroupVersion, []interface{}{
		knetworkingv1.NetworkPolicy{},
		knetworkingv1.Ingress{},
	}, nil)
	generator.GenerateNativeTypes(batchv1.SchemeGroupVersion, []interface{}{
		batchv1.Job{},
	}, nil)
	generator.GenerateNativeTypes(batchv1beta1.SchemeGroupVersion, []interface{}{
		batchv1beta1.CronJob{},
	}, nil)
	generator.GenerateNativeTypes(extv1beta1.SchemeGroupVersion,
		[]interface{}{
			extv1beta1.Ingress{},
		},
		nil,
	)
	generator.GenerateNativeTypes(policyv1beta1.SchemeGroupVersion,
		nil,
		[]interface{}{
			policyv1beta1.PodSecurityPolicy{},
		},
	)
	generator.GenerateNativeTypes(storagev1.SchemeGroupVersion,
		nil,
		[]interface{}{
			storagev1.StorageClass{},
		},
	)
	generator.GenerateNativeTypes(
		k8sschema.GroupVersion{Group: monitoring.GroupName, Version: monitoringv1.Version},
		[]interface{}{
			monitoringv1.Prometheus{},
			monitoringv1.Alertmanager{},
			monitoringv1.PrometheusRule{},
			monitoringv1.ServiceMonitor{},
		},
		nil,
	)
	generator.GenerateNativeTypes(scalingv2beta2.SchemeGroupVersion,
		[]interface{}{
			scalingv2beta2.HorizontalPodAutoscaler{},
		},
		nil,
	)
	generator.GenerateNativeTypes(istiov1alpha3.SchemeGroupVersion,
		[]interface{}{
			istiov1alpha3.VirtualService{},
			istiov1alpha3.DestinationRule{},
		},
		nil,
	)
	generator.GenerateNativeTypes(apiregistrationv1.SchemeGroupVersion,
		nil,
		[]interface{}{
			apiregistrationv1.APIService{},
		},
	)
}
