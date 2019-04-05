package main

import (
	"github.com/exsver/netconf/helpers/junos"
	"log"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	spew.Dump(sw.LoadConfigurationRolback(1))
}
