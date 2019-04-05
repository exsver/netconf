package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		data, err := sw.GetDataARP()
		spew.Dump(data, err)
	}
}
