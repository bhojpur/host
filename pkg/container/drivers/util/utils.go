package util

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
	"context"
	"fmt"
	"time"

	v3 "github.com/bhojpur/host/pkg/engine/types"
	errs "github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	defaultNamespace          = "default"
	bhojpurNamespace          = "bhojpur-system"
	clusterAdmin              = "cluster-admin"
	netesDefault              = "netes-default"
	hostfarmEngine            = "hostfarm-engine"
	oldClusterRoleBindingName = "netes-default-clusterRoleBinding"
	newClusterRoleBindingName = "system-netes-default-clusterRoleBinding"
)

// GenerateServiceAccountToken generate a serviceAccountToken for clusterAdmin given a rest clientset
func GenerateServiceAccountToken(clientset kubernetes.Interface) (string, error) {
	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: bhojpurNamespace,
		},
	}, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return "", err
	}

	serviceAccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: hostfarmEngine,
		},
	}

	_, err = clientset.CoreV1().ServiceAccounts(bhojpurNamespace).Create(context.TODO(), serviceAccount, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return "", fmt.Errorf("error creating service account: %v", err)
	}

	adminRole := &v1beta1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterAdmin,
		},
		Rules: []v1beta1.PolicyRule{
			{
				APIGroups: []string{"*"},
				Resources: []string{"*"},
				Verbs:     []string{"*"},
			},
			{
				NonResourceURLs: []string{"*"},
				Verbs:           []string{"*"},
			},
		},
	}
	clusterAdminRole, err := clientset.RbacV1beta1().ClusterRoles().Get(context.TODO(), clusterAdmin, metav1.GetOptions{})
	if err != nil {
		clusterAdminRole, err = clientset.RbacV1beta1().ClusterRoles().Create(context.TODO(), adminRole, metav1.CreateOptions{})
		if err != nil {
			return "", fmt.Errorf("error creating admin role: %v", err)
		}
	}

	clusterRoleBinding := &v1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: newClusterRoleBindingName,
		},
		Subjects: []v1beta1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      serviceAccount.Name,
				Namespace: bhojpurNamespace,
				APIGroup:  v1.GroupName,
			},
		},
		RoleRef: v1beta1.RoleRef{
			Kind:     "ClusterRole",
			Name:     clusterAdminRole.Name,
			APIGroup: v1beta1.GroupName,
		},
	}
	if _, err = clientset.RbacV1beta1().ClusterRoleBindings().Create(context.TODO(), clusterRoleBinding, metav1.CreateOptions{}); err != nil && !errors.IsAlreadyExists(err) {
		return "", fmt.Errorf("error creating role bindings: %v", err)
	}

	start := time.Millisecond * 250
	for i := 0; i < 5; i++ {
		time.Sleep(start)
		if serviceAccount, err = clientset.CoreV1().ServiceAccounts(bhojpurNamespace).Get(context.TODO(), serviceAccount.Name, metav1.GetOptions{}); err != nil {
			return "", fmt.Errorf("error getting service account: %v", err)
		}

		if len(serviceAccount.Secrets) > 0 {
			secret := serviceAccount.Secrets[0]
			secretObj, err := clientset.CoreV1().Secrets(bhojpurNamespace).Get(context.TODO(), secret.Name, metav1.GetOptions{})
			if err != nil {
				return "", fmt.Errorf("error getting secret: %v", err)
			}
			if token, ok := secretObj.Data["token"]; ok {
				return string(token), nil
			}
		}
		start = start * 2
	}

	return "", errs.New("failed to fetch serviceAccountToken")
}

func DeleteLegacyServiceAccountAndRoleBinding(clientset kubernetes.Interface) error {
	_, err := clientset.CoreV1().ServiceAccounts(defaultNamespace).Get(context.TODO(), netesDefault, metav1.GetOptions{})
	if !errors.IsNotFound(err) {
		err = clientset.CoreV1().ServiceAccounts(defaultNamespace).Delete(context.TODO(), netesDefault, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	_, err = clientset.RbacV1beta1().ClusterRoleBindings().Get(context.TODO(), oldClusterRoleBindingName, metav1.GetOptions{})
	if !errors.IsNotFound(err) {
		err = clientset.RbacV1beta1().ClusterRoleBindings().Delete(context.TODO(), oldClusterRoleBindingName, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func ConvertToBkeConfig(config string) (v3.BhojpurKubernetesEngineConfig, error) {
	var bkeConfig v3.BhojpurKubernetesEngineConfig
	if err := yaml.Unmarshal([]byte(config), &bkeConfig); err != nil {
		return bkeConfig, err
	}
	return bkeConfig, nil
}
