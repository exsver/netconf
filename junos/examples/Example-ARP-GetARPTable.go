package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Messages()

	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	ARPTable, err := sw.GetARPTableInformation()

	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, ARP := range ARPTable.ARPTableEntries {
		fmt.Printf("%s %s %s\n", ARP.MACAddress, ARP.IPAddress, ARP.InterfaceName)
	}

	fmt.Printf("Total entries: %v", ARPTable.ARPEntryCount)
}
