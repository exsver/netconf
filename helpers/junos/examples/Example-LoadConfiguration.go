package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf/helpers/junos"
	"log"
)

func main() {
	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	config := `vlans {v109 {description test-vlan109;vlan-id 109;}}`
	err = sw.LoadConfigurationText(config, "merge")
	spew.Dump(err)

}
