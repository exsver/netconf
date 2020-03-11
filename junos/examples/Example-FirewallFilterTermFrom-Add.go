package main

import (
	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
	"log"
)

/*
set firewall family inet filter testFilterCounter term count from destination-address 10.0.0.1/32
set firewall family inet filter testFilterCounter term count then count testCounter

test-mx80# show firewall family inet filter testFilterCounter
term count {
    from {
        destination-address {
            10.0.0.1/32;
        }
    }
    then count testCounter;
}
*/

func main() {
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// set firewall family inet filter testFilterCounter term count from destination-address 10.0.0.2/32
	term := junos.Term{
		Name: "count",
		From: &junos.FilterFrom{
			DestinationAddress: []junos.FilterDestinationAddress{
				{Name: "10.0.0.2/32"},
			},
		},
	}

	err = device.EditConfig(term.ConvertToConfig(true, "testFilterCounter"), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

}
