package comware

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ACL struct {
	/* top level
	   ACL
	     Base
	     Capability                     ***ReadOnly***
	     Groups
	       []Group
	     IPv4AdvanceRules
	       []IPv4AdvanceRule
	     IPv4BasicRules
	       []IPv4BasicRules
	     IPv4NamedAdvanceRules
	       []IPv4NamedAdvanceRule
	     IPv4NamedBasicRules
	       []IPv4NamedBasicRule
	     IPv6AdvanceRules
	       []IPv6AdvanceRule
	     IPv6BasicRules
	       []IPv6BasicRule
	     IPv6NamedAdvanceRules
	       []IPv6NamedAdvanceRule
	     IPv6NamedBasicRules
	       []IPv6NamedBasicRule
	     MACNamedRules
	       []MACNamedRule
	     MACRules
	       []MACRule
	     NamedGroups
	     PfilterApply
	     PfilterDefAction
	     PfilterGroupRunInfo
	     PfilterIgnoreAction
	     PfilterRuleRunInfo
	     PfilterStatisticSum
	*/
	Groups                *Groups                `xml:"Groups"`
	NamedGroups           *NamedGroups           `xml:"NamedGroups"`
	IPv4AdvanceRules      *IPv4AdvanceRules      `xml:"IPv4AdvanceRules"`
	IPv4NamedAdvanceRules *IPv4NamedAdvanceRules `xml:"IPv4NamedAdvanceRules"`
	IPv4BasicRules        *IPv4BasicRules        `xml:"IPv4BasicRules"`
	IPv4NamedBasicRules   *IPv4NamedBasicRules   `xml:"IPv4NamedBasicRules"`
	IPv6AdvanceRules      *IPv6AdvanceRules      `xml:"IPv6AdvanceRules"`
	IPv6NamedAdvanceRules *IPv6NamedAdvanceRules `xml:"IPv6NamedAdvanceRules"`
	IPv6BasicRules        *IPv6BasicRules        `xml:"IPv6BasicRules"`
	IPv6NamedBasicRules   *IPv6NamedBasicRules   `xml:"IPv6NamedBasicRules"`
	MACRules              *MACRules              `xml:"MACRules"`
	PfilterDefAction      *PfilterDefAction      `xml:"PfilterDefAction"`
	PfilterApply          *PfilterApply          `xml:"PfilterApply"`
	PfilterGroupRunInfo   *PfilterGroupRunInfo   `xml:"PfilterGroupRunInfo"`
}

// Groups table contains ACL information.
type Groups struct {
	XMLName xml.Name `xml:"Groups"`
	Groups  []Group  `xml:"Group"`
}

// NamedGroups table contains named ACL information.
type NamedGroups struct {
	XMLName xml.Name     `xml:"NamedGroups"`
	Groups  []NamedGroup `xml:"Group"`
}

type Group struct {
	XMLName xml.Name `xml:"Group"`
	// GroupType specifies the type of ACL:
	// 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined.
	GroupType ACLGroupType `xml:"GroupType"`
	// GroupID specifies the ACL number.
	// The value range depends on the GroupType column.
	// - 2000 to 5999 if GroupType is 1 (IPv4).
	//    IPv4 basic ACL: 2000 to 2999.
	//    IPv4 advanced ACL: 3000 to 3999.
	//    Ethernet frame header ACL: 4000 to 4999.
	//    User-defined ACL: 5000 to	5999.
	// - 2000 to 3999 if GroupType is 2 (IPv6).
	//    IPv6 basic ACL: 2000 to 2999.
	//    IPv6 advanced ACL: 3000 to 3999.
	GroupID int `xml:"GroupID"`
	// Match order: 1 - config (Ascending order of rule IDs), 2 - auto (Depth-first order). Default: config.
	// The match order can only be	modified for ACLs that do not contain any rules.
	// The match order can only be config for user-defined ACLs.
	MatchOrder int `xml:"MatchOrder,omitempty"`
	// Step length, range from 1 to 20. Default: 5. The rule numbering step for user-defined ACLs can only be 5.
	Step int `xml:"Step,omitempty"`
	// Description of the ACL group, a case-sensitive string of 1 to 127 characters.
	Description string `xml:"Description,omitempty"`
	// The rules number of the ACL group. Range from 0 to 65535.
	// ReadOnly
	RuleNum int `xml:"RuleNum,omitempty"`
}

