package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()
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
		fmt.Printf("IfIndex: %v, IfType, %v, Name: %s, Abbreviated Name: %s, Description: %s\n", ifIndex, iface.IfType, iface.Name, iface.AbbreviatedName, iface.Description)
	}
}
