package main

import (
	"fmt"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	netconf.LogLevel.Messages()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	ifIdentity, err := sw.GetIfIdentity()
	if err != nil {
		log.Fatalf("%s", err)
	}
	for _, iface := range ifIdentity {
		fmt.Printf("%s %s\n", iface.AbbreviatedName, iface.Description)
	}
}
