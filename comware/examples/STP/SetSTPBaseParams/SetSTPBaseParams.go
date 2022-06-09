package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf/comware"
)

func main() {
	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	data, err := sw.GetDataSTP()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print running configuration
	spew.Dump(data.Base)

	// Chaging parameters
	config := data.Base
	config.Mode = comware.STPModeMSTP
	config.HelloTime = 7

	err = sw.SetSTPBaseParams(*config)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("OK\n")
}
