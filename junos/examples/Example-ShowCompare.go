package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
)

func main() {
	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	output, err := device.CompareConfigurationRollback(0, "text")
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("Configuration diff:\n %s", bytes.TrimSpace(output))
}
