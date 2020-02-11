package main

import (
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	acl := comware.NamedGroup{
		GroupType:     comware.ACLGroupTypeIPv4,
		GroupCategory: comware.ACLGroupCategoryAdvanced,
		GroupIndex:    "testACL", // ACL name
		Description:   "ACL example",
	}

	err = sw.ACLCreate(&acl)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
