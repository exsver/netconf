package main

import (
	"log"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

// Run ./Example-FirewallFilterTermFrom-Add.go first

func main() {
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	term := junos.Term{
		Name: "count",
		From: &junos.FilterFrom{
			DestinationAddress: []junos.FilterDestinationAddress{
				{
					NetconfConfigOperation: "remove",
					Name:             "10.0.0.2/32",
				},
			},
		},
	}

	err = device.EditConfig(term.ConvertToConfig("inet", "testFilterCounter"), "none")
	if err != nil {
		log.Fatalf("%s", err)
	}
}
