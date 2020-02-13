package main

import (
	"log"

	"github.com/exsver/netconf/comware"
	"github.com/exsver/netconf/netconf"
)

/*
#
acl advanced name testACL
 description ACL example
 rule 0 deny icmp destination 10.100.100.1 0
 rule 0 comment Deny ICMP
 rule 5 deny udp destination 10.100.100.1 0 counting
 rule 5 comment Deny UDP
 rule 10 permit tcp destination 10.100.100.1 0 destination-port eq www
 rule 10 comment Permit WWW
 rule 15 deny ip destination 10.100.100.1 0
 rule 15 comment Deny All IP
#

*/

func main() {
	netconf.LogLevel.Verbose()

	sw, err := comware.NewTargetDevice("10.10.10.10", "netconf-user", "netconf-password")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// CLI equivalent:
	// acl advanced name testACL
	//  description ACL example
	acl := comware.NamedGroup{
		GroupType:     comware.ACLGroupTypeIPv4,
		GroupCategory: comware.ACLGroupCategoryAdvanced,
		GroupIndex:    "testACL",     // ACL name
		Description:   "ACL example",
	}

	err = sw.ACLCreate(&acl)
	if err != nil {
		log.Fatalf("%s", err)
	}

	rules := comware.IPv4NamedAdvanceRules{
		IPv4NamedAdvanceRules: []comware.IPv4NamedAdvanceRule{
			// rule 0 deny icmp destination 10.100.100.1 0
			// rule 0 comment Deny ICMP
			{
				GroupIndex:   acl.GroupIndex,
				RuleID:       comware.ACLRuleIDAuto,
				Action:       comware.ACLRuleActionDeny,
				ProtocolType: comware.ProtocolICMP,
				DstAny:       comware.NewFalse(),
				DstIPv4: &comware.DstIPv4{
					DstIPv4Addr:     "10.100.100.1",
					DstIPv4Wildcard: "0.0.0.0",
				},
				Comment: "Deny ICMP", // a case-sensitive string of 1 to 127 characters
			},
			// rule 5 deny udp destination 10.100.100.1 0 counting
			// rule 5 comment Deny UDP
			{
				GroupIndex:   acl.GroupIndex,
				RuleID:       comware.ACLRuleIDAuto,
				Action:       comware.ACLRuleActionDeny,
				ProtocolType: comware.ProtocolUDP,
				DstAny:       comware.NewFalse(),
				DstIPv4: &comware.DstIPv4{
					DstIPv4Addr:     "10.100.100.1",
					DstIPv4Wildcard: "0.0.0.0",
				},
				Counting: true,
				Comment:  "Deny UDP", // a case-sensitive string of 1 to 127 characters
			},
			// rule 10 permit tcp destination 10.100.100.1 0 destination-port eq 80
			// rule 10 comment Permit WWW
			{
				GroupIndex:   acl.GroupIndex,
				RuleID:       comware.ACLRuleIDAuto,
				Action:       comware.ACLRuleActionPermit,
				ProtocolType: comware.ProtocolTCP,
				DstAny:       comware.NewFalse(),
				DstIPv4: &comware.DstIPv4{
					DstIPv4Addr:     "10.100.100.1",
					DstIPv4Wildcard: "0.0.0.0",
				},
				DstPort: &comware.DstPort{
					DstPortOp:     comware.OperationEqual,
					DstPortValue1: 80,
					DstPortValue2: 65536,
				},
				Comment: "Permit WWW", // a case-sensitive string of 1 to 127 characters
			},
			// rule 15 deny ip destination 10.100.100.1 0
			// rule 15 comment Deny All IP
			{
				GroupIndex:   acl.GroupIndex,
				RuleID:       comware.ACLRuleIDAuto,
				Action:       comware.ACLRuleActionDeny,
				ProtocolType: comware.ProtocolAny,
				DstAny:       comware.NewFalse(),
				DstIPv4: &comware.DstIPv4{
					DstIPv4Addr:     "10.100.100.1",
					DstIPv4Wildcard: "0.0.0.0",
				},
				Comment: "Deny All IP", // a case-sensitive string of 1 to 127 characters
			},
		},
	}

	err = sw.ACLIPv4NamedAdvanceRulesAdd(&rules)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
