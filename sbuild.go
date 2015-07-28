/* {{{ Copyright (c) Paul R. Tagliamonte <paultag@debian.org>, 2015
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE. }}} */

package sbuild

import (
	"fmt"
	"os/exec"
)

type Sbuild struct {
	args []string
}

func (sbuild Sbuild) Partial(args ...string) Sbuild {
	sbuild.args = append(sbuild.args, args...)
	return sbuild
}

func (sbuild *Sbuild) AddFlag(flag string) {
	sbuild.args = append(sbuild.args, flag)
}

func (sbuild *Sbuild) AddArgument(key, value string) {
	sbuild.args = append(sbuild.args, fmt.Sprintf("--%s=%s", key, value))
}

func (sbuild *Sbuild) BuildCommand(dsc string) (*exec.Cmd, error) {
	s := sbuild.Partial(dsc)
	cmd := exec.Command("sbuild", s.args...)
	return cmd, nil
}

func (sbuild *Sbuild) Build(dsc string) (*exec.Cmd, error) {
	cmd, err := sbuild.BuildCommand(dsc)
	if err != nil {
		return nil, err
	}
	return cmd, cmd.Start()
}

func NewSbuild(chroot string, suite string, args ...string) Sbuild {
	sbuild := Sbuild{
		args: append([]string{"-c", chroot, "-d", suite}, args...),
	}
	return sbuild
}

// vim: foldmethod=marker