type NamedGroup struct {
	XMLName xml.Name `xml:"Group"`
	// GroupType specifies the type of ACL:
	// 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined.
	GroupType ACLGroupType `xml:"GroupType"`
	// GroupCategory specifies the category of ACL: 0 - invalid, 1 - basic, 2 - advanced.
	// The value range depends on the GroupType column.
	// - 1 to 2 if GroupType is 1 or 2.
	//    basic ACL: 1.
	//    advanced ACL: 2.
	// - 0 if GroupType is 3 or 4.
	GroupCategory ACLGroupCategory `xml:"GroupCategory"`
	// GroupIndex specifies ACL name STRING<1-63> (case-insensitive string) or ACL index.
	// Can't use 'all' as ACL name. If it's Index, range from 2000 to 5999.
	GroupIndex string `xml:"GroupIndex"`
	// Match order: 1 - config (Ascending order of rule IDs), 2 - auto (Depth-first order). Default: config.
	// The match order can only be	modified for ACLs that do not contain any rules.
	// The match order can only be config for user-defined ACLs.
	MatchOrder int `xml:"MatchOrder,omitempty"`
	// Step length, range from 1 to 20. Default: 5. The rule numbering step for user-defined ACLs can only be 5.
	Step int `xml:"Step,omitempty"`
	// Description of the ACL group, a case-sensitive string of 1 to 127 characters.
	Description string `xml:"Description,omitempty"`
	// The rules number of the ACL group. Range from 0 to 65535.
	// ReadOnly
	RuleNum int `xml:"RuleNum,omitempty"`
}

// IPv4AdvanceRules table contains IPv4 advanced ACL rule information.
type IPv4AdvanceRules struct {
	IPv4AdvanceRules []IPv4AdvanceRule `xml:"Rule"`
}

// IPv4NamedAdvanceRules table contains information about named IPv4 advanced ACL rules.
type IPv4NamedAdvanceRules struct {
	IPv4NamedAdvanceRules []IPv4NamedAdvanceRule `xml:"Rule"`
}

// IPv4BasicRules table contains information about IPv4 basic ACL rules.
type IPv4BasicRules struct {
	IPv4BasicRules []IPv4BasicRule `xml:"Rule"`
}

// IPv4NamedBasicRules table contains information about named IPv4 basic ACL rules.
type IPv4NamedBasicRules struct {
	IPv4NamedBasicRules []IPv4NamedBasicRule `xml:"Rule"`
}

// IPv6AdvanceRules table contains IPv6 advanced ACL rule information.
type IPv6AdvanceRules struct {
	IPv6AdvanceRules []IPv6AdvanceRule `xml:"Rule"`
}

// IPv6NamedAdvanceRules table contains IPv6 advanced ACL rule information.
type IPv6NamedAdvanceRules struct {
	IPv6NamedAdvanceRules []IPv6NamedAdvanceRule `xml:"Rule"`
}

// IPv6BasicRules table contains information about IPv6 basic ACL rules.
type IPv6BasicRules struct {
	IPv6BasicRules []IPv6BasicRule `xml:"Rule"`
}

// IPv6NamedBasicRules table contains information about IPv6 basic ACL rules.
type IPv6NamedBasicRules struct {
	IPv6NamedBasicRules []IPv6NamedBasicRule `xml:"Rule"`
}

// MACRules table contains Ethernet frame header ACL rule information.
type MACRules struct {
	MACRule []MACRule `xml:"Rule"`
}

