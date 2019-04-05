package main

import (
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/helpers/comware"
	"log"
)

func main() {
	netconf.LogLevel.Verbose()
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf", "netconf")
	if err != nil {
		log.Fatalf("%s", err)
	}
	acl := comware.NamedGroup{
		GroupType:     1,
		GroupCategory: 2,
		GroupIndex:    "testACL",
	}
	err = sw.ACLRemove(&acl)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
