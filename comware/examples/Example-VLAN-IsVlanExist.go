package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	ok, err := sw.IsVlanExist(99)
	if err != nil {
		log.Fatalf("%s", err)
	}
	if ok {
		fmt.Println("Vlan exist")
	} else {
		fmt.Println("Vlan not exist")
	}
}
