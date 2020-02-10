package main

import (
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

	v := junos.Vlan{
		Name:        "test-vlan",
		Description: "test-vlan-descr",
		VlanID:      40,
	}

	err = sw.EditConfig(*v.ConvertToConfig(), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}
}
