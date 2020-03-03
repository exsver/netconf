package main

import (
	"log"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

// Run ./Example-FirewallFilterTerm-Add.go first

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

func main() {
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// delete firewall family inet filter testFilterFamilyInet term drop10
	term := junos.Term{
		NetconfConfigOperation: "remove",
		Name:                   "drop10",
	}

	err = device.EditConfig(term.ConvertToConfig("inet", "testFilterFamilyInet"), "none")
	if err != nil {
		log.Fatalf("%s", err)
	}
}
