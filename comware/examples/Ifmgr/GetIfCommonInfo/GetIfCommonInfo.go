package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
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

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatal(err)
	}

	data, err := sw.GetIfCommonInfo(10)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(data)
}
