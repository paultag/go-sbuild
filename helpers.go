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

func (sbuild *Sbuild) Verbose() {
	sbuild.AddFlag("-v")
}

func (sbuild *Sbuild) Quiet() {
	sbuild.AddFlag("-q")
}

func (sbuild *Sbuild) Source() {
	sbuild.AddFlag("-s")
}

func (sbuild *Sbuild) NoSource() {
	sbuild.AddFlag("--no-source")
}

func (sbuild *Sbuild) ArchAll() {
	sbuild.AddFlag("-A")
}

func (sbuild *Sbuild) NoArchAll() {
	sbuild.AddFlag("--no-arch-all")
}

func (sbuild *Sbuild) Arch(arch string) {
	sbuild.AddArgument("arch", arch)
}

func (sbuild *Sbuild) HostArch(arch string) {
	sbuild.AddArgument("host", arch)
}

func (sbuild *Sbuild) BuildArch(arch string) {
	sbuild.AddArgument("build", arch)
}

func (sbuild *Sbuild) BuildDepResolver(resolver string) {
	sbuild.AddArgument("build-dep-resolver", resolver)
}

// vim: foldmethod=marker
