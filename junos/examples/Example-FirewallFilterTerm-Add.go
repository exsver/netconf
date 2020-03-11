package main

import (
	"log"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

/*
set firewall family inet filter testFilterFamilyInet term accept then accept

test-mx80# show firewall family inet filter testFilterFamilyInet
term accept {
    then accept;
}
*/

func main() {
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// set firewall family inet filter testFilterFamilyInet term drop10 from destination-address 10.0.0.0/8
	// set firewall family inet filter testFilterFamilyInet term drop10 then discard
	term1 := junos.Term{
		Name: "drop10",
		From: &junos.FilterFrom{
			DestinationAddress: []junos.FilterDestinationAddress{
				{Name: "10.0.0.0/8"},
			},
		},
		Then: &junos.FilterThen{
			Discard: &junos.FilterThenDiscard{},
		},
	}

	// family: "" | "inet" | "inet6"
	// operations:
	//  "merge" -  The device merges new configuration data into the existing configuration data. This is the default.
	//  "replace" - The device replaces existing configuration data with the new configuration data.
	err = device.EditConfig(term1.ConvertToConfig(true, "testFilterFamilyInet"), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// edit firewall family inet filter testFilterFamilyInet
	// insert term drop10 before term accept
	term1 = junos.Term{
		Name:                  "drop10",
		NetconfInsertPosition: "first",
	}

	err = device.EditConfig(term1.ConvertToConfig(true, "testFilterFamilyInet"), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// set firewall family inet filter testFilterFamilyInet term drop172 from destination-address 172.16.0.0/12
	// set firewall family inet filter testFilterFamilyInet term drop172 then discard
	term2 := junos.Term{
		Name: "drop172",
		From: &junos.FilterFrom{
			DestinationAddress: []junos.FilterDestinationAddress{
				{Name: "172.16.0.0/12"},
			},
		},
		Then: &junos.FilterThen{
			Discard: &junos.FilterThenDiscard{},
		},
	}

	// edit firewall family inet filter testFilterFamilyInet
	// insert term drop172 after term drop10
	err = device.EditConfig(term2.ConvertToConfig(true, "testFilterFamilyInet"), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	term2 = junos.Term{
		Name:                      "drop172",
		NetconfInsertPosition:     "after",
		NetconfInsertPositionName: "drop10",
	}

	err = device.EditConfig(term2.ConvertToConfig(true, "testFilterFamilyInet"), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	/*
		test-mx80# show firewall family inet filter testFilterFamilyInet
		term drop10 {
		    from {
		        destination-address {
		            10.0.0.0/8;
		        }
		    }
		    then {
		        discard;
		    }
		}
		term drop172 {
		    from {
		        destination-address {
		            172.16.0.0/12;
		        }
		    }
		    then {
		        discard;
		    }
		}
		term accept {
		    then accept;
		}
	*/
}
