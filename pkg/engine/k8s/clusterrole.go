package k8s

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

	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func UpdateClusterRoleBindingFromYaml(k8sClient *kubernetes.Clientset, clusterRoleBindingYaml string) error {
	clusterRoleBinding := rbacv1.ClusterRoleBinding{}
	if err := DecodeYamlResource(&clusterRoleBinding, clusterRoleBindingYaml); err != nil {
		return err
	}
	return retryTo(updateClusterRoleBinding, k8sClient, clusterRoleBinding, DefaultRetries, DefaultSleepSeconds)
}

func updateClusterRoleBinding(k8sClient *kubernetes.Clientset, crb interface{}) error {
	clusterRoleBinding := crb.(rbacv1.ClusterRoleBinding)
	if _, err := k8sClient.RbacV1().ClusterRoleBindings().Create(context.TODO(), &clusterRoleBinding, metav1.CreateOptions{}); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return err
		}
		if _, err := k8sClient.RbacV1().ClusterRoleBindings().Update(context.TODO(), &clusterRoleBinding, metav1.UpdateOptions{}); err != nil {
			return err
		}
	}
	return nil
}

func UpdateClusterRoleFromYaml(k8sClient *kubernetes.Clientset, clusterRoleYaml string) error {
	clusterRole := rbacv1.ClusterRole{}
	if err := DecodeYamlResource(&clusterRole, clusterRoleYaml); err != nil {
		return err
	}

	return retryTo(updateClusterRole, k8sClient, clusterRole, DefaultRetries, DefaultSleepSeconds)
}

func updateClusterRole(k8sClient *kubernetes.Clientset, cr interface{}) error {
	clusterRole := cr.(rbacv1.ClusterRole)
	if _, err := k8sClient.RbacV1().ClusterRoles().Create(context.TODO(), &clusterRole, metav1.CreateOptions{}); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return err
		}
		if _, err := k8sClient.RbacV1().ClusterRoles().Update(context.TODO(), &clusterRole, metav1.UpdateOptions{}); err != nil {
			return err
		}
	}
	return nil
}
