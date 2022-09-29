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
	netconf.LogLevel.Verbose()

	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	data, err := sw.GetDataVLANs()
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, vlan := range data.VLANs {
		fmt.Printf("ID: %v, Name: '%s', Description: '%s', AccessPortList: '%s', UntaggedPortList: '%s', TaggedPortList: '%s'\n",
			vlan.ID,
			vlan.Name,
			vlan.Description,
			vlan.AccessPortList,
			vlan.UntaggedPortList,
			vlan.TaggedPortList,
		)
	}
}
