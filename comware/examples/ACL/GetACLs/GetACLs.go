package main

import (
	"fmt"
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

func main() {
	// Setting the Log Level for netconf lib.
	// One of:
	//   netconf.LogLevel.Silent()
	//   netconf.LogLevel.Default() - default
	//   netconf.LogLevel.Messages()
	//   netconf.LogLevel.Verbose()
	netconf.LogLevel.Verbose()

	// Creating a new device.
	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Filters examples:
	// all items -                        nil
	// ACL with name testACL -            []string{`<GroupIndex>testACL</GroupIndex>`}
	// ACLs with Description "aclDescr" - []string{`<Description>aclDescr</Description>`}
	// All IPv4 ACLs -                    []string{`<GroupType>1</GroupType>`}
	// All IPv6 ACLs -                    []string{`<GroupType>2</GroupType>`}
	// All advanced ACLs -                []string{`<GroupCategory>2</GroupCategory>`}
	// All advanced IPv4 ACLs -           []string{`<GroupCategory>2</GroupCategory>`, `<GroupType>1</GroupType>`}
	// All ACLs with rule number 1 -      []string{`<RuleNum>1</RuleNum>`}

	acls, err := sw.ACLGetNamedGroups([]string{`<GroupType>1</GroupType>`})
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print
	for _, v := range acls {
		fmt.Printf("ACL Type %v, ACL Number/Name: %v, ACL Decription: %s, RuleNumber: %v\n", v.GroupType, v.GroupIndex, v.Description, v.RuleNum)
	}
}
