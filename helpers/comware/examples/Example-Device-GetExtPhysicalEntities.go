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
	entities, err := sw.GetIndexBoards()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Boards: %s\n", entities)
}
