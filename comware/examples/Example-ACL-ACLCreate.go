package main

import (
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
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
		Description:   "ACL example",
	}
	err = sw.ACLCreate(&acl)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
