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
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

const (
	HostnameLabel             = "kubernetes.io/hostname"
	InternalAddressAnnotation = "bke.bhojpur.net/internal-ip"
	ExternalAddressAnnotation = "bke.bhojpur.net/external-ip"
	AWSCloudProvider          = "aws"
	MaxRetries                = 5
	RetryInterval             = 5
)

func DeleteNode(k8sClient *kubernetes.Clientset, nodeName, cloudProvider string) error {
	// If cloud provider is configured, the node name can be set by the cloud provider, which can be different from the original node name
	if cloudProvider != "" {
		node, err := GetNode(k8sClient, nodeName)
		if err != nil {
			return err
		}
		nodeName = node.Name
	}
	return k8sClient.CoreV1().Nodes().Delete(context.TODO(), nodeName, metav1.DeleteOptions{})
}

func GetNodeList(k8sClient *kubernetes.Clientset) (*v1.NodeList, error) {
	return k8sClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
}

func GetNode(k8sClient *kubernetes.Clientset, nodeName string) (*v1.Node, error) {
	var listErr error
	for retries := 0; retries < MaxRetries; retries++ {
		logrus.Debugf("Checking node list for node [%v], try #%v", nodeName, retries+1)
		nodes, err := GetNodeList(k8sClient)
		if err != nil {
			listErr = err
			time.Sleep(time.Second * RetryInterval)
			continue
		}
		// reset listErr back to nil
		listErr = nil
		for _, node := range nodes.Items {
			if strings.ToLower(node.Labels[HostnameLabel]) == strings.ToLower(nodeName) {
				return &node, nil
			}
		}
		time.Sleep(time.Second * RetryInterval)
	}
	if listErr != nil {
		return nil, listErr
	}
	return nil, apierrors.NewNotFound(schema.GroupResource{}, nodeName)
}

func CordonUncordon(k8sClient *kubernetes.Clientset, nodeName string, cordoned bool) error {
	updated := false
	for retries := 0; retries < MaxRetries; retries++ {
		node, err := GetNode(k8sClient, nodeName)
		if err != nil {
			logrus.Debugf("Error getting node %s: %v", nodeName, err)
			// no need to retry here since GetNode already retries
			return err
		}
		if node.Spec.Unschedulable == cordoned {
			logrus.Debugf("Node %s is already cordoned: %v", nodeName, cordoned)
			return nil
		}
		node.Spec.Unschedulable = cordoned
		_, err = k8sClient.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
		if err != nil {
			logrus.Debugf("Error setting cordoned state for node %s: %v", nodeName, err)
			time.Sleep(time.Second * RetryInterval)
			continue
		}
		updated = true
	}
	if !updated {
		return fmt.Errorf("Failed to set cordonded state for node: %s", nodeName)
	}
	return nil
}

func IsNodeReady(node v1.Node) bool {
	nodeConditions := node.Status.Conditions
	for _, condition := range nodeConditions {
		if condition.Type == v1.NodeReady && condition.Status == v1.ConditionTrue {
			return true
		}
	}
	return false
}

func RemoveTaintFromNodeByKey(k8sClient *kubernetes.Clientset, nodeName, taintKey string) error {
	updated := false
	var err error
	var node *v1.Node
	for retries := 0; retries <= 5; retries++ {
		node, err = GetNode(k8sClient, nodeName)
		if err != nil {
			if apierrors.IsNotFound(err) {
				logrus.Debugf("[hosts] Can't find node by name [%s]", nodeName)
				return nil
			}
			return err
		}
		foundTaint := false
		for i, taint := range node.Spec.Taints {
			if taint.Key == taintKey {
				foundTaint = true
				node.Spec.Taints = append(node.Spec.Taints[:i], node.Spec.Taints[i+1:]...)
				break
			}
		}
		if !foundTaint {
			return nil
		}
		_, err = k8sClient.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
		if err != nil {
			logrus.Debugf("Error updating node [%s] with new set of taints: %v", node.Name, err)
			time.Sleep(time.Second * 5)
			continue
		}
		updated = true
		break
	}
	if !updated {
		return fmt.Errorf("Timeout waiting for node [%s] to be updated with new set of taints: %v", node.Name, err)
	}
	return nil
}

func SyncNodeLabels(node *v1.Node, toAddLabels, toDelLabels map[string]string) {
	oldLabels := map[string]string{}
	if node.Labels == nil {
		node.Labels = map[string]string{}
	}

	for k, v := range node.Labels {
		oldLabels[k] = v
	}

	// Delete Labels
	for key := range toDelLabels {
		if _, ok := node.Labels[key]; ok {
			delete(node.Labels, key)
		}
	}
	// ADD Labels
	for key, value := range toAddLabels {
		node.Labels[key] = value
	}
}

func SyncNodeTaints(node *v1.Node, toAddTaints, toDelTaints []string) {
	// Add taints to node
	for _, taintStr := range toAddTaints {
		if isTaintExist(toTaint(taintStr), node.Spec.Taints) {
			continue
		}
		node.Spec.Taints = append(node.Spec.Taints, toTaint(taintStr))
	}
	// Remove Taints from node
	for _, taintStr := range toDelTaints {
		node.Spec.Taints = delTaintFromList(node.Spec.Taints, toTaint(taintStr))
	}
}

func isTaintExist(taint v1.Taint, taintList []v1.Taint) bool {
	for _, t := range taintList {
		if t.Key == taint.Key && t.Value == taint.Value && t.Effect == taint.Effect {
			return true
		}
	}
	return false
}

func toTaint(taintStr string) v1.Taint {
	taintStruct := strings.Split(taintStr, "=")
	tmp := strings.Split(taintStruct[1], ":")
	key := taintStruct[0]
	value := tmp[0]
	effect := v1.TaintEffect(tmp[1])
	return v1.Taint{
		Key:    key,
		Value:  value,
		Effect: effect,
	}
}

func SetNodeAddressesAnnotations(node *v1.Node, internalAddress, externalAddress string) {
	currentExternalAnnotation := node.Annotations[ExternalAddressAnnotation]
	currentInternalAnnotation := node.Annotations[ExternalAddressAnnotation]
	if currentExternalAnnotation == externalAddress && currentInternalAnnotation == internalAddress {
		return
	}
	node.Annotations[ExternalAddressAnnotation] = externalAddress
	node.Annotations[InternalAddressAnnotation] = internalAddress
}

func delTaintFromList(l []v1.Taint, t v1.Taint) []v1.Taint {
	r := []v1.Taint{}
	for _, i := range l {
		if i == t {
			continue
		}
		r = append(r, i)
	}
	return r
}
