package main

import (
	"fmt"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		err := sw.RestoreInterfaceDefaultConfiguration(10)
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
