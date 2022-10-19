package main

import (
	"fmt"
	"github.com/exsver/netconf/comware"
	"log"
)

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Preparing a set of commands
	commands := `display current-configuration`

	// Executing commands in device
	output, err := sw.RunCLICommand(commands, false)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Printing output
	fmt.Printf("%s", comware.CorrectNewLines(output))
}