type IPv4AdvanceRule struct {
	XMLName xml.Name `xml:"Rule"`
	GroupID int      `xml:"GroupID"`
	// RuleID int in range 0-65534
	RuleID int `xml:"RuleID"`
	// Action: 1 - Deny, 2 - Permit
	Action ACLRuleAction `xml:"Action"`
	// ProtocolType defines:
	// Protocol number INTEGER<0-255>, 256 - any IP protocol
	// 1 - ICMP
	// 6 - TCP
	// 17 - UDP
	// ...
	// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	ProtocolType int           `xml:"ProtocolType,omitempty"`
	Count        int           `xml:"Count,omitempty"`
	Status       ACLRuleStatus `xml:"Status,omitempty"`
	Fragment     bool          `xml:"Fragment,omitempty"`
	Logging      bool          `xml:"Logging,omitempty"`
	Counting     bool          `xml:"Counting,omitempty"`
	SrcAny       *bool         `xml:"SrcAny,omitempty"`
	DstAny       *bool         `xml:"DstAny,omitempty"`
	SrcIPv4      *SrcIPv4      `xml:"SrcIPv4,omitempty"`
	DstIPv4      *DstIPv4      `xml:"DstIPv4,omitempty"`
	SrcPort      *SrcPort      `xml:"SrcPort,omitempty"`
	DstPort      *DstPort      `xml:"DstPort,omitempty"`
}

type IPv4NamedAdvanceRule struct {
	XMLName xml.Name `xml:"Rule"`
	// GroupIndex - ACL name, or ACL number.
	// You must create an ACL first before you create, merge, or replace rules for it.
	GroupIndex string `xml:"GroupIndex"`
	// Value range: 0 to 65534.
	// If you set this column to 65535, the system automatically assigns a new rule ID.
	// This rule ID is the nearest higher multiple of the numbering step to the current highest rule ID, starting from 0.
	RuleID int `xml:"RuleID"`
	// Action on packets matching the rule.
	// Action: 1 - Deny, 2 - Permit
	Action ACLRuleAction `xml:"Action,omitempty"`
	// Protocol type.
	// Value range: 0 to 256. The value 256 represents all IPv4 protocols.
	ProtocolType int   `xml:"ProtocolType,omitempty"`
	Count        int64 `xml:"Count,omitempty"`
	// Rule status.
	// Status: 1 - active, 2 - inactive.
	Status ACLRuleStatus `xml:"Status,omitempty"`
	// Fragment - the flag of matching fragmented packet.
	// If an ACL is for QoS traffic classification or packet filtering do not specify the fragment.
	Fragment *bool    `xml:"Fragment,omitempty"`
	Logging  *bool    `xml:"Logging,omitempty"`  // The logging takes effect only when the module (for example, packet filtering) that uses the ACL supports logging.
	Counting *bool    `xml:"Counting,omitempty"` // Counts times the ACL rule has been matched.
	SrcAny   *bool    `xml:"SrcAny,omitempty"`
	DstAny   *bool    `xml:"DstAny,omitempty"`
	SrcIPv4  *SrcIPv4 `xml:"SrcIPv4,omitempty"`
	DstIPv4  *DstIPv4 `xml:"DstIPv4,omitempty"`
	SrcPort  *SrcPort `xml:"SrcPort,omitempty"`
	DstPort  *DstPort `xml:"DstPort,omitempty"`
	// Rule comment,
	// a case-sensitive string of 1 to 127 characters.
	Comment string `xml:"Comment,omitempty"`
}

type IPv4BasicRule struct {
	XMLName xml.Name `xml:"Rule"`
	// Group ID. Range from 2000 to 2999.
	GroupID int `xml:"GroupID"`
	// Rule ID. Range from 0 to 65534
	// If you set this column to 65535, the system automatically assigns a new rule ID.
	RuleID int `xml:"RuleID"`
	// Action on packets matching the rule.
	// Action: 1 - Deny, 2 - Permit
	Action ACLRuleAction `xml:"Action,omitempty"`
	// Fragment - the flag of matching fragmented packet.
	//  false - the rule applies to all fragments and non-fragments,
	//  true - the rule applies to only non-first fragments.
	Fragment *bool `xml:"Fragment,omitempty"`
	Logging  *bool `xml:"Logging,omitempty"`
	// Counts times the ACL rule has been matched.
	Counting *bool `xml:"Counting,omitempty"`
	// SrcAny - The flag of matching any IP address.
	SrcAny *bool `xml:"SrcAny,omitempty"`
	// SrcIPv4 - Source IP address.
	SrcIPv4 *SrcIPv4 `xml:"SrcIPv4,omitempty"`
	// Rule comment,
	// a case-sensitive string of 1 to 127 characters.
	Comment string `xml:"Comment,omitempty"`
	// Specifies a time range for the rule, a case-insensitive string of 1 to 32 characters.
	// It must start with an English letter.
	TimeRange string `xml:"TimeRange"`
}

