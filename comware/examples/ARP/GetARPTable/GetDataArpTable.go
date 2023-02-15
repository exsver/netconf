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

	// Getting data
	data, err := sw.GetDataARPTable()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print
	for _, arpEntry := range data.ArpEntries {
		fmt.Printf("IfIndex: %v, PortIndex: %v, VLANID: %v, IP: %s, MAC: %s", arpEntry.IfIndex, arpEntry.PortIndex, arpEntry.VLANID, arpEntry.Ipv4Address, arpEntry.MacAddress)
	}
}
