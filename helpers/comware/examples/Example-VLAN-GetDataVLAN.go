package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		data, err := sw.GetDataVLAN()
		spew.Dump(data, err)
	}
}
