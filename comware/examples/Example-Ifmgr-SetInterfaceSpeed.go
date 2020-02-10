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
		err := sw.SetInterfaceSpeed(10, "100") //auto | 10 | 100 | 1000 | 1G | 10000 | 10G | 40000 | 40G
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
