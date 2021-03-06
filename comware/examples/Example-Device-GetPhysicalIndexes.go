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

	boards, err := sw.GetIndexBoards()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Boards: %v\n", boards)

	sensors, err := sw.GetPhysicalIndexSensors()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Sensors: %v\n", sensors)

	psu, err := sw.GetPhysicalIndexPSU()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("PSU: %v\n", psu)

	fans, err := sw.GetPhysicalIndexFan()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("FANs: %v\n", fans)
}
