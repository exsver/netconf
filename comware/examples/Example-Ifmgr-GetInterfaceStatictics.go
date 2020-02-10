package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		ifaces, err := sw.GetEthPortStatistics()
		spew.Dump(ifaces, err)
	}
}
