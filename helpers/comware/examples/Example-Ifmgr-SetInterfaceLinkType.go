package main

import (
	"fmt"
	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		err := sw.SetInterfaceLinkType(10, "hybrid") //access | trunk | hybrid
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