type IPv4NamedBasicRule struct {
	XMLName    xml.Name      `xml:"Rule"`
	GroupIndex string        `xml:"GroupIndex"`
	RuleID     int           `xml:"RuleID"`
	Action     ACLRuleAction `xml:"Action,omitempty"` // Action: 1 - Deny, 2 - Permit
	SrcAny     *bool         `xml:"SrcAny,omitempty"`
	Fragment   *bool         `xml:"Fragment,omitempty"`
	Counting   *bool         `xml:"Counting,omitempty"`
	Logging    *bool         `xml:"Logging,omitempty"`
	Count      int           `xml:"Count,omitempty"`
	Status     ACLRuleStatus `xml:"Status,omitempty"`
	SrcIPv4    *SrcIPv4      `xml:"SrcIPv4,omitempty"`
	// Rule comment,
	// a case-sensitive string of 1 to 127 characters.
	Comment string `xml:"Comment,omitempty"`
}

type SrcIPv4 struct {
	SrcIPv4Addr     string `xml:"SrcIPv4Addr"`
	SrcIPv4Wildcard string `xml:"SrcIPv4Wildcard"`
}

func (ip *SrcIPv4) String() string {
	mask, err := wildcardToPrefix(ip.SrcIPv4Wildcard)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s/%v", ip.SrcIPv4Addr, mask)
}

type DstIPv4 struct {
	DstIPv4Addr     string `xml:"DstIPv4Addr"`
	DstIPv4Wildcard string `xml:"DstIPv4Wildcard"`
}

func (ip *DstIPv4) String() string {
	mask, err := wildcardToPrefix(ip.DstIPv4Wildcard)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s/%v", ip.DstIPv4Addr, mask)
}

type SrcPort struct {
	// SrcPortOp:
	// 1 - lt    less than given port number
	// 2 - eq    Equal to given port number
	// 3 - gt    Greater than given port number
	// 4 - neq   Not equal to given port number
	// 5 - range Between two port numbers
	SrcPortOp ACLRulePortOperation `xml:"SrcPortOp"`
	// SrcPortValue1 specify a source port
	SrcPortValue1 int `xml:"SrcPortValue1"`
	// SrcPortValue2 used only in range case
	SrcPortValue2 int `xml:"SrcPortValue2,omitempty"`
}

type DstPort struct {
	// DstPortOp:
	// 1 - lt    less than given port number
	// 2 - eq    Equal to given port number
	// 3 - gt    Greater than given port number
	// 4 - neq   Not equal to given port number
	// 5 - range Between two port numbers
	DstPortOp ACLRulePortOperation `xml:"DstPortOp"`
	// DstPortValue1 specify a destination port
	DstPortValue1 int `xml:"DstPortValue1"`
	// DstPortValue2 used only in range case
	DstPortValue2 int `xml:"DstPortValue2,omitempty"`
}

type IPv6AdvanceRule struct {
	XMLName        xml.Name      `xml:"Rule"`
	GroupID        int           `xml:"GroupID"`
	RuleID         int           `xml:"RuleID"`
	Action         ACLRuleAction `xml:"Action,omitempty"`
	ProtocolType   int           `xml:"ProtocolType,omitempty"`
	Fragment       *bool         `xml:"Fragment,omitempty"`
	RoutingTypeAny bool          `xml:"RoutingTypeAny,omitempty"`
	HopTypeAny     bool          `xml:"HopTypeAny,omitempty"`
	SrcAny         *bool         `xml:"SrcAny,omitempty"`
	DstAny         *bool         `xml:"DstAny,omitempty"`
	SrcIPv6        *SrcIPv6      `xml:"SrcIPv6,omitempty"`
	DstIPv6        *DstIPv6      `xml:"DstIPv6,omitempty"`
	SrcPort        *SrcPort      `xml:"SrcPort,omitempty"`
	DstPort        *DstPort      `xml:"DstPort,omitempty"`
	// Rule comment,
	// a case-sensitive string of 1 to 127 characters.
	Comment string `xml:"Comment,omitempty"`
}

