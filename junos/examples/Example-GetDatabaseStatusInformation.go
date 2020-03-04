package main

import (
	"bytes"
	"fmt"
	"github.com/exsver/netconf/junos"
	"log"
)

func main() {
	device, err := junos.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	info, err :=  device.GetDatabaseStatusInformation()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("Users currently editing the configuration:")
	for _, user := range info.DatabaseStatus {
		fmt.Printf("%s (pid %s) since %s idle %s\n", user.User, user.PID, user.StartTime, user.IdleTime)
	}

	fmt.Println("Configuration Diff (show | compare):")
	output, err := device.CompareConfigurationRollback(0, "text")
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("%s", bytes.TrimSpace(output))
}
