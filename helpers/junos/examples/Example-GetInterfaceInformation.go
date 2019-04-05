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
	i, err := sw.GetInterfaceInformation()
	spew.Dump(i, err)
}
