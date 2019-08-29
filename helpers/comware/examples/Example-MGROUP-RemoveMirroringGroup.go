package main

import (
	"fmt"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	err = sw.RemoveMirroringGroup(1)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("Removed or not exist")
}
