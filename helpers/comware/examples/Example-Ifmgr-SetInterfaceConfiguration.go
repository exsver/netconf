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
		iface := comware.Interface{

			IfIndex:     10,
			Description: "h1000000",
			ConfigSpeed: 2,
			AdminStatus: 2,
		}

		err := sw.Configure(*iface.ConvertToTop(), "merge")
		spew.Dump(err)
		sw.SetInterfaceLinkType(10, "trunk")
	}
}
