package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/junos"
	"log"
)

func main() {
	netconf.LogLevel.Messages()
	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	data, err := sw.GetConfig("candidate", "interfaces")
	// Source: running | candidate
	// Subtree: system | interfaces | forwarding-options | routing-options | routing-options/static | protocols | protocols/bgp | vlans
	spew.Dump(data, err)

}
