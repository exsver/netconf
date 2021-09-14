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
	sw.AddIPSourceBinding("8", "10.10.10.11", "AA-BB-CC-DD-EE-FF", "220")

	// Get and print all bindings
	b, _ := sw.GetIPSourceBindings()
	for _, v := range b {
		fmt.Printf("Port: %s IP: %s MAC: %s VLAN: %s\n", v.IfIndex, v.Ipv4Address, v.MacAddress, v.VLANID)
	}

	// Delete binding
	sw.DeleteIPSourceBinding("8", "10.10.10.11", "AA-BB-CC-DD-EE-FF", "220")
}
