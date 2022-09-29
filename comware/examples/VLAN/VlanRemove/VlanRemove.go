package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

// Run examples/VLAN/VlanCreate/VlanCreate.go first

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	vlans := &comware.VLANs{
		VLANs: []comware.VLANID{
			{ID: 300},
			{ID: 301},
		},
	}

	err = sw.VlanRemove(vlans)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("Ok")
}
