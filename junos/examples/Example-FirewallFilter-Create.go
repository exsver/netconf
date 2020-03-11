package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	filter := junos.Filter{
		Name: "testFilter",
		Terms: []junos.Term{
			{
				Name: "DenyUDP",
				From: &junos.FilterFrom{
					Protocol: []string{"udp"},
				},
				Then: &junos.FilterThen{
					Count:   "udp-drop-counter",
					Discard: nil,
				},
			},
			{
				Name: "Accept",
				Then: &junos.FilterThen{
					Accept: true,
				},
			},
		},
	}

	err = device.EditConfig(filter.ConvertToConfig(false), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	commitResult, err := device.Commit()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("Commit Ok - %v", commitResult.OK)
}
