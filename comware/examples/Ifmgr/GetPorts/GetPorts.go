package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()
	netconf.LogLevel.Verbose()

	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Getting data
	ports, err := sw.GetPortsRegExp("1/0/")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print
	for _, port := range ports {
		fmt.Printf("ifIndex:%v, portIndex:%v, name:%s\n", port.IfIndex, port.PortIndex, port.Name)
	}
}