type IPv6NamedAdvanceRule struct {
	XMLName xml.Name `xml:"Rule"`
	// GroupIndex - Acl Group Name or Index.
	// If it's Index, range from 3000 to 3999.
	GroupIndex string `xml:"GroupIndex"`
	// Rule ID. Range from 0 to 65534
	// If you set this column to 65535, the system automatically assigns a new rule ID.
	RuleID int `xml:"RuleID"`
	// Action on packets matching the rule.
	// Action: 1 - Deny, 2 - Permit
	Action ACLRuleAction `xml:"Action"`
	// ProtocolType defines:
	// Protocol number INTEGER<0-255>, 256 - any IP protocol
	// 6 - TCP
	// 17 - UDP
	// 58 - ICMPv6
	// ...
	// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	ProtocolType int `xml:"ProtocolType"`
	// SrcAny - the flag of matching any source IPv6 address.
	SrcAny *bool `xml:"SrcAny,omitempty"`
	// SrcIPv6 - Source IPv6, including SrcIPv6Address and SrcIPv6Prefix.
	SrcIPv6 *SrcIPv6 `xml:"SrcIPv6,omitempty"`
	// DstAny - the flag of matching any destination IPv6 address.
	DstAny *bool `xml:"DstAny,omitempty"`
	// DstIPv6 - Destination IPv6.
	DstIPv6 *DstIPv6 `xml:"DstIPv6,omitempty"`
	// DSCP - the value of DSCP of IPv6 packet.
	// Range: 0 - 63 inclusive.
	DSCP int `xml:"DSCP,omitempty"`
	// The value of flow label of IPv6 packet header.
	// Range: 0 - 1048575 inclusive.
	FlowLabel int `xml:"FlowLabel,omitempty"`
	// RoutingTypeAny - the flag of matching any routing header type.
	RoutingTypeAny bool `xml:"RoutingTypeAny,omitempty"`
	// RoutingTypeValue - the value of routing header type.
	// Range: 0 - 255 inclusive.
	RoutingTypeValue int      `xml:"RoutingTypeValue,omitempty"`
	SrcPort          *SrcPort `xml:"SrcPort,omitempty"`
	DstPort          *DstPort `xml:"DstPort,omitempty"`
	// Fragment - hhe flag of matching fragmented packet.
	//  0: the rule applies to all fragments and non-fragments,
	//  1: the rule applies to only non-first fragments.
	Fragment *bool `xml:"Fragment,omitempty"`
	Counting *bool `xml:"Counting,omitempty"`
	// Count the number of times the ACL rule has been matched.
	Count uint64 `xml:"Count,omitempty"`
	// Rule status:
	// 1: active,
	// 2: inactive.
	Status ACLRuleStatus `xml:"Status,omitempty"`
	// Logging - enables logs matching packets.
	Logging *bool `xml:"Logging,omitempty"`
	// Rule comment,
	// a case-sensitive string of 1 to 127 characters.
	Comment string `xml:"Comment,omitempty"`
}

type IPv6BasicRule struct {
	XMLName xml.Name `xml:"Rule"`
	// Group ID. Range from 2000 to 2999.
	GroupID int `xml:"GroupID"`
	// Rule ID. Range from 0 to 65534.
	// If you set this column to 65535, the system automatically assigns a new rule ID.
	RuleID int `xml:"RuleID"`
	// Action on packets matching the rule.
	// Action: 1 - Deny, 2 - Permit.
	Action ACLRuleAction `xml:"Action,omitempty"`
	// Rule status.
	// Status: 1 - active, 2 - inactive.
	Status ACLRuleStatus `xml:"Status,omitempty"`
	Count  int64         `xml:"Count,omitempty"`
	// The value of routing header type.
	// 0 .. 255
	RoutingTypeValue int `xml:"RoutingTypeValue,omitempty"`
	// RoutingTypeAny - the flag of matching any routing header type.
	RoutingTypeAny bool `xml:"RoutingTypeAny,omitempty"`
	// Fragment - the flag of matching fragmented packet.
	//  false - the rule applies to all fragments and non-fragments,
	//  true - the rule applies to only non-first fragments.
	Fragment *bool `xml:"Fragment,omitempty"`
	Logging  *bool `xml:"Logging,omitempty"`
	// Counts times the ACL rule has been matched.
	Counting *bool `xml:"Counting,omitempty"`
	// SrcAn - the flag of matching any source IPv6 address.
	SrcAny *bool `xml:"SrcAny,omitempty"`
	// Source IPv6, including SrcIPv6Address and SrcIPv6Prefix.
	SrcIPv6 *SrcIPv6 `xml:"SrcIPv6,omitempty"`
	// Rule comment,
	// a case-sensitive string of 1 to 127 characters.
	Comment string `xml:"Comment,omitempty"`
}

