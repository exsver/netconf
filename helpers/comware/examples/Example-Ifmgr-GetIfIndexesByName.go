package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/helpers/comware"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}
	ifIndexes, err := sw.GetIfIndexesByName("Bridge", true)
	if err != nil {
		log.Fatalf("%s", err)
	}
	if len(ifIndexes) == 0 {
		fmt.Println("No interfaces found")
	} else {
		fmt.Printf("%v", ifIndexes)
	}
}
