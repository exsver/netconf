package main

import (
	"fmt"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	ifIndexes, err := sw.GetIfIndexesByDecription("Uplink", false)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if len(ifIndexes) == 0 {
		fmt.Println("No interfaces found with")
	} else {
		fmt.Printf("%v", ifIndexes)
	}
}
