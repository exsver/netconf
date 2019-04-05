package main

import (
	"fmt"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	// netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	slots, err := sw.GetSlotsInfo()
	if err != nil {
		log.Fatalf("%s", err)
	}
	for _, slot := range slots {
		fmt.Printf("Slot: %v, Model: %s, Description: %s, SoftwareRev: %s, SerialNumber: %s\n", slot.Slot, slot.Model, slot.Description, slot.SoftwareRev, slot.SerialNumber)
	}

}
