package main

import (
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

	pfilter := comware.Pfilter{
		AppObjType:   comware.PFilterAppObjTypeInterface,   // 1 - interface, 2 - vlan, 3 - global.
		AppObjIndex:  8,                                    // For interface ifIndex.
		AppDirection: comware.PFilterApplyDirectionInbound, // Apply Direction: 1 - inbound, 2 - outbound.
		AppACLType:   comware.ACLGroupTypeIPv4,             // 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined, 5 - default.
		AppACLGroup:  "testACL",                            // ACLName
	}

	err = sw.PfilterApply(&pfilter)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
