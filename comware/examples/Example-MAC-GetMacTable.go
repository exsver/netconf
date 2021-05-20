package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	// netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Available filters are:
	//  - VLANID
	//  - MacAddress
	//  - PortIndex
	//  - Status
	//  - Aging
	// Filter examples:
	//  - all mac-addresses									   nil
	//  - all mac-addresses with VLANID 99  --                 []comware.XMLFilter{{Key: "VLANID", Value:"99", IsRegExp:false,},}
	//  - all mac-addresses with VLANID 99 and PortIndex 1 --  []comware.XMLFilter{{Key: "VLANID", Value:"99", IsRegExp:false,},{Key: "PortIndex", Value:"1", IsRegExp:false,},}
	//  - all mac-addresses starts with "40-B0-34" --          []comware.XMLFilter{{Key: "MacAddress", Value:"^40-B0-34", IsRegExp:true,},}

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
