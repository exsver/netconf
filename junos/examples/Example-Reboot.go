package main

import (
	"log"

	"github.com/exsver/netconf/junos"
)

func main() {
	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	sw.Reboot()
}
