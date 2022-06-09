package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Add binding
	// interface GigabitEthernet1/0/8
	//   ip source binding ip-address 10.10.10.11 mac-address AABB-CCDD-EEFF
	sw.AddIPSourceBinding("8", "10.10.10.11", "AA-BB-CC-DD-EE-FF", "220")

	// interface GigabitEthernet1/0/8
	//   ip verify source ip-address mac-address
	err = sw.AddIpVerifySource(10, true, true)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Get and print all bindings
	b, _ := sw.GetIPSourceBindings()
	for _, v := range b {
		fmt.Printf("Port: %s IP: %s MAC: %s VLAN: %s\n", v.IfIndex, v.Ipv4Address, v.MacAddress, v.VLANID)
	}

	// Delete binding
	// interface GigabitEthernet1/0/8
	//   undo ip source binding ip-address 10.10.10.11 mac-address AABB-CCDD-EEFF
	sw.DeleteIPSourceBinding("8", "10.10.10.11", "AA-BB-CC-DD-EE-FF", "220")
}