type IPv6NamedBasicRule struct {
	XMLName xml.Name `xml:"Rule"`
	// GroupIndex - Acl Group Name or Index.
	// If it's Index, range from 2000 to 2999.
	GroupIndex string `xml:"GroupIndex"`
	// Rule ID. Range from 0 to 65534
	// If you set this column to 65535, the system automatically assigns a new rule ID.
	RuleID int `xml:"RuleID"`
	// Action on packets matching the rule.
	// Action: 1 - Deny, 2 - Permit
	Action ACLRuleAction `xml:"Action,omitempty"`
	// SrcAny - the flag of matching any source IPv6 address.
	SrcAny *bool `xml:"SrcAny,omitempty"`
	// SrcIPv6 - Source IPv6, including SrcIPv6Address and SrcIPv6Prefix.
	SrcIPv6 *SrcIPv6 `xml:"SrcIPv6,omitempty"`
	// RoutingTypeAny - the flag of matching any routing header type.
	RoutingTypeAny bool `xml:"RoutingTypeAny,omitempty"`
	// RoutingTypeValue - the value of routing header type.
	// Range: 0 - 255 inclusive.
	RoutingTypeValue int `xml:"RoutingTypeValue,omitempty"`
	// Fragment - hhe flag of matching fragmented packet.
	//  0: the rule applies to all fragments and non-fragments,
	//  1: the rule applies to only non-first fragments.
	Fragment *bool `xml:"Fragment,omitempty"`
	Counting *bool `xml:"Counting,omitempty"`
	// Count the number of times the ACL rule has been matched.
	Count uint64 `xml:"Count,omitempty"`
	// Rule status:
	// 1: active,
	// 2: inactive.
	Status ACLRuleStatus `xml:"Status,omitempty"`
	// Logging - enables logs matching packets.
	Logging *bool `xml:"Logging,omitempty"`
	// Rule comment,
	// a case-sensitive string of 1 to 127 characters.
	Comment string `xml:"Comment,omitempty"`
}

type SrcIPv6 struct {
	SrcIPv6Addr   string `xml:"SrcIPv6Addr"`
	SrcIPv6Prefix string `xml:"SrcIPv6Prefix"`
}

func (ip *SrcIPv6) String() string {
	return fmt.Sprintf("%s/%s", ip.SrcIPv6Addr, ip.SrcIPv6Prefix)
}

type DstIPv6 struct {
	DstIPv6Addr   string `xml:"DstIPv6Addr"`
	DstIPv6Prefix string `xml:"DstIPv6Prefix"`
}

func (ip *DstIPv6) String() string {
	return fmt.Sprintf("%s/%s", ip.DstIPv6Addr, ip.DstIPv6Prefix)
}

type MACRule struct {
	XMLName    xml.Name      `xml:"Rule"`
	GroupID    int           `xml:"GroupID"`
	RuleID     int           `xml:"RuleID"`
	Action     ACLRuleAction `xml:"Action,omitempty"`
	SrcMACAddr SrcMACAddr    `xml:"SrcMACAddr,omitempty"`
	DstMACAddr DstMACAddr    `xml:"DstMACAddr,omitempty"`
	Protocol   Protocol      `xml:"Protocol,omitempty"`
}

