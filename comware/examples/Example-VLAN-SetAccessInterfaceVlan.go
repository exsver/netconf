package main

import (
	"fmt"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		err := sw.SetAccessInterfaceVlan(10, 300)
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
