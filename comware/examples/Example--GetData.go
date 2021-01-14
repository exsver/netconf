package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	// Setting the Log Level for netconf lib.
	// one of:
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

	// You can use any GetData* function.
	// For example:
	//  GetDataACL()
	//  GetDataARP()
	//  GetDataDevice()
	//  GetDataDHCPServer()
	//  GetDataDHCPSP()
	//  GetDataIfmgr()
	//  GetDataIRF()
	//  GetDataMAC()
	//  GetDataMGROUP()
	//  GetDataResourceMonitor()
	//  GetDataVLAN()
	//
	//  and others (one for each subtree)
	data, err := sw.GetDataDevice()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Printing the data structures
	spew.Dump(data)
}
