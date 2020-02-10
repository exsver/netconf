package main

import (
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	pfilter := comware.Pfilter{
		AppObjType:   1,         // 1 - interface, 2 - vlan, 3 - global.
		AppObjIndex:  8,         // For interface ifIndex.
		AppDirection: 1,         // Apply Direction: 1 - inbound, 2 - outbound.
		AppACLType:   1,         // 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined, 5 - default.
		AppACLGroup:  "testACL", // ACLName
	}
	err = sw.PfilterApply(&pfilter)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
