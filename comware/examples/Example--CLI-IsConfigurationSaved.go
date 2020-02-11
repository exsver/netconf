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

	ok, diff, err := sw.IsConfigurationSaved()
	if err != nil {
		log.Fatalf("%s", err)
	}

	if ok {
		fmt.Println("Configuration changes are saved.")
	} else {
		fmt.Printf("Configuration changes are not saved.\n %s", string(diff))
	}
}
