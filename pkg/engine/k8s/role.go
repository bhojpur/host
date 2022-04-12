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

func UpdateRoleBindingFromYaml(k8sClient *kubernetes.Clientset, roleBindingYaml, namespace string) error {
	roleBinding := rbacv1.RoleBinding{}
	if err := DecodeYamlResource(&roleBinding, roleBindingYaml); err != nil {
		return err
	}
	roleBinding.Namespace = namespace
	return retryTo(updateRoleBinding, k8sClient, roleBinding, DefaultRetries, DefaultSleepSeconds)
}

func updateRoleBinding(k8sClient *kubernetes.Clientset, rb interface{}) error {
	roleBinding := rb.(rbacv1.RoleBinding)
	if _, err := k8sClient.RbacV1().RoleBindings(roleBinding.Namespace).Create(context.TODO(), &roleBinding, metav1.CreateOptions{}); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return err
		}
		if _, err := k8sClient.RbacV1().RoleBindings(roleBinding.Namespace).Update(context.TODO(), &roleBinding, metav1.UpdateOptions{}); err != nil {
			return err
		}
	}
	return nil
}

func UpdateRoleFromYaml(k8sClient *kubernetes.Clientset, roleYaml, namespace string) error {
	role := rbacv1.Role{}
	if err := DecodeYamlResource(&role, roleYaml); err != nil {
		return err
	}
	role.Namespace = namespace
	return retryTo(updateRole, k8sClient, role, DefaultRetries, DefaultSleepSeconds)
}

func updateRole(k8sClient *kubernetes.Clientset, r interface{}) error {
	role := r.(rbacv1.Role)
	if _, err := k8sClient.RbacV1().Roles(role.Namespace).Create(context.TODO(), &role, metav1.CreateOptions{}); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return err
		}
		if _, err := k8sClient.RbacV1().Roles(role.Namespace).Update(context.TODO(), &role, metav1.UpdateOptions{}); err != nil {
			return err
		}
	}
	return nil
}
