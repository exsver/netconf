package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	ports, err := sw.GetPortsRegExp("1/0/")
	if err != nil {
		log.Fatalf("%s", err)
	}
	for _, port := range ports {
		fmt.Printf("ifIndex:%v, portIndex:%v, name:%s\n", port.IfIndex, port.PortIndex, port.Name)
	}

}
