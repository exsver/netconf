package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Messages()

	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Source: running | candidate
	// Subtree: system | interfaces | forwarding-options | routing-options | routing-options/static | protocols | protocols/bgp | vlans | firewall | ...
	data, err := sw.GetConfig("candidate", "")
	if err != nil {
		log.Fatalf("%s", err)
	}

	spew.Dump(data)
}
