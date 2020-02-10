package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	ok, err := sw.IsVlanInTrunk("Bridge-Aggregation1", 101)
	if err != nil {
		log.Fatalf("%s", err)
	}
	spew.Dump(ok)
}
