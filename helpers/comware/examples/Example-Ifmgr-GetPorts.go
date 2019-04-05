package main

import (
	"fmt"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
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
