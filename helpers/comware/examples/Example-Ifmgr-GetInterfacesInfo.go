package main

import (
	"log"

	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	interfaces, err := sw.GetInterfacesInfo([]string{`<IfIndex>10</IfIndex>`})
	for _, v := range interfaces {
		spew.Dump(v, err)
	}
}
