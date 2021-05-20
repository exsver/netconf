package main

import (
	"fmt"
	"log"
	"time"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

/* switch configuration
system-view
vlan 300
quit
interface GigabitEthernet 1/0/10
 port link-type trunk
 port trunk permit vlan 300
 quit
*/

func checkIsVlanExist(sw *comware.TargetDevice, vlan int) error {
	ok, err := sw.IsVlanExist(vlan)
	if err != nil {
		return err
	}
	if ok {
		fmt.Printf("vlan %s exist", vlan)
	} else {
		fmt.Printf("vlan %s does not exist", vlan)
	}
	return nil
}

func checkIsVlanInTrunk(sw *comware.TargetDevice, ifName string, vlan int) error {
	ok, err := sw.IsVlanInTrunk(ifName, vlan)
	if err != nil {
		return err
	}
	if ok {
		fmt.Printf("vlan %v is allowed in the trunk", vlan)
	} else {
		fmt.Printf("vlan %v is not allowed in the trunk", vlan)
	}
	return nil
}

func main() {
	netconf.LogLevel.Messages()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = sw.Connect(time.Minute * 10)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer sw.Disconnect()

	err = checkIsVlanExist(sw, 300)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = checkIsVlanInTrunk(sw, "GigabitEthernet1/0/10", 300)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
