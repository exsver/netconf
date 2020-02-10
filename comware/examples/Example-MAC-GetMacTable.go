package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	// netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	// create a filter to get mac-addresses with VLANID 99 and PortIndex 1
	filters := []comware.XMLFilter{
		{
			Key:      "VLANID",
			Value:    "99",
			IsRegExp: false,
		},
		{
			Key:      "PortIndex",
			Value:    "1",
			IsRegExp: false,
		},
	}
	macs, err := sw.GetMacTable(filters)
	if err != nil {
		log.Fatalf("%s", err)
	}
	if len(macs) == 0 {
		log.Fatalf("no MAC's are found.")
	}
	for _, v := range macs {
		fmt.Printf("vlan: %s MAC: %s PortIndex: %v\n", v.VLANID, v.MacAddress, v.PortIndex)

	}
}
