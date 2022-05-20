package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	config := comware.STPBase{
		Mode:           0,
		TcThreshold:    0,
		PathCostMethod: 0,
		HelloTime:      0,
		MaxHops:        0,
		MaxAge:         0,
		ForwardDelay:   0,
		TcSnooping:     false,
		DigestSnooping: false,
		BPDUProtect:    false,
		TcProtect:      false,
		Enable:         false,
	}

	err = sw.SetSTPBaseParams(config)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("OK\n")
}