type SrcMACAddr struct {
	SrcMACAddress string `xml:"SrcMACAddress"`
	//SrcMACMask represents 48-bit hardware address mask
	// ffff-ffff-ffff - one MAC
	// 0000-0000-0000 - all MAC
	SrcMACMask string `xml:"SrcMACMask"`
}

type DstMACAddr struct {
	DstMACAddress string `xml:"DstMACAddress"`
	DstMACMask    string `xml:"DstMACMask"`
}

// Protocol Represents EtherType
// https://tools.ietf.org/html/rfc7042#appendix-B
// For example:
// 0800  IPv4
// 86DD  IPv6
// 0806  ARP
type Protocol struct {
	ProtocolType     string `xml:"ProtocolType"`
	ProtocolTypeMask string `xml:"ProtocolTypeMask"`
}

type TCPFlag struct {
	XMLName xml.Name `xml:"TcpFlag"`
	SYN     int      `xml:"SYN,omitempty"`
	ACK     int      `xml:"ACK,omitempty"`
	FIN     int      `xml:"FIN,omitempty"`
	RST     int      `xml:"RST,omitempty"`
	PSH     int      `xml:"PSH,omitempty"`
	URG     int      `xml:"URG,omitempty"`
}

type ICMP struct {
	XMLName  xml.Name `xml:"ICMP"`
	ICMPType int      `xml:"ICMPType,omitempty"`
	ICMPCode int      `xml:"ICMPCode,omitempty"`
}

type PfilterDefAction struct {
	XMLName       xml.Name      `xml:"PfilterDefAction"`
	DefaultAction ACLRuleAction `xml:"DefaultAction"` // ACL Default Action. 1:Permit, 2:Deny.
}

// PfilterApply table contains information about packet filter application.
type PfilterApply struct {
	XMLName  xml.Name  `xml:"PfilterApply"`
	Pfilters []Pfilter `xml:"Pfilter"`
}

type Pfilter struct {
	XMLName      xml.Name          `xml:"Pfilter"`
	AppObjType   int               `xml:"AppObjType"`            // Object type: 1 - interface, 2 - vlan, 3 - global.
	AppObjIndex  int               `xml:"AppObjIndex"`           // Object Index.
	AppDirection ACLApplyDirection `xml:"AppDirection"`          // Apply Direction: 1 - inbound, 2 - outbound.
	AppACLType   ACLGroupType      `xml:"AppAclType"`            // ACL Group type: 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined, 5 - default.
	AppACLGroup  string            `xml:"AppAclGroup"`           // ACL Name or Index
	HardCount    int               `xml:"HardCount,omitempty"`   // Hardware count flag: 1 - true, 2 - false. Default:false
	AppSequence  int               `xml:"AppSequence,omitempty"` // 1-4294967295
}

// PfilterGroupRunInfo table contains running information about packet filter ACLs on application objects.
type PfilterGroupRunInfo struct {
	XMLName       xml.Name       `xml:"PfilterGroupRunInfo"`
	GroupsRunInfo []GroupRunInfo `xml:"GroupRunInfo"`
}

type GroupRunInfo struct {
	XMLName             xml.Name          `xml:"GroupRunInfo"`
	AppObjType          int               `xml:"AppObjType"`          // Object type: 1 - interface, 2 - vlan, 3 - global.
	AppObjIndex         int               `xml:"AppObjIndex"`         // Object Index.
	AppDirection        ACLApplyDirection `xml:"AppDirection"`        // Apply Direction: 1 - inbound, 2 - outbound.
	AppACLType          ACLGroupType      `xml:"AppAclType"`          // ACL Group type: 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined, 5 - default.
	AppACLGroup         string            `xml:"AppAclGroup"`         // ACL Name or Index.
	ACLGroupStatus      int               `xml:"AclGroupStatus"`      // The status of ACL group applied: 1 - success, 2 - failed, 3 - partialSuccess.
	ACLGroupCountStatus int               `xml:"AclGroupCountStatus"` // The status of enabling hardware count: 1 - success, 2 - failed, 3 - partialSuccess.
	ACLGroupPermitPkts  int               `xml:"AclGroupPermitPkts"`  // The number of packets permitted.
	ACLGroupPermitBytes int               `xml:"AclGroupPermitBytes"` // The number of bytes permitted.
	ACLGroupDenyPkts    int               `xml:"AclGroupDenyPkts"`    // The number of packets denied.
	ACLGroupDenyBytes   int               `xml:"AclGroupDenyBytes"`   // The number of bytes denied.
}

