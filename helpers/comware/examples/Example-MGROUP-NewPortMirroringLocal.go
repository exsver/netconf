package main

import (
	"log"

	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// add 2 ports (with indexes 10 and 11) as source ports
	sourcePorts := []comware.PortMirroringSourcePort{
		{
			IfIndex:   10, // Interface index
			Direction: 3,  // Direction of source port: 1—Inbound, 2—Outbound, 3—Both.
		},
		{
			IfIndex:   11, // Interface index
			Direction: 3,  // Direction of source port: 1—Inbound, 2—Outbound, 3—Both.
		},
	}

	// Create new Port Mirroring group (local)
	err = sw.NewMirroringGroupLocal(1, 28, sourcePorts)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
