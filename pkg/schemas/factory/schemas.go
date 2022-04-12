package factory

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
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/factory"
	m "github.com/bhojpur/host/pkg/core/types/mapper"
	"github.com/bhojpur/host/pkg/schemas/mapper"
)

func Schemas(version *types.APIVersion) *types.Schemas {
	schemas := factory.Schemas(version)
	baseFunc := schemas.DefaultMappers
	schemas.DefaultMappers = func() []types.Mapper {
		mappers := append([]types.Mapper{
			&mapper.Status{},
		}, baseFunc()...)
		mappers = append(mappers, &m.Scope{
			If: types.NamespaceScope,
			Mappers: []types.Mapper{
				&mapper.NamespaceIDMapper{},
			},
		}, &mapper.NamespaceReference{
			VersionPath: "/v3/project",
		})
		return mappers
	}
	basePostFunc := schemas.DefaultPostMappers
	schemas.DefaultPostMappers = func() []types.Mapper {
		return append(basePostFunc(), &mapper.Creator{})
	}
	return schemas
}
