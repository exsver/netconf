package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		r, err := sw.GetACLConfig()
		spew.Dump(r, err)
	}

}
