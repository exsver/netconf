package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	ok, err := sw.IsVlanInTrunk("Bridge-Aggregation1", 101)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("%v", ok)
}
