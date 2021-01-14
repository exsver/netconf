package main

import (
	"fmt"
	"github.com/exsver/netconf/comware"

	"log"
)

func main() {
	// Creating a new device.
	sw, err := comware.NewTargetDevice("172.21.1.249", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// CLI equivalent: "port-security enable"
	config := comware.PortSecurityCommon{
		Enable: true,
	}

	// CLI equivalent: "undo port-security enable"
	//  config := comware.PortSecurityCommon{
	// 	  Enable: false,
	//  }

	err = sw.PortSecurityEnable(&config)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("OK")
}
