package main

import (
	"fmt"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
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
