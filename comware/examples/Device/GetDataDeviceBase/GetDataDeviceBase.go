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

	data, err := sw.GetDataDeviceBase()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("Hostname: '%s', Uptime: %v seconds, LocalTime: %s, BridgeMAC: %s, HostDescription: '%s'",
		data.HostName,
		data.Uptime,
		data.LocalTime,
		data.BridgeMAC,
		data.HostDescription,
	)
}
