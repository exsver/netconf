package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	interfaceConfig := comware.Interface{
		IfIndex:      10,                             // interface GigabitEthernet 1/0/10
		Description:  "test-descr",                   // description test-descr
		ConfigSpeed:  comware.InterfaceSpeed1G,       // speed 1000
		ConfigDuplex: comware.InterfaceDuplexFull,    // duplex full
		AdminStatus:  comware.InterfaceAdminStatusUP, // no shutdown
	}

	err = sw.Configure(*interfaceConfig.ConvertToTop(), "merge")
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("OK\n")
}
