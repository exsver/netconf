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
		err := sw.SetTrunkInterfaceVlans(10, []int{300, 301, 302}, 0)
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
