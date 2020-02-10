package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	data, err := sw.GetDataResourceMonitor()
	if err != nil {
		log.Fatalf("%s", err)
	}
	for _, monitor := range data.Monitors.Monitors {
		fmt.Printf("Slot: %v, Name: %s, Used: %v, Total: %v\n", monitor.DeviceNode.Slot, monitor.Name, monitor.Used, monitor.Total)
	}
}
