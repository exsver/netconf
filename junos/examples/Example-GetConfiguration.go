package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
)

func main() {
	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	// Format: text | set | xml | json
	config, err := sw.GetConfiguration("text")
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("%s", config)
}
