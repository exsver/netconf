package main

import (
	"github.com/davecgh/go-spew/spew"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	dhcpInfo, err := sw.GetDataDHCPServer()
	if err != nil {
		log.Fatalf("%s", err)
	}

	spew.Dump(dhcpInfo)
}
