package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = sw.MirroringGroupRemove(1)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("Removed or not exist")
}
