package provision

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
	"io/ioutil"

	"github.com/bhojpur/host/pkg/core/provision/pkgaction"
)

func WithCustomScript(provisioner Provisioner, customScriptPath string) error {
	if provisioner == nil {
		return nil
	}

	if err := provisioner.SetHostname(provisioner.GetDriver().GetMachineName()); err != nil {
		return err
	}

	for _, pkg := range provisioner.GetPackages() {
		if err := provisioner.Package(pkg, pkgaction.Install); err != nil {
			return err
		}
	}

	customScriptContents, err := ioutil.ReadFile(customScriptPath)
	if err != nil {
		return fmt.Errorf("unable to read file %s: %v", customScriptPath, err)
	}

	if output, err := provisioner.SSHCommand(fmt.Sprintf("cat <<'OEOF' >/tmp/install_script.sh\n%s\nOEOF", string(customScriptContents))); err != nil {
		return fmt.Errorf("error uploading custom script: output: %s, error: %s", output, err)
	}
	if output, err := provisioner.SSHCommand("sudo sh /tmp/install_script.sh"); err != nil {
		return fmt.Errorf("error running custom script: output: %s, error: %s", output, err)
	}

	return nil
}
