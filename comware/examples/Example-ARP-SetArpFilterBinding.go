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

	// interface GigabitEthernet1/0/10
	//   ip source binding ip-address 10.0.0.1 mac-address AABB-CCDD-EEFF
	err = sw.SetArpFilterBinding("10", "10.0.0.1", "AA-BB-CC-DD-EE-FF")
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("End")
}
