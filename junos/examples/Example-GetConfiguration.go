package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Messages()
	sw, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	//Format: text | set | xml | json
	config, _ := sw.GetConfiguration("xml")

	fmt.Printf("%s", config)
}
