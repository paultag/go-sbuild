go-sbuild
=========

```
package main

import (
	"fmt"
	"os"

	"pault.ag/go/sbuild"
)

func main() {
	sbuild := sbuild.NewSbuild("unstable", "unstable")
	sbuild.Verbose()
	sbuild.ArchAll()
	cmd, err := sbuild.BuildCommand("/some/path/somepackage.dsc")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("%s\n", cmd.Run())
}
```
