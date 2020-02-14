package main

import (
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

// Run ./Example-ACL-ACLCreate.go first

func main() {
	netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	acl := comware.NamedGroup{
		GroupType:     comware.ACLGroupTypeIPv4,
		GroupCategory: comware.ACLGroupCategoryAdvanced,
		GroupIndex:    "testACL",
	}

	err = sw.ACLRemove(&acl)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
