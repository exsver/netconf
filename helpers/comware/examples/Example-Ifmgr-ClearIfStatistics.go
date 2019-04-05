package main

import (
	"fmt"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	err = sw.ClearIfStatistics(1)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("OK")
	}
}
