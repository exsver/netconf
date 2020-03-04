package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
)

func main() {
	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// set policy-options prefix-list rfc1918 10.0.0.0/8
	// set policy-options prefix-list rfc1918 172.16.0.0/12
	// set policy-options prefix-list rfc1918 192.168.0.0/16
	prefixList := junos.PrefixList{
		Name: "rfc1918",
		PrefixListItems: []junos.PrefixListItem{
			{Name: "10.0.0.0/8"},
			{Name: "172.16.0.0/12"},
			{Name: "192.168.0.0/16"},
		},
	}

	err = device.EditConfig(prefixList.ConvertToConfig(), "merge")
	if err != nil {
		fmt.Println("Fail")
	} else {
		fmt.Println("OK")
	}
}
