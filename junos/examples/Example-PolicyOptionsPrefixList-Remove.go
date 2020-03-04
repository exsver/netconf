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

	// delete policy-options prefix-list rfc1918
	prefixList := junos.PrefixList{
		Name:                   "rfc1918",
		NetconfConfigOperation: "remove",
	}

	err = device.EditConfig(prefixList.ConvertToConfig(), "none")
	if err != nil {
		fmt.Println("Fail")
	} else {
		fmt.Println("OK")
	}
}
