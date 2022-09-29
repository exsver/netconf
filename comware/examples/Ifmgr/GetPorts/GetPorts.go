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
	netconf.LogLevel.Default()

	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Getting all ports
	ports, err := sw.GetPorts()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print
	fmt.Println("All ports:")
	for _, port := range ports {
		fmt.Printf("ifIndex:%v, portIndex:%v, name:%s\n", port.IfIndex, port.PortIndex, port.Name)
	}

	// Getting ports by RegExp
	// RegExp Examples:
	//   "^GigabitEthernet"       -- all GigabitEthernet ports
	//   "^Ten-GigabitEthernet"   -- all Ten-GigabitEthernet ports
	//   "1/0/"                   -- slot 1 ports
	ports, err = sw.GetPortsRegExp("^Ten-GigabitEthernet")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print
	fmt.Println("Ten-GigabitEthernet ports:")
	for _, port := range ports {
		fmt.Printf("ifIndex:%v, portIndex:%v, name:%s\n", port.IfIndex, port.PortIndex, port.Name)
	}
}
