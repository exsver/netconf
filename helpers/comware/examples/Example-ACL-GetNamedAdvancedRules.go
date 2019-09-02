package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	data, err := sw.GetListOfIPv4NamedAdvanceRules()
	if err != nil {
		log.Fatalf("%s", err)
	}
	spew.Dump(data)
}

