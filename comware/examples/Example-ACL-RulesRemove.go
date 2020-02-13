package main

import (
	"log"
	"time"

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

	// Use one netconf session for all operations
	err = sw.Connect(time.Minute * 5)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer sw.Disconnect()

	rules := comware.IPv4NamedAdvanceRules{
		IPv4NamedAdvanceRules: []comware.IPv4NamedAdvanceRule{
			{
				GroupIndex:   "testACL",
				RuleID:       0,
			},
			{
				GroupIndex:   "testACL",
				RuleID:       5,
			},
			{
				GroupIndex:   "testACL",
				RuleID:       10,
			},
			{
				GroupIndex:   "testACL",
				RuleID:       15,
			},
		},
	}

	err = sw.ACLIPv4NamedAdvanceRulesRemove(&rules)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

