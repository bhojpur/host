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

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func UpdateServiceAccountFromYaml(k8sClient *kubernetes.Clientset, serviceAccountYaml string) error {
	serviceAccount := v1.ServiceAccount{}

	if err := DecodeYamlResource(&serviceAccount, serviceAccountYaml); err != nil {
		return err
	}
	return retryTo(updateServiceAccount, k8sClient, serviceAccount, DefaultRetries, DefaultSleepSeconds)
}

func updateServiceAccount(k8sClient *kubernetes.Clientset, s interface{}) error {
	serviceAccount := s.(v1.ServiceAccount)
	if _, err := k8sClient.CoreV1().ServiceAccounts(metav1.NamespaceSystem).Create(context.TODO(), &serviceAccount, metav1.CreateOptions{}); err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return err
		}
		if _, err := k8sClient.CoreV1().ServiceAccounts(metav1.NamespaceSystem).Update(context.TODO(), &serviceAccount, metav1.UpdateOptions{}); err != nil {
			return err
		}
	}
	return nil
}
