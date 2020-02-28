package main

import (
	"fmt"
	"log"
	"time"

	"github.com/exsver/netconf/junos"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Messages()

	device, err := junos.NewTargetDevice("172.21.1.250", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatal(err)
	}

	err = device.Connect(time.Minute * 5)
	if err != nil {
		log.Fatal(err)
	}
	defer device.Disconnect()

	openConfigurationMessage, err := device.OpenConfigurationPrivate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(openConfigurationMessage)

	// more LoadConfiguration examples: ./Example-LoadConfiguration.go
	loadConfigurationResults, err := device.LoadConfiguration("text", "set system host-name test-mx80","set")
	if err != nil {
		log.Fatal(err)
	}

	if loadConfigurationResults.GetErrors() != nil {
		log.Fatal(loadConfigurationResults.GetErrors())
	}

	commitResult, err := device.Commit()
	if err != nil {
		log.Fatal(err)
	}

	if !commitResult.OK {
		log.Fatal("commit failed")
	}

	fmt.Println("Ok")
}
