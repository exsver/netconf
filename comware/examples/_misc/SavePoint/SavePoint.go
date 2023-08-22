package main

import (
	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
	"log"
	"time"
)

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()
	netconf.LogLevel.Messages()

	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = sw.Connect(time.Minute * 10)
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer sw.Disconnect()

	commitID, err := sw.SavePointBegin(60)
	if err != nil {
		log.Fatalf("%s", err)
	}

	vlans1 := &comware.VLANs{
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

	err = sw.VlanCreate(vlans1)
	if err != nil {
		log.Println(sw.SavePointRollback(commitID))
	}

	vlans2 := &comware.VLANs{
		VLANs: []comware.VLANID{
			{
				ID:   301,
				Name: "v301-test",
			},
			{
				ID:   302,
				Name: "v302-test",
			},
		},
	}

	err = sw.VlanCreate(vlans2)
	if err != nil {
		log.Println(sw.SavePointRollback(commitID))
	}

	log.Println(sw.SavePointEnd())
}
