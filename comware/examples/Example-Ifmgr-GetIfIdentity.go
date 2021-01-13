package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Messages()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	ifIdentity, err := sw.GetIfIdentity()
	if err != nil {
		log.Fatalf("%s", err)
	}

	for ifIndex, iface := range ifIdentity {
		fmt.Printf("Index: %v, Name: %s, Description: %s\n", ifIndex, iface.AbbreviatedName, iface.Description)
	}
}
