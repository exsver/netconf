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

	// Getting interfaces info
	ifIdentity, err := sw.GetIfIdentity()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Getting ARP table
	data, err := sw.GetDataARPTable()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print
	for _, arpEntry := range data.ArpEntries {
		fmt.Printf("IfIndex: %v, IfName: %s, PortIndex: %v, PortName: %s, VLANID: %v, IP: %s, MAC: %s",
			arpEntry.IfIndex,
			ifIdentity[arpEntry.IfIndex].AbbreviatedName,
			arpEntry.PortIndex,
			ifIdentity[arpEntry.PortIndex].AbbreviatedName,
			arpEntry.VLANID,
			arpEntry.Ipv4Address,
			arpEntry.MacAddress,
		)
	}
}
