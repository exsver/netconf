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

	vlans := &comware.VLANs{
		VLANs: []comware.VLANID{
			{
				ID:   300,
				Name: "v300-test",
			},
			{
				ID:   301,
				Name: "v301-test",
			},
		},
	}

	err = sw.VlanCreate(vlans)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("Ok")
}
