package main

import (
	"fmt"
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

	//Filters examples:
	// all items -                        nil
	// ACL with name testACL -            []string{`<GroupIndex>testACL</GroupIndex>`}
	// ACLs with Description "aclDescr" - []string{`<Description>aclDescr</Description>`}
	// All IPv4 ACLs -                    []string{`<GroupType>1</GroupType>`}
	// All IPv6 ACLs -                    []string{`<GroupType>2</GroupType>`}
	// All advanced ACLs -                []string{`<GroupCategory>2</GroupCategory>`}
	// All advanced IPv4 ACLs -           []string{`<GroupCategory>2</GroupCategory>`, `<GroupType>1</GroupType>`}
	// All ACLs with rule number 1 -      []string{`<RuleNum>1</RuleNum>`}

	acls, err := sw.GetListOfNamedACL([]string{`<GroupType>2</GroupType>`})
	if err != nil {
		log.Fatalf("%s", err)
	}
	for _, v := range acls {
		fmt.Printf("ACL Type %v, ACL Number/Name: %v, ACL Decription: %s, RuleNumber: %v\n", v.GroupType, v.GroupIndex, v.Description, v.RuleNum)
	}
}
