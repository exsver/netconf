package main

import (
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = sw.SetInterfaceSpeed(10, "100") // auto | 10 | 100 | 1000 | 1G | 10000 | 10G | 40000 | 40G
	if err != nil {
		log.Fatalf("%s", err)
	}
}
