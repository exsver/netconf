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

	// interface GigabitEthernet1/0/10
	//   port link-type trunk
	err = sw.SetInterfaceLinkType(10, "trunk") // access | trunk | hybrid
	if err != nil {
		log.Fatalf("%s", err)
	}

	// interface GigabitEthernet1/0/10
	//   undo port trunk permit vlan 1
	//   port trunk permit vlan 300 301 302
	err = sw.SetTrunkInterfaceVlans(10, []int{300, 301, 302}, 0) // ifIndex, vlans, pvid
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("OK\n")
}
