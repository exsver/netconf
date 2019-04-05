package main

import (
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

//CLI equivalent
// ARPRateLimitLogEnable(true) - "arp rate-limit log enable"
// ARPRateLimitLogEnable(false) - "undo arp rate-limit log enable"

func main() {
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	err = sw.ARPRateLimitLogEnable(true)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
