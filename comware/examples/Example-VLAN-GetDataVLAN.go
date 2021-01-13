package main

import (
	"github.com/davecgh/go-spew/spew"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	data, err := sw.GetDataVLAN()
	spew.Dump(data, err)
}
