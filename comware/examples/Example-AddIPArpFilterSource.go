package main

import (
	"log"

	"github.com/exsver/netconf/comware"
)

func main() {
	// netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	if err := sw.AddIPArpFilterSource(10, []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}); err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("End")

}
