package generators

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
	"fmt"
	"io"

	args2 "github.com/bhojpur/host/pkg/common/controller-gen/args"
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
)

func FactoryGo(group string, args *args.GeneratorArgs, customArgs *args2.CustomArgs) generator.Generator {
	return &factory{
		group:      group,
		args:       args,
		customArgs: customArgs,
		DefaultGen: generator.DefaultGen{
			OptionalName: "factory",
			OptionalBody: []byte(factoryBody),
		},
	}
}

type factory struct {
	generator.DefaultGen

	group      string
	args       *args.GeneratorArgs
	customArgs *args2.CustomArgs
}

func (f *factory) Imports(*generator.Context) []string {
	imports := Imports

	for gv, types := range f.customArgs.TypesByGroup {
		if f.group == gv.Group && len(types) > 0 {
			imports = append(imports,
				fmt.Sprintf("%s \"%s\"", gv.Version, types[0].Package))
		}
	}

	return imports
}

func (f *factory) Init(c *generator.Context, w io.Writer) error {
	if err := f.DefaultGen.Init(c, w); err != nil {
		return err
	}

	sw := generator.NewSnippetWriter(w, c, "{{", "}}")
	m := map[string]interface{}{
		"groupName": upperLowercase(f.group),
	}

	sw.Do("\n\nfunc (c *Factory) {{.groupName}}() Interface {\n", m)
	sw.Do("	return New(c.ControllerFactory())\n", m)
	sw.Do("}\n\n", m)

	return sw.Error()
}

var factoryBody = `
type Factory struct {
	*generic.Factory
}

func NewFactoryFromConfigOrDie(config *rest.Config) *Factory {
	f, err := NewFactoryFromConfig(config)
	if err != nil {
		panic(err)
	}
	return f
}

func NewFactoryFromConfig(config *rest.Config) (*Factory, error) {
	return NewFactoryFromConfigWithOptions(config, nil)
}

func NewFactoryFromConfigWithNamespace(config *rest.Config, namespace string) (*Factory, error) {
	return NewFactoryFromConfigWithOptions(config, &FactoryOptions{
		Namespace: namespace,
	})
}

type FactoryOptions = generic.FactoryOptions

func NewFactoryFromConfigWithOptions(config *rest.Config, opts *FactoryOptions) (*Factory, error) {
	f, err := generic.NewFactoryFromConfigWithOptions(config, opts)
	return &Factory{
		Factory: f,
	}, err
}

func NewFactoryFromConfigWithOptionsOrDie(config *rest.Config, opts *FactoryOptions) *Factory {
    f, err := NewFactoryFromConfigWithOptions(config, opts)
	if err != nil {
		panic(err)
	}
	return f
}

`
