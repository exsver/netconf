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

	// Create new Port Mirroring group (remote destination)
	err = sw.NewMirroringGroupRemoteDestination(1, 12, 8)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