// ACLGroupType
//
//	1 - IPv4,
//	2 - IPv6,
//	3 - MAC,
//	4 - User-defined,
//	5 - Default.
type ACLGroupType int

func (gType ACLGroupType) String() string {
	switch gType {
	case ACLGroupTypeIPv4:
		return ACLGroupTypeIPv4String
	case ACLGroupTypeIPv6:
		return ACLGroupTypeIPv6String
	case ACLGroupTypeMAC:
		return ACLGroupTypeMACString
	case ACLGroupTypeUserDefined:
		return ACLGroupTypeUserDefinedString
	case ACLGroupTypeDefault:
		return ACLGroupTypeDefaultString
	}

	return UnknownString
}

// ACLGroupCategory
//
//	0 - Invalid/None,
//	1 - Basic,
//	2 - Advanced.
type ACLGroupCategory int

func (category ACLGroupCategory) String() string {
	switch category {
	case ACLGroupCategoryBasic:
		return ACLGroupCategoryBasicString
	case ACLGroupCategoryAdvanced:
		return ACLGroupCategoryAdvancedString
	case ACLGroupCategoryNone:
		return ACLGroupCategoryNoneString
	}

	return UnknownString
}

// ACLRuleStatus
//
//	1 - Active,
//	2 - Inactive.
type ACLRuleStatus int

func (status ACLRuleStatus) String() string {
	switch status {
	case ACLRuleStatusActive:
		return AclRuleStatusActiveString
	case ACLRuleStatusInactive:
		return ACLRuleStatusInactiveString
	}

	return UnknownString
}

// ACLRuleAction
//
//	1 - Deny,
//	2 - Permit
type ACLRuleAction int

func (action ACLRuleAction) String() string {
	switch action {
	case ACLRuleActionDeny:
		return ACLRuleActionDenyString
	case ACLRuleActionPermit:
		return ACLRuleActionPermitString
	}

	return UnknownString
}

// ACLApplyDirection
//
//	1 - Inbound,
//	2 - Outbound.
type ACLApplyDirection int

func (direction ACLApplyDirection) String() string {
	switch direction {
	case PFilterApplyDirectionInbound:
		return PfilterApplyDirectionInboundString
	case PFilterApplyDirectionOutbound:
		return PfilterApplyDirectionOutboundString
	}

	return UnknownString
}

// ACLRulePortOperation
//
//	1 - Less
//	2 - Equal
//	3 - Greater
//	4 - NotEqual
//	5 - Range
type ACLRulePortOperation int

func (operation ACLRulePortOperation) String() string {
	switch operation {
	case OperationLess:
		return OperationLessString
	case OperationEqual:
		return OperationEqualString
	case OperationGreater:
		return OperationGreaterString
	case OperationNotEqual:
		return OperationNotEqualString
	case OperationRange:
		return OperationRangeString
	}

	return UnknownString
}

func wildcardToPrefix(wildcardAddress string) (int, error) {
	wildcardOctets := strings.Split(wildcardAddress, ".")
	if len(wildcardOctets) != 4 {
		return -1, fmt.Errorf("bad wildcard length")
	}

	binaryString := ""

	for _, octetString := range wildcardOctets {
		octetInt, err := strconv.Atoi(octetString)
		if err != nil {
			return -1, fmt.Errorf("bad octet '%s': %w", octetString, err)
		}

		binaryString = fmt.Sprintf("%s%08s", binaryString, strconv.FormatInt(int64(octetInt), 2))
	}

	if len(binaryString) != 32 {
		return -1, fmt.Errorf("bad wildcard string '%s'", binaryString)
	}

	pattern := "^0{0,32}1{0,32}$"

	match, _ := regexp.MatchString(pattern, binaryString)
	if !match {
		return -1, fmt.Errorf("bad wildcard string '%s'", binaryString)
	}

	return strings.Count(binaryString, "0"), nil
}
