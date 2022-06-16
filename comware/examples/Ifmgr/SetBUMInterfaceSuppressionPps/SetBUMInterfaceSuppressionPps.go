package main

import (
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

	// interface GigabitEthernet1/0/10
	//    broadcast-suppression pps 100
	//    multicast-suppression pps 100
	//    unicast-suppression pps 2000
	if err := sw.SetInterfaceSuppressionPps(10, 100, 100, 2000); err != nil {
		log.Fatalf("%s", err)
	}

	// fmt.Println("End")

}
