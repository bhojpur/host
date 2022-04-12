package schema

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
	"reflect"
	"strings"

	v3 "github.com/bhojpur/host/pkg/apis/cluster.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/types"
	m "github.com/bhojpur/host/pkg/core/types/mapper"
	"github.com/bhojpur/host/pkg/schemas/factory"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
)

var (
	Version = types.APIVersion{
		Version:          "v3",
		Group:            "cluster.bhojpur.net",
		Path:             "/v3/cluster",
		SubContext:       true,
		SubContextSchema: "/v3/schemas/cluster",
	}

	Schemas = factory.Schemas(&Version).
		Init(namespaceTypes).
		Init(persistentVolumeTypes).
		Init(storageClassTypes).
		Init(tokens).
		Init(apiServiceTypes)
)

func namespaceTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.NamespaceSpec{},
			&m.Drop{Field: "finalizers"},
		).
		AddMapperForType(&Version, v1.Namespace{},
			&m.AnnotationField{Field: "description"},
			&m.AnnotationField{Field: "projectId"},
			&m.AnnotationField{Field: "resourceQuota", Object: true},
			&m.AnnotationField{Field: "containerDefaultResourceLimit", Object: true},
			&m.Drop{Field: "status"},
		).
		MustImport(&Version, NamespaceResourceQuota{}).
		MustImport(&Version, ContainerResourceLimit{}).
		MustImport(&Version, v1.Namespace{}, struct {
			Description                   string `json:"description"`
			ProjectID                     string `bhojpur:"type=reference[/v3/schemas/project],noupdate"`
			ResourceQuota                 string `json:"resourceQuota,omitempty" bhojpur:"type=namespaceResourceQuota"`
			ContainerDefaultResourceLimit string `json:"containerDefaultResourceLimit,omitempty" bhojpur:"type=containerResourceLimit"`
		}{}).
		MustImport(&Version, NamespaceMove{}).
		MustImportAndCustomize(&Version, v1.Namespace{}, func(schema *types.Schema) {
			schema.ResourceActions["move"] = types.Action{
				Input: "namespaceMove",
			}
		})
}

func persistentVolumeTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.PersistentVolume{},
			&m.AnnotationField{Field: "description"},
		).
		AddMapperForType(&Version, v1.HostPathVolumeSource{},
			m.Move{From: "type", To: "kind"},
			m.Enum{
				Options: []string{
					"DirectoryOrCreate",
					"Directory",
					"FileOrCreate",
					"File",
					"Socket",
					"CharDevice",
					"BlockDevice",
				},
				Field: "kind",
			},
		).
		MustImport(&Version, v1.PersistentVolumeSpec{}, struct {
			StorageClassName *string `json:"storageClassName,omitempty" bhojpur:"type=reference[storageClass]"`
		}{}).
		MustImport(&Version, v1.PersistentVolume{}, struct {
			Description string `json:"description"`
		}{}).
		MustImportAndCustomize(&Version, v1.PersistentVolume{}, func(schema *types.Schema) {
			schema.MustCustomizeField("name", func(field types.Field) types.Field {
				field.Type = "hostname"
				field.Nullable = false
				field.Required = true
				return field
			})
			schema.MustCustomizeField("volumeMode", func(field types.Field) types.Field {
				field.Update = false
				return field
			})
			// All fields of PersistentVolumeSource are immutable
			val := reflect.ValueOf(v1.PersistentVolumeSource{})
			for i := 0; i < val.Type().NumField(); i++ {
				if tag, ok := val.Type().Field(i).Tag.Lookup("json"); ok {
					name := strings.Split(tag, ",")[0]
					schema.MustCustomizeField(name, func(field types.Field) types.Field {
						field.Update = false
						return field
					})
					pvSchema := schemas.Schema(&Version, val.Type().Field(i).Type.String()[4:])
					for name := range pvSchema.ResourceFields {
						pvSchema.MustCustomizeField(name, func(field types.Field) types.Field {
							field.Update = false
							return field
						})
					}
				}
			}
		})
}

func storageClassTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, storagev1.StorageClass{},
			&m.AnnotationField{Field: "description"},
		).
		MustImport(&Version, storagev1.StorageClass{}, struct {
			Description   string `json:"description"`
			ReclaimPolicy string `json:"reclaimPolicy,omitempty" bhojpur:"type=enum,options=Recycle|Delete|Retain"`
		}{})
}

func tokens(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImportAndCustomize(&Version, v3.ClusterAuthToken{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{}
		}).
		MustImportAndCustomize(&Version, v3.ClusterUserAttribute{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{}
		})
}

func apiServiceTypes(Schemas *types.Schemas) *types.Schemas {
	return Schemas.
		AddMapperForType(&Version, apiregistrationv1.APIService{},
			&m.Embed{Field: "status"},
		).
		MustImport(&Version, apiregistrationv1.APIService{})
}
