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
	"reflect"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func UpdateConfigMap(k8sClient *kubernetes.Clientset, configYaml []byte, configMapName string) (bool, error) {
	cfgMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: metav1.NamespaceSystem,
		},
		Data: map[string]string{
			configMapName: string(configYaml),
		},
	}
	updated := false
	// let's try to get the config map from k8s
	existingConfigMap, err := GetConfigMap(k8sClient, configMapName)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return updated, err
		}
		// the config map is not in k8s, I will create it and return updated=false
		if _, err := k8sClient.CoreV1().ConfigMaps(metav1.NamespaceSystem).Create(context.TODO(), cfgMap, metav1.CreateOptions{}); err != nil {
			return updated, err
		}
		return updated, nil
	}
	if !reflect.DeepEqual(existingConfigMap.Data, cfgMap.Data) {
		if _, err := k8sClient.CoreV1().ConfigMaps(metav1.NamespaceSystem).Update(context.TODO(), cfgMap, metav1.UpdateOptions{}); err != nil {
			return updated, err
		}
		updated = true
	}
	return updated, nil
}

func GetConfigMap(k8sClient *kubernetes.Clientset, configMapName string) (*v1.ConfigMap, error) {
	return k8sClient.CoreV1().ConfigMaps(metav1.NamespaceSystem).Get(context.TODO(), configMapName, metav1.GetOptions{})
}

func DeleteConfigMap(k8sClient *kubernetes.Clientset, configMapName string) error {
	return k8sClient.CoreV1().ConfigMaps(metav1.NamespaceSystem).Delete(context.TODO(), configMapName, metav1.DeleteOptions{})
}
