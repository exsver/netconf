package main

import (
	"log"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"

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

	// all rules (all ACLs)
	data, err := sw.ACLIPv4NamedAdvanceRulesGet(nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	spew.Dump(data)

	// all rules for acl named testACL
	data, err = sw.ACLIPv4NamedAdvanceRulesGet([]comware.XMLFilter{{Key: "GroupIndex", Value: "testACL", IsRegExp: false}})
	if err != nil {
		log.Fatalf("%s", err)
	}

	spew.Dump(data)

	//  all rules with Action "Deny" (all ACLs)
	data, err = sw.ACLIPv4NamedAdvanceRulesGet([]comware.XMLFilter{{Key: "Action", Value: strconv.Itoa(comware.ACLRuleActionDeny), IsRegExp: false}})
	if err != nil {
		log.Fatalf("%s", err)
	}

	spew.Dump(data)

	//  all rules with ProtocolType "ICMP" and Action "Deny" for acl named testACL
	data, err = sw.ACLIPv4NamedAdvanceRulesGet(
		[]comware.XMLFilter{
			{Key: "GroupIndex", Value: "testACL", IsRegExp: false},
			{Key: "ProtocolType", Value: strconv.Itoa(comware.ProtocolICMP), IsRegExp: false},
			{Key: "Action", Value: strconv.Itoa(comware.ACLRuleActionDeny), IsRegExp: false},
		})
	if err != nil {
		log.Fatalf("%s", err)
	}

	spew.Dump(data)
}
