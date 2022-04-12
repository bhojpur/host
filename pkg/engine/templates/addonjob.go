package templates

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

const AddonJobTemplate = `
{{- $addonName := .AddonName }}
{{- $nodeName := .NodeName }}
{{- $image := .Image }}
{{- $OSLabel := .OSLabel }}
apiVersion: batch/v1
kind: Job
metadata:
{{- if eq .DeleteJob "true" }}
  name: {{$addonName}}-delete-job
{{- else }}
  name: {{$addonName}}-deploy-job
{{- end }}
  namespace: kube-system
spec:
  backoffLimit: 10
  template:
    metadata:
       name: bke-deploy
    spec:
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:
              nodeSelectorTerms:
                - matchExpressions:
                  - key: {{$OSLabel}}
                    operator: NotIn
                    values:
                      - windows
        tolerations:
        - operator: Exists
        hostNetwork: true
        serviceAccountName: bke-job-deployer
        nodeName: {{$nodeName}}
        containers:
          {{- if eq .DeleteJob "true" }}
          - name: {{$addonName}}-delete-pod
          {{- else }}
          - name: {{$addonName}}-pod
          {{- end }}
            image: {{$image}}
            {{- if eq .DeleteJob "true" }}
            command: ["/bin/sh"]
            args: ["-c" ,"kubectl get --ignore-not-found=true -f /etc/config/{{$addonName}}.yaml -o custom-columns=NAME:.metadata.name,NAMESPACE:.metadata.namespace,KIND:.kind --no-headers | while read name namespace kind; do if [ \"x${namespace}\" = \"x<none>\" ]; then kubectl delete $kind $name; else kubectl -n $namespace delete $kind $name; fi; done"]
            {{- else }}
            command: [ "kubectl", "apply", "-f" , "/etc/config/{{$addonName}}.yaml"]
            {{- end }}
            volumeMounts:
            - name: config-volume
              mountPath: /etc/config
        volumes:
          - name: config-volume
            configMap:
              # Provide the name of the ConfigMap containing the files you want
              # to add to the container
              name: {{$addonName}}
              items:
                - key: {{$addonName}}
                  path: {{$addonName}}.yaml
        restartPolicy: Never`
