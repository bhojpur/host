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
	"k8s.io/gengo/namer"
)

func GroupInterfaceGo(group string, args *args.GeneratorArgs, customArgs *args2.CustomArgs) generator.Generator {
	return &interfaceGo{
		group:      group,
		args:       args,
		customArgs: customArgs,
		DefaultGen: generator.DefaultGen{
			OptionalName: "interface",
			OptionalBody: []byte(interfaceBody),
		},
	}
}

type interfaceGo struct {
	generator.DefaultGen

	group      string
	args       *args.GeneratorArgs
	customArgs *args2.CustomArgs
}

func (f *interfaceGo) Imports(context *generator.Context) []string {
	packages := Imports

	for gv := range f.customArgs.TypesByGroup {
		if gv.Group != f.group {
			continue
		}
		packages = append(packages, fmt.Sprintf("%s \"%s/controllers/%s/%s\"", gv.Version, f.customArgs.Package,
			groupPackageName(gv.Group, f.customArgs.Options.Groups[gv.Group].OutputControllerPackageName), gv.Version))
	}

	return packages
}

func (f *interfaceGo) Init(c *generator.Context, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "{{", "}}")
	sw.Do("type Interface interface {\n", nil)
	for gv := range f.customArgs.TypesByGroup {
		if gv.Group != f.group {
			continue
		}

		sw.Do("{{.upperVersion}}() {{.version}}.Interface\n", map[string]interface{}{
			"upperVersion": namer.IC(gv.Version),
			"version":      gv.Version,
		})
	}
	sw.Do("}\n", nil)

	if err := f.DefaultGen.Init(c, w); err != nil {
		return err
	}

	for gv := range f.customArgs.TypesByGroup {
		if gv.Group != f.group {
			continue
		}

		m := map[string]interface{}{
			"upperGroup":   upperLowercase(f.group),
			"upperVersion": namer.IC(gv.Version),
			"version":      gv.Version,
		}
		sw.Do("\nfunc (g *group) {{.upperVersion}}() {{.version}}.Interface {\n", m)
		sw.Do("return {{.version}}.New(g.controllerFactory)\n", m)
		sw.Do("}\n", m)
	}

	return sw.Error()
}

var interfaceBody = `
type group struct {
	controllerFactory controller.SharedControllerFactory
}

// New returns a new Interface.
func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &group{
		controllerFactory: controllerFactory,
	}
}
`
