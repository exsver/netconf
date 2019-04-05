package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	dhcpSnooping, err := sw.GetDataDHCPSP()
	if err != nil {
		log.Fatalf("%s", err)
	}
	spew.Dump(dhcpSnooping)
}
