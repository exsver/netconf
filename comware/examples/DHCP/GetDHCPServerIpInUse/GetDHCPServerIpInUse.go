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

	// Getting IP List
	data, err := sw.GetDataDHCPServerIpInUse()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print
	for _, ip := range data.IPInUse {
		if ip.Type == 5 {
			fmt.Printf("VLAN: %v, Ipv4Address: %s, CID: %s, EndLease: %s\n", ip.VLANID, ip.Ipv4Address, ip.CID, ip.EndLease)
		}
	}
}
