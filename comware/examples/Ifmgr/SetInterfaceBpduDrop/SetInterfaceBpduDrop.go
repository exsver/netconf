package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if err := sw.SetInterfaceBpduDrop(10, true); err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("End")
}
