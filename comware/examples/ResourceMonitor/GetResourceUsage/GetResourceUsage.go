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
	data, err := sw.GetDataResourceMonitor()
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, monitor := range data.Monitors.Monitors {
		fmt.Printf("Slot: %v, Name: %s, Used: %v, Total: %v\n", monitor.DeviceNode.Slot, monitor.Name, monitor.Used, monitor.Total)
	}
}
