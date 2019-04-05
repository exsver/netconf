package main

import (
	"fmt"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/junos"
	"log"
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
