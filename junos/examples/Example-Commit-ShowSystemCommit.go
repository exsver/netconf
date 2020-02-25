package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/junos"
)

// CLI equivalent: run show system commit
func main() {
	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	commits, err := device.GetSystemCommit()
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, commit := range commits {
		fmt.Printf("%2v %s by %s via %s\n", commit.SequenceNumber, commit.DateTime, commit.User, commit.Client)
	}
}
