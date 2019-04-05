package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	output, err := sw.RunCLICommand(`
#
interface GigabitEthernet1/0/3
 default
 description to_switch_205
 port link-type trunk
 undo port trunk permit vlan 1
 port trunk permit vlan 300 to 302
#
`,
		true)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", output)
		spew.Dump(output)
	}

}
