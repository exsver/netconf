package main

import (
	"log"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatal(err)
	}

	filter := junos.Filter{
		Name: "filterWithPrefixList",
		Terms: []junos.Term{
			// set firewall family inet filter filterWithPrefixList term whitelist from destination-address 10.100.100.100/32
			// set firewall family inet filter filterWithPrefixList term whitelist from source-prefix-list Whitelist
			// set firewall family inet filter filterWithPrefixList term whitelist then accept
			{
				Name: "whitelist",
				From: &junos.FilterFrom{
					DestinationAddress: []junos.FilterDestinationAddress{{Name: "10.100.100.100"}},
					SourcePrefixList:   []junos.FilterSourcePrefixList{{Name: "Whitelist"}},
				},
				Then: &junos.FilterThen{
					Accept: true,
				},
			},
			// set firewall family inet filter filterWithPrefixList term Discard from destination-address 10.100.100.100/32
			// set firewall family inet filter filterWithPrefixList term Discard then discard
			{
				Name: "Discard",
				From: &junos.FilterFrom{
					DestinationAddress: []junos.FilterDestinationAddress{{Name: "10.100.100.100"}},
				},
				Then: &junos.FilterThen{
					Discard: &junos.FilterThenDiscard{},
				},
			},
		},
	}

	// set policy-options prefix-list Whitelist 10.0.0.42/32
	// set policy-options prefix-list Whitelist 10.0.0.73/32
	prefixList := junos.PrefixList{
		Name: "Whitelist",
		PrefixListItems: []junos.PrefixListItem{
			{Name: "10.0.0.42/32"},
			{Name: "10.0.0.73/32"},
		},
	}

	var config junos.Config

	filter.AppendToConfig(true, &config)
	prefixList.AppendToConfig(&config)

	err = device.EditConfig(&config, "merge")
	if err != nil {
		log.Fatal(err)
	}
}
