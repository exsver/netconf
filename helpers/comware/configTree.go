package comware

import "encoding/xml"

//Data used to collect information from device
type Data struct {
	XMLName   xml.Name   `xml:"data"`
	Top       *Top       `xml:"top"`
	SavePoint *SavePoint `xml:"save-point"`
}

type Filter struct {
	XMLName xml.Name `xml:"filter"`
	Top     Top      `xml:"top"`
}

//Top used to configure device
type Top struct {
	XMLName         xml.Name         `xml:"top"`
	ACL             *ACL             `xml:"ACL"`
	ARP             *ARP             `xml:"ARP"`
	Device          *Device          `xml:"Device"`
	DHCP            *DHCP            `xml:"DHCP"`
	DHCPSP          *DHCPSP          `xml:"DHCPSP"`
	IPCIM           *IPCIM           `xml:"IPCIM"`
	Ifmgr           *Ifmgr           `xml:"Ifmgr"`
	MAC             *MAC             `xml:"MAC"`
	MGROUP          *MGROUP          `xml:"MGROUP"`
	ResourceMonitor *ResourceMonitor `xml:"ResourceMonitor"`
	STP             *STP             `xml:"STP"`
	Syslog          *Syslog          `xml:"Syslog"`
	VLAN            *VLAN            `xml:"VLAN"`
	IRF             *IRF             `xml:"IRF"`
}

type ACL struct {
	/* top level
	   ACL
	     Base
	     Capability
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
	     MACRules
	       []MACRule
	*/
	Groups                *Groups                `xml:"Groups"`
	NamedGroups           *NamedGroups           `xml:"NamedGroups"`
	IPv4AdvanceRules      *IPv4AdvanceRules      `xml:"IPv4AdvanceRules"`
	IPv4NamedAdvanceRules *IPv4NamedAdvanceRules `xml:"IPv4NamedAdvanceRules"`
	IPv4BasicRules        *IPv4BasicRules        `xml:"IPv4BasicRules"`
	IPv4NamedBasicRules   *IPv4NamedBasicRules   `xml:"IPv4NamedBasicRules"`
	IPv6AdvanceRules      *IPv6AdvanceRules      `xml:"IPv6AdvanceRules"`
	IPv6BasicRules        *IPv6BasicRules        `xml:"IPv6BasicRules"`
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
	GroupType int `xml:"GroupType"`
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
	GroupType int `xml:"GroupType"`
	// GroupCategory specifies the category of ACL: 0 - invalid, 1 - basic, 2 - advanced.
	// The value range depends on the GroupType column.
	// - 1 to 2 if GroupType is 1 or 2.
	//    basic ACL: 1.
	//    advanced ACL: 2.
	// - 0 if GroupType is 3 or 4.
	GroupCategory int `xml:"GroupCategory"`
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

// IPv6BasicRules table contains information about IPv6 basic ACL rules.
type IPv6BasicRules struct {
	IPv6BasicRules []IPv6BasicRule `xml:"Rule"`
}

// MACRules table contains Ethernet frame header ACL rule information.
type MACRules struct {
	MACRule []MACRule `xml:"Rule"`
}

type IPv4AdvanceRule struct {
	XMLName xml.Name `xml:"Rule"`
	GroupID int      `xml:"GroupID"`
	//RuleID int in range 0-65534
	RuleID int `xml:"RuleID"`
	//Action: 1 - Deny, 2 - Permit
	Action int `xml:"Action"`
	//ProtocolType defines:
	//Protocol number INTEGER<0-255>, 256 - any IP protocol
	// 1 - ICMP
	// 6 - TCP
	// 17 - UDP
	// ...
	//https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	ProtocolType int      `xml:"ProtocolType,omitempty"`
	Count        int      `xml:"Count,omitempty"`
	Status       int      `xml:"Status,omitempty"`
	Fragment     bool     `xml:"Fragment,omitempty"`
	Logging      bool     `xml:"Logging,omitempty"`
	Counting     bool     `xml:"Counting,omitempty"`
	SrcAny       bool     `xml:"SrcAny,omitempty"`
	DstAny       bool     `xml:"DstAny,omitempty"`
	SrcIPv4      *SrcIPv4 `xml:"SrcIPv4,omitempty"`
	DstIPv4      *DstIPv4 `xml:"DstIPv4,omitempty"`
	SrcPort      *SrcPort `xml:"SrcPort,omitempty"`
	DstPort      *DstPort `xml:"DstPort,omitempty"`
}

type IPv4NamedAdvanceRule struct {
	XMLName      xml.Name `xml:"Rule"`
	GroupIndex   string   `xml:"GroupIndex"`
	RuleID       int      `xml:"RuleID"`
	Action       int      `xml:"Action"` //Action: 1 - Deny, 2 - Permit
	ProtocolType int      `xml:"ProtocolType,omitempty"`
	Count        int      `xml:"Count,omitempty"`
	Status       int      `xml:"Status,omitempty"`
	Fragment     bool     `xml:"Fragment,omitempty"` // If an ACL is for QoS traffic classification or packet filtering do not specify the fragment.
	Logging      bool     `xml:"Logging,omitempty"`  // The logging takes effect only when the module (for example, packet filtering) that uses the ACL supports logging.
	Counting     bool     `xml:"Counting,omitempty"`
	SrcAny       bool     `xml:"SrcAny,omitempty"`
	DstAny       bool     `xml:"DstAny,omitempty"`
	SrcIPv4      *SrcIPv4 `xml:"SrcIPv4,omitempty"`
	DstIPv4      *DstIPv4 `xml:"DstIPv4,omitempty"`
	SrcPort      *SrcPort `xml:"SrcPort,omitempty"`
	DstPort      *DstPort `xml:"DstPort,omitempty"`
}

type IPv4BasicRule struct {
	XMLName  xml.Name `xml:"Rule"`
	GroupID  int      `xml:"GroupID"`
	RuleID   int      `xml:"RuleID"`
	Action   int      `xml:"Action"` //Action: 1 - Deny, 2 - Permit
	Fragment bool     `xml:"Fragment,omitempty"`
	Counting bool     `xml:"Counting,omitempty"`
	SrcAny   bool     `xml:"SrcAny,omitempty"`
	SrcIPv4  *SrcIPv4 `xml:"SrcIPv4,omitempty"`
}

type IPv4NamedBasicRule struct {
	XMLName    xml.Name `xml:"Rule"`
	GroupIndex string   `xml:"GroupIndex"`
	RuleID     int      `xml:"RuleID"`
	Action     int      `xml:"Action"` //Action: 1 - Deny, 2 - Permit
	SrcAny     bool     `xml:"SrcAny,omitempty"`
	Fragment   bool     `xml:"Fragment,omitempty"`
	Counting   bool     `xml:"Counting,omitempty"`
	Logging    bool     `xml:"Logging,omitempty"`
	Count      int      `xml:"Count,omitempty"`
	Status     int      `xml:"Status,omitempty"`
	SrcIPv4    *SrcIPv4 `xml:"SrcIPv4,omitempty"`
}

type SrcIPv4 struct {
	SrcIPv4Addr     string `xml:"SrcIPv4Addr"`
	SrcIPv4Wildcard string `xml:"SrcIPv4Wildcard"`
}

type DstIPv4 struct {
	DstIPv4Addr     string `xml:"DstIPv4Addr"`
	DstIPv4Wildcard string `xml:"DstIPv4Wildcard"`
}

type SrcPort struct {
	//DstPortOp:
	// 1 - lt    less than given port number
	// 2 - eq    Equal to given port number
	// 3 - gt    Greater than given port number
	// 4 - neq   Not equal to given port number
	// 5 - range Between two port numbers
	SrcPortOp int `xml:"SrcPortOp"`
	// DstPortValue1 specify a destination port
	SrcPortValue1 int `xml:"SrcPortValue1"`
	// DstPortValue2 used only in range case
	SrcPortValue2 int `xml:"SrcPortValue2"`
}

type DstPort struct {
	DstPortOp     int `xml:"DstPortOp"`
	DstPortValue1 int `xml:"DstPortValue1"`
	DstPortValue2 int `xml:"DstPortValue2"`
}

type IPv6AdvanceRule struct {
	XMLName        xml.Name `xml:"Rule"`
	GroupID        int      `xml:"GroupID"`
	RuleID         int      `xml:"RuleID"`
	Action         int      `xml:"Action"`
	ProtocolType   int      `xml:"ProtocolType,omitempty"`
	Fragment       bool     `xml:"Fragment,omitempty"`
	RoutingTypeAny bool     `xml:"RoutingTypeAny,omitempty"`
	HopTypeAny     bool     `xml:"HopTypeAny,omitempty"`
	SrcAny         bool     `xml:"SrcAny,omitempty"`
	DstAny         bool     `xml:"DstAny,omitempty"`
	SrcIPv6        SrcIPv6  `xml:"SrcIPv6,omitempty"`
	DstIPv6        DstIPv6  `xml:"DstIPv6,omitempty"`
	SrcPort        SrcPort  `xml:"SrcPort,omitempty"`
	DstPort        DstPort  `xml:"DstPort,omitempty"`
}

type IPv6BasicRule struct {
	XMLName        xml.Name `xml:"Rule"`
	GroupID        int      `xml:"GroupID"`
	RuleID         int      `xml:"RuleID"`
	Action         int      `xml:"Action"`
	Fragment       bool     `xml:"Fragment,omitempty"`
	RoutingTypeAny bool     `xml:"RoutingTypeAny,omitempty"`
	SrcAny         bool     `xml:"SrcAny,omitempty"`
	SrcIPv6        SrcIPv6  `xml:"SrcIPv6,omitempty"`
}

type SrcIPv6 struct {
	SrcIPv6Addr   string `xml:"SrcIPv6Addr"`
	SrcIPv6Prefix string `xml:"SrcIPv6Prefix"`
}

type DstIPv6 struct {
	DstIPv6Addr   string `xml:"DstIPv6Addr"`
	DstIPv6Prefix string `xml:"DstIPv6Prefix"`
}

type MACRule struct {
	XMLName    xml.Name   `xml:"Rule"`
	GroupID    int        `xml:"GroupID"`
	RuleID     int        `xml:"RuleID"`
	Action     int        `xml:"Action"`
	SrcMACAddr SrcMACAddr `xml:"SrcMACAddr,omitempty"`
	DstMACAddr DstMACAddr `xml:"DstMACAddr,omitempty"`
	Protocol   Protocol   `xml:"Protocol,omitempty"`
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

//Protocol Represents EtherType
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
	XMLName       xml.Name `xml:"PfilterDefAction"`
	DefaultAction int      `xml:"DefaultAction"` // ACL Default Action. 1:Permit, 2:Deny.
}

// PfilterApply table contains information about packet filter application.
type PfilterApply struct {
	XMLName  xml.Name  `xml:"PfilterApply"`
	Pfilters []Pfilter `xml:"Pfilter"`
}

type Pfilter struct {
	XMLName      xml.Name `xml:"Pfilter"`
	AppObjType   int      `xml:"AppObjType"`            // Object type: 1 - interface, 2 - vlan, 3 - global.
	AppObjIndex  int      `xml:"AppObjIndex"`           // Object Index.
	AppDirection int      `xml:"AppDirection"`          // Apply Direction: 1 - inbound, 2 - outbound.
	AppACLType   int      `xml:"AppAclType"`            // ACL Group type: 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined, 5 - default.
	AppACLGroup  string   `xml:"AppAclGroup"`           // ACL Name or Index
	HardCount    int      `xml:"HardCount,omitempty"`   // Hardware count flag: 1 - true, 2 - false. Default:false
	AppSequence  int      `xml:"AppSequence,omitempty"` // 1-4294967295
}

// PfilterGroupRunInfo table contains running information about packet filter ACLs on application objects.
type PfilterGroupRunInfo struct {
	XMLName       xml.Name       `xml:"PfilterGroupRunInfo"`
	GroupsRunInfo []GroupRunInfo `xml:"GroupRunInfo"`
}

type GroupRunInfo struct {
	XMLName             xml.Name `xml:"GroupRunInfo"`
	AppObjType          int      `xml:"AppObjType"`          // Object type: 1 - interface, 2 - vlan, 3 - global.
	AppObjIndex         int      `xml:"AppObjIndex"`         // Object Index.
	AppDirection        int      `xml:"AppDirection"`        // Apply Direction: 1 - inbound, 2 - outbound.
	AppACLType          int      `xml:"AppAclType"`          // ACL Group type: 1 - IPv4, 2 - IPv6, 3 - MAC, 4 - User-defined, 5 - default.
	AppACLGroup         string   `xml:"AppAclGroup"`         // ACL Name or Index.
	ACLGroupStatus      int      `xml:"AclGroupStatus"`      // The status of ACL group applied: 1 - success, 2 - failed, 3 - partialSuccess.
	ACLGroupCountStatus int      `xml:"AclGroupCountStatus"` // The status of enabling hardware count: 1 - success, 2 - failed, 3 - partialSuccess.
	ACLGroupPermitPkts  int      `xml:"AclGroupPermitPkts"`  // The number of packets permitted.
	ACLGroupPermitBytes int      `xml:"AclGroupPermitBytes"` // The number of bytes permitted.
	ACLGroupDenyPkts    int      `xml:"AclGroupDenyPkts"`    // The number of packets denied.
	ACLGroupDenyBytes   int      `xml:"AclGroupDenyBytes"`   // The number of bytes denied.
}

type ARP struct {
	/* top level
		ARP
	      ArpAttackProtect
	      ArpAuthorized
	      ArpDetection
	        []Detect
	      ArpDetectionTrust
	        []Interface
	      ArpFilterSource
	        []FilterSource
	      ArpGratuitous
	      ArpGratuitousInterval
	        []GratuitousInterval
	      ArpLearnLimit
	        []LearnLimit
	      ArpProxy
	        []Proxy
	      ArpRateLimit
	        []RateLimit
	      ArpRateLimitLog
	      ArpSnooping
	        []Snooping
	      ArpSpecification
	      ArpTable
	        []ArpEntry
	*/
	ArpDetection      *ArpDetection      `xml:"ArpDetection"`
	ArpDetectionTrust *ArpDetectionTrust `xml:"ArpDetectionTrust"`
	ArpFilterSource   *ArpFilterSource   `xml:"ArpFilterSource"`
	ArpGratuitous     *ArpGratuitous     `xml:"ArpGratuitous"`
	ArpRateLimit      *ArpRateLimit      `xml:"ArpRateLimit"`
	ArpRateLimitLog   *ArpRateLimitLog   `xml:"ArpRateLimitLog"`
	ArpSpecification  *ArpSpecification  `xml:"ArpSpecification"`
	ArpTable          *ArpTable          `xml:"ArpTable"`
}

// ArpDetection table contains the information about arp detections.
type ArpDetection struct {
	XMLName   xml.Name `xml:"ArpDetection"`
	ArpDetect []Detect `xml:"Detect"`
}

type Detect struct {
	XMLName                    xml.Name `xml:"Detect"`
	VLANID                     int      `xml:"VLANID"`
	DetectionEnable            bool     `xml:"DetectionEnable"`
	RestrictedForwardingEnable bool     `xml:"RestrictedForwardingEnable"`
}

// ArpDetectionTrust table contains the information of interface that arp trusted.
type ArpDetectionTrust struct {
	XMLName    xml.Name                     `xml:"ArpDetectionTrust"`
	Interfaces []ArpDetectionTrustInterface `xml:"Interface"`
}

type ArpDetectionTrustInterface struct {
	XMLName     xml.Name `xml:"Interface"`
	IfIndex     int      `xml:"IfIndex"`
	TrustEnable bool     `xml:"TrustEnable"`
}

// ArpFilterSource table contains the information about arp gateway protection and protected ip address.
type ArpFilterSource struct {
	XMLName       xml.Name       `xml:"ArpFilterSource"`
	FilterSources []FilterSource `xml:"FilterSource"`
}

// ArpGratuitous contains gratuitous ARP packets enable information.
type ArpGratuitous struct {
	XMLName     xml.Name `xml:"ArpGratuitous"`
	LearnEnable bool     `xml:"LearnEnable"`
	SendEnable  bool     `xml:"SendEnable"`
}

type FilterSource struct {
	XMLName     xml.Name `xml:"FilterSource"`
	IfIndex     int      `xml:"IfIndex"`
	Ipv4Address string   `xml:"Ipv4Address"`
}

// ArpSpecification table contains ARP specification information.
type ArpSpecification struct {
	XMLName              xml.Name `xml:"ArpSpecification"`
	SupportMultiport     bool     `xml:"SupportMultiport"`
	SupportDetection     bool     `xml:"SupportDetection"`
	SupportFilter        bool     `xml:"SupportFilter"`
	SupportSnooping      bool     `xml:"SupportSnooping"`
	SupportModeUNI       bool     `xml:"SupportModeUNI"`
	RateLimitDefault     int      `xml:"RateLimitDefault"`
	RateLimitLowerLimit  int      `xml:"RateLimitLowerLimit"`
	RateLimitUpperLimit  int      `xml:"RateLimitUpperLimit"`
	SMACDefaultThreshold int      `xml:"SMACDefaultThreshold"`
	SMACMinThreshold     int      `xml:"SMACMinThreshold"`
	SMACMaxThreshold     int      `xml:"SMACMaxThreshold"`
	IntervalDefault      int      `xml:"IntervalDefault"`
	LearnMaxNumLimit     int      `xml:"LearnMaxNumLimit"`
	LearnMaxNumDefault   int      `xml:"LearnMaxNumDefault"`
}

// ArpRateLimit table contains ARP rate limit information.
type ArpRateLimit struct {
	XMLName             xml.Name                `xml:"ArpRateLimit"`
	RateLimitInterfaces []ArpRateLimitInterface `xml:"RateLimit"`
}

type ArpRateLimitInterface struct {
	XMLName         xml.Name `xml:"RateLimit"`
	IfIndex         int      `xml:"IfIndex"`
	RateLimitEnable bool     `xml:"RateLimitEnable"`
	RateLimitNum    int      `xml:"RateLimitNum"`
}

// ArpRateLimitLog contains ARP rate limit log information.
type ArpRateLimitLog struct {
	XMLName     xml.Name `xml:"ArpRateLimitLog"`
	LogEnable   bool     `xml:"LogEnable"`
	LogInterval int      `xml:"LogInterval,omitempty"`
}

// ArpTable table contains ARP table information.
type ArpTable struct {
	XMLName    xml.Name   `xml:"ArpTable"`
	ArpEntries []ArpEntry `xml:"ArpEntry"`
}

type ArpEntry struct {
	XMLName     xml.Name `xml:"ArpEntry"`
	IfIndex     int      `xml:"IfIndex"`
	VLANID      int      `xml:"VLANID"`
	Ipv4Address string   `xml:"Ipv4Address"`
	MacAddress  string   `xml:"MacAddress"`
	PortIndex   int      `xml:"PortIndex"`
	VrfIndex    int      `xml:"VrfIndex"`
	ArpType     int      `xml:"ArpType"`
}

type IRF struct {
	/* top level
	   IRF
	     Capability
	     Configuration
	     IRFPorts
	       []IRFPort
	     Members
	       []Member
	*/
	Capability    *IRFCapability    `xml:"Capability"`
	Configuration *IRFConfiguration `xml:"Configuration"`
	IRFPorts      *IRFPorts         `xml:"IRFPorts"`
	Members       *Members          `xml:"Members"`
}

type IRFCapability struct {
	XMLName            xml.Name `xml:"Capability"`
	MaxMemberCount     int      `xml:"MaxMemberCount"`
	MaxPriority        int      `xml:"MaxPriority"`
	MaxIfNumPerIRFPort int      `xml:"MaxIfNumPerIRFPort"`
}

type IRFConfiguration struct {
	XMLName     xml.Name `xml:"Configuration"`
	AutoUpgrade string   `xml:"AutoUpgrade"`
	Domain      int      `xml:"Domain"`
	LinkDelay   int      `xml:"LinkDelay"`
	MacPersist  int      `xml:"MacPersist"`
	BridgeMac   string   `xml:"BridgeMac"`
	TopoType    int      `xml:"TopoType"`
	StackMode   bool     `xml:"StackMode"`
	MemberCount int      `xml:"MemberCount"`
}

// IRFPorts table contains configuration of the IRF ports.
type IRFPorts struct {
	XMLName  xml.Name  `xml:"IRFPorts"`
	IRFPorts []IRFPort `xml:"IRFPort"`
}

type IRFPort struct {
	XMLName   xml.Name          `xml:"IRFPort"`
	MemberID  int               `xml:"MemberID,omitempty"`
	Port      int               `xml:"Port,omitempty"`
	Neighbor  int               `xml:"Neighbor"`
	State     int               `xml:"State"`
	Interface *IRFPortInterface `xml:"Interface"`
}

type IRFPortInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfName  string   `xml:"IfName"`
	//Physical IRF port mode
	// 1 Normal, 2 - Enhanced 3 - Extended
	Mode      int `xml:"Mode"`
	LinkState int `xml:"LinkState"`
}

// Members table contains configuration of the IRF members.
type Members struct {
	XMLName xml.Name `xml:"Members"`
	Members []Member `xml:"Member"`
}

type Member struct {
	XMLName     xml.Name `xml:"Member"`
	MemberID    int      `xml:"MemberID,omitempty"`
	Description string   `xml:"Description,omitempty"`
	NewMemberID int      `xml:"NewMemberID,omitempty"`
	Priority    int      `xml:"Priority,omitempty"`
}

type Device struct {
	/* top level
	Device
	  Base
	    TimeZone
	  Boards						 ***ReadOnly***
	    []Board                      ***ReadOnly***
	  CPUs                           ***ReadOnly***
	    []CPU                        ***ReadOnly***
	  SummerTime
	    WeekBased
	  FanDirections
	    []Fan
	  PhysicalEntities               ***ReadOnly***
	    []Entity                     ***ReadOnly***
	  ExtPhysicalEntities	         ***ReadOnly***
	    []ExtPhysicalEntity			 ***ReadOnly***
	  TemperatureSensors			 ***ReadOnly***
	    []Sensor					 ***ReadOnly***
	  Transceivers                   ***ReadOnly***
	    []Interface					 ***ReadOnly***
	  TransceiversChannels			 ***ReadOnly***   use for FortyGigE interfaces
	    []Interface					 ***ReadOnly***
	*/
	Base                 *Base                 `xml:"Base"`
	Boards               *Boards               `xml:"Boards"`
	CPUs                 *CPUs                 `xml:"CPUs"`
	PhysicalEntities     *PhysicalEntities     `xml:"PhysicalEntities"`
	ExtPhysicalEntities  *ExtPhysicalEntities  `xml:"ExtPhysicalEntities"`
	FanDirections        *FanDirections        `xml:"FanDirections"`
	TemperatureSensors   *TemperatureSensors   `xml:"TemperatureSensors"`
	Transceivers         *Transceivers         `xml:"Transceivers"`
	TransceiversChannels *TransceiversChannels `xml:"TransceiversChannels"`
	SummerTime           *SummerTime           `xml:"SummerTime"`
}

type Base struct {
	XMLName         xml.Name `xml:"Base"`
	Uptime          uint64   `xml:"Uptime,omitempty"`          //Time that elapsed after the device started up(units:second)
	HostName        string   `xml:"HostName,omitempty"`        //Device name. String length constraints must be in range(0..64). String must not begin and end with space character
	HostDescription string   `xml:"HostDescription,omitempty"` //Device description.String length constraints must be in range(0..255).
	LocalTime       string   `xml:"LocalTime,omitempty"`
	BridgeMAC       string   `xml:"BridgeMAC,omitempty"`
	LocalBridgeMAC  string   `xml:"LocalBridgeMAC,omitempty"`
	MinChassisNum   int      `xml:"MinChassisNum,omitempty"`
	MaxChassisNum   int      `xml:"MaxChassisNum,omitempty"`
	MinSlotNum      int      `xml:"MinSlotNum,omitempty"`
	MaxSlotNum      int      `xml:"MaxSlotNum,omitempty"`
	MinCPUIDNum     int      `xml:"MinCPUIDNum,omitempty"`
	MaxCPUIDNum     int      `xml:"MaxCPUIDNum,omitempty"`
}

//Boards table contains board information.
type Boards struct {
	XMLName xml.Name `xml:"Boards"`
	Boards  []Board  `xml:"Board"`
}

type Board struct {
	XMLName       xml.Name    `xml:"Board"`
	DeviceNode    *DeviceNode `xml:"DeviceNode"`
	PhysicalIndex int         `xml:"PhysicalIndex"`
	Status        int         `xml:"Status"` //Status of the device node: 1 Absent, 2 Normal, 3 Fault
	Role          int         `xml:"Role"`   //Role of the device node: 1 - Unknown, 2 - ActiveMPU, 3 - StandbyMPU, 4 - LPU. Role is Unknown when status is Absent or Fault.
}

type CPUs struct {
	XMLName xml.Name `xml:"CPUs"`
	CPUs    []CPU    `xml:"CPU"`
}

type CPU struct {
	XMLName       xml.Name `xml:"CPU"`
	PhysicalIndex int      `xml:"PhysicalIndex"`
	Chassis       int      `xml:"Chassis"`
	Slot          int      `xml:"Slot"`
	CPUID         int      `xml:"CPUID"`
	CPUUsage      int      `xml:"CPUUsage"`
}

// PhysicalEntities table contains basic information about an entity.
// ReadOnly struct
type PhysicalEntities struct {
	XMLName          xml.Name         `xml:"PhysicalEntities"`
	PhysicalEntities []PhysicalEntity `xml:"Entity"`
}

//ReadOnly struct
type PhysicalEntity struct {
	XMLName       xml.Name `xml:"Entity"`
	PhysicalIndex int      `xml:"PhysicalIndex"`
	Chassis       int      `xml:"Chassis"`
	Slot          int      `xml:"Slot"`
	SubSlot       int      `xml:"SubSlot"`
	Description   string   `xml:"Description"`
	VendorType    string   `xml:"VendorType"`
	ContainedIn   int      `xml:"ContainedIn"`
	Class         int      `xml:"Class"` //3 - physical node (slot),  6 - PSU,  7 - FAN UNIT, 8 - Sensor, 9 - Board, SubBoard, 10 - Interfaces
	ParentRelPos  int      `xml:"ParentRelPos"`
	Name          string   `xml:"Name"`
	HardwareRev   string   `xml:"HardwareRev"`
	FirmwareRev   string   `xml:"FirmwareRev"`
	SoftwareRev   string   `xml:"SoftwareRev"`
	SerialNumber  string   `xml:"SerialNumber"`
	MfgName       string   `xml:"MfgName"`
	Model         string   `xml:"Model"`
	Alias         string   `xml:"Alias"`
	AssetID       string   `xml:"AssetID"`
	FRU           bool     `xml:"FRU"`
	MfgDate       string   `xml:"MfgDate"`
	Uris          string   `xml:"Uris"`
}

// ExtPhysicalEntities table contains extended information about an entity.
type ExtPhysicalEntities struct {
	XMLName             xml.Name            `xml:"ExtPhysicalEntities"`
	ExtPhysicalEntities []ExtPhysicalEntity `xml:"Entity"`
}

type ExtPhysicalEntity struct {
	XMLName                      xml.Name `xml:"Entity"`
	PhysicalIndex                int      `xml:"PhysicalIndex"`
	Uptime                       int      `xml:"Uptime"`                       //Time that elapsed after the entity started up (units:seconds)
	AdminState                   int      `xml:"AdminState"`                   //1 - Unsupported, 2 - Disabled, 3 - Shut down, 4 - Enabled
	OperState                    int      `xml:"OperState"`                    //1 - Unsupported, 2 - Unusable, 3 - Usable, 4 - Dangerous, cannot be used
	StandbyState                 int      `xml:"StandbyState"`                 //1 - Unsupported, 2 - Hot standby, 3 - Cold standby, 4 - In service
	AlarmLight                   int      `xml:"AlarmLight"`                   //Alarm LED status
	CPUUsage                     int      `xml:"CpuUsage"`                     //CPU use ratio (percentage)
	CPUMaxUsage                  int      `xml:"CpuMaxUsage"`                  //Maximum CPU use ratio (percentage)
	CPUAvgUsage                  int      `xml:"CpuAvgUsage"`                  //Average CPU usage (percentage)
	CPUUsageThreshold            int      `xml:"CpuUsageThreshold"`            //CPU use ratio threshold (percentage)
	MemUsage                     int      `xml:"MemUsage"`                     //Memory use ratio (percentage)
	MemAvgUsage                  int      `xml:"MemAvgUsage"`                  //Average memory usage (percentage)
	MemUsageThreshold            int      `xml:"MemUsageThreshold"`            //Memory use ratio threshold (percentage)
	MemSize                      int      `xml:"MemSize"`                      //Memory size (units: bytes)
	PhyMemSize                   int      `xml:"PhyMemSize"`                   //Physical memory size (units: bytes)
	Temperature                  int      `xml:"Temperature"`                  //Temperature of the entity (units: C)
	TemperatureThreshold         int      `xml:"TemperatureThreshold"`         //High-temperature threshold (units: C)
	TemperatureCriticalThreshold int      `xml:"TemperatureCriticalThreshold"` //Critical high-voltage threshold (units: C)
	TemperatureLowThreshold      int      `xml:"TemperatureLowThreshold"`      //Low temperature threshold (units: C)
	TemperatureShutdownThreshold int      `xml:"TemperatureShutdownThreshold"` //Shutdown temperature threshold (units: C)
	ErrorStatus                  int      `xml:"ErrorStatus"`                  //Error status of the entity. Numerial value must be more than 0
	MAC                          string   `xml:"MAC"`
}

type FanDirections struct {
	XMLName xml.Name `xml:"FanDirections"`
	Fans    []Fan    `xml:"Fan"`
}

type Fan struct {
	XMLName          xml.Name `xml:"Fan"`
	Chassis          int      `xml:"Chassis,omitempty"`
	Slot             int      `xml:"Slot,omitempty"`
	CPUID            int      `xml:"CPUID,omitempty"`
	Direction        int      `xml:"Direction,omitempty"`
	DefaultDirection int      `xml:"DefaultDirection,omitempty"`
}

// TemperatureSensors table contains the temperature sensor information.
type TemperatureSensors struct {
	XMLName xml.Name `xml:"TemperatureSensors"`
	Sensors []Sensor `xml:"Sensor"`
}

type Sensor struct {
	XMLName      xml.Name `xml:"Sensor"`
	Chassis      int      `xml:"Chassis"`
	Slot         int      `xml:"Slot"`
	CPUID        int      `xml:"CPUID"`
	SensorType   int      `xml:"SensorType"`
	SensorIndex  int      `xml:"SensorIndex"`
	LowerLimit   int      `xml:"LowerLimit"`
	WarningLimit int      `xml:"WarningLimit"`
	AlarmLimit   int      `xml:"AlarmLimit"`
}

type Transceivers struct {
	XMLName    xml.Name               `xml:"Transceivers"`
	Interfaces []TransceiverInterface `xml:"Interface"`
}

type TransceiverInterface struct {
	XMLName            xml.Name `xml:"Interface"`
	IfIndex            int      `xml:"IfIndex"`
	Name               string   `xml:"Name"`
	HardwareType       string   `xml:"HardwareType"`
	TransceiverType    string   `xml:"TransceiverType"`
	RevisionNumber     string   `xml:"RevisionNumber"`
	TransceiverErrors  string   `xml:"TransceiverErrors"`
	WaveLength         string   `xml:"WaveLength"`
	VendorName         string   `xml:"VendorName"`
	SerialNumber       string   `xml:"SerialNumber"`
	FiberDiameterType  int      `xml:"FiberDiameterType"`
	TransferDistance   int      `xml:"TransferDistance"`
	Diagnostic         int      `xml:"Diagnostic"`
	CurRxPower         string   `xml:"CurRxPower"`
	MaxRxPower         string   `xml:"MaxRxPower"`
	MinRxPower         string   `xml:"MinRxPower"`
	CurTxPower         string   `xml:"CurTxPower"`
	MaxTxPower         string   `xml:"MaxTxPower"`
	MinTxPower         string   `xml:"MinTxPower"`
	Temperature        string   `xml:"Temperature"`
	Voltage            string   `xml:"Voltage"`
	BiasCurrent        string   `xml:"BiasCurrent"`
	TemperatureHiAlarm string   `xml:"TemperatureHiAlarm"`
	TemperatureLoAlarm string   `xml:"TemperatureLoAlarm"`
	TemperatureHiWarn  string   `xml:"TemperatureHiWarn"`
	TemperatureLoWarn  string   `xml:"TemperatureLoWarn"`
	VccHiAlarm         string   `xml:"VccHiAlarm"`
	VccLoAlarm         string   `xml:"VccLoAlarm"`
	VccHiWarn          string   `xml:"VccHiWarn"`
	VccLoWarn          string   `xml:"VccLoWarn"`
	BiasHiAlarm        string   `xml:"BiasHiAlarm"`
	BiasLoAlarm        string   `xml:"BiasLoAlarm"`
	BiasHiWarn         string   `xml:"BiasHiWarn"`
	BiasLoWarn         string   `xml:"BiasLoWarn"`
	PwrOutHiAlarm      string   `xml:"PwrOutHiAlarm"`
	PwrOutLoAlarm      string   `xml:"PwrOutLoAlarm"`
	PwrOutHiWarn       string   `xml:"PwrOutHiWarn"`
	PwrOutLoWarn       string   `xml:"PwrOutLoWarn"`
	RcvPwrHiAlarm      string   `xml:"RcvPwrHiAlarm"`
	RcvPwrLoAlarm      string   `xml:"RcvPwrLoAlarm"`
	RcvPwrHiWarn       string   `xml:"RcvPwrHiWarn"`
	RcvPwrLoWarn       string   `xml:"RcvPwrLoWarn"`
	VendorOUI          string   `xml:"VendorOUI"`
	Frequency          string   `xml:"Frequency"`
	ActiveITUChannel   string   `xml:"ActiveITUChannel"`
	CurWaveErr         string   `xml:"CurWaveErr"`
	WaveErrHiAlarm     string   `xml:"WaveErrHiAlarm"`
	WaveErrLoAlarm     string   `xml:"WaveErrLoAlarm"`
	CurFreqErr         string   `xml:"CurFreqErr"`
	FreqErrHiAlarm     string   `xml:"FreqErrHiAlarm"`
	FreqErrLoAlarm     string   `xml:"FreqErrLoAlarm"`
}

type TransceiversChannels struct {
	XMLName    xml.Name                         `xml:"TransceiversChannels"`
	Interfaces []TransceiversChannelsInterfaces `xml:"Interface"`
}

type TransceiversChannelsInterfaces struct {
	XMLName            xml.Name `xml:"Interface"`
	IfIndex            int      `xml:"IfIndex"`
	ChannelIndex       int      `xml:"ChannelIndex"`
	ChannelCurTXPower  string   `xml:"ChannelCurTXPower"`
	ChannelCurRXPower  string   `xml:"ChannelCurRXPower"`
	ChannelTemperature string   `xml:"ChannelTemperature"`
	ChannelBiasCurrent string   `xml:"ChannelBiasCurrent"`
}

type SummerTime struct {
	XMLName   xml.Name   `xml:"SummerTime"`
	Name      string     `xml:"Name,omitempty"`
	AddTime   string     `xml:"AddTime,omitempty"`
	WeekBased *WeekBased `xml:"WeekBased"`
}

type WeekBased struct {
	XMLName      xml.Name `xml:"WeekBased"`
	BeginMonth   int      `xml:"BeginMonth,omitempty"`
	BeginWeek    int      `xml:"BeginWeek,omitempty"`
	BeginWeekDay int      `xml:"BeginWeekDay,omitempty"`
	BeginHour    int      `xml:"BeginHour,omitempty"`
	BeginMinute  int      `xml:"BeginMinute,omitempty"`
	BeginSecond  int      `xml:"BeginSecond,omitempty"`
	EndMonth     int      `xml:"EndMonth,omitempty"`
	EndWeek      int      `xml:"EndWeek,omitempty"`
	EndWeekDay   int      `xml:"EndWeekDay,omitempty"`
	EndHour      int      `xml:"EndHour,omitempty"`
	EndMinute    int      `xml:"EndMinute,omitempty"`
	EndSecond    int      `xml:"EndSecond,omitempty"`
}

type DHCPSP struct {
	/* top level
	DHCPSP
	  DHCPSPBindingDatabase
	  DHCPSPConfig
	  DHCPSPInterface
	    []Interface
	  DHCPSPOpt82
	    []Option82
	  DHCPSPSpecification
	*/
	DHCPSPBindingDatabase *DHCPSPBindingDatabase `xml:"DHCPSPBindingDatabase"`
	DHCPSPConfig          *DHCPSPConfig          `xml:"DHCPSPConfig"`
	DHCPSPInterface       *DHCPSPInterface       `xml:"DHCPSPInterface"`
}

type DHCPSPBindingDatabase struct {
	XMLName        xml.Name `xml:"DHCPSPBindingDatabase"`
	UpdateInterval int      `xml:"UpdateInterval"`
}

type DHCPSPConfig struct {
	XMLName      xml.Name `xml:"DHCPSPConfig"`
	DHCPSPEnable bool     `xml:"DHCPSPEnable"`
}

type DHCPSPInterface struct {
	XMLName    xml.Name                `xml:"DHCPSPInterface"`
	Interfaces []DHCPSnoopingInterface `xml:"Interface"`
}

type DHCPSnoopingInterface struct {
	XMLName             xml.Name `xml:"Interface"`
	IfIndex             int      `xml:"IfIndex"`
	BindingRecord       bool     `xml:"BindingRecord"`
	CheckMacAddress     bool     `xml:"CheckMacAddress"`
	CheckRequestMessage bool     `xml:"CheckRequestMessage"`
	Trust               bool     `xml:"Trust"`
	LearnMaxNum         int      `xml:"LearnMaxNum"`
	RateLimitNum        int      `xml:"RateLimitNum"`
}

type DHCP struct {
	/* top level
	DHCP
	  DHCPConfig
	  DHCPIfMode
	  DHCPServerIpPool
	    []IpPool
	  DHCPServerPoolStatic
	    []IpPoolStatic
	  DHCPServerIPInUse         ***ReadOnly***
	    []IpInUse               ***ReadOnly***
	*/
	DHCPConfig           *DHCPConfig           `xml:"DHCPConfig"`
	DHCPIfMode           *DHCPIfMode           `xml:"DHCPIfMode"`
	DHCPServerIPPool     *DHCPServerIPPool     `xml:"DHCPServerIpPool"`
	DHCPServerPoolStatic *DHCPServerPoolStatic `xml:"DHCPServerPoolStatic"`
	DHCPServerIPInUse    *DHCPServerIPInUse    `xml:"DHCPServerIpInUse"`
}

type DHCPConfig struct {
	XMLName      xml.Name          `xml:"DHCPConfig"`
	DHCPEnable   bool              `xml:"DHCPEnable,omitempty"`
	ServerConfig *DHCPServerConfig `xml:"ServerConfig"`
}

type DHCPServerConfig struct {
	XMLName           xml.Name `xml:"ServerConfig"`
	AlwaysBroadcast   bool     `xml:"AlwaysBroadcast,omitempty"`
	IgnoreBOOTP       bool     `xml:"IgnoreBOOTP,omitempty"`
	BOOTPReplyRFC1048 bool     `xml:"BOOTPReplyRFC1048,omitempty"`
	Opt82Enabled      bool     `xml:"Opt82Enabled,omitempty"`
	PingNumber        int      `xml:"PingNumber,omitempty"`  //Valid values are:0-10
	PingTimeout       int      `xml:"PingTimeout,omitempty"` //Valid values are:0-10000
}

type DHCPRelayConfig struct {
	XMLName           xml.Name `xml:"RelayConfig"`
	UserInfoRecord    bool     `xml:"UserInfoRecord,omitempty"`
	UserInfoRefresh   bool     `xml:"UserInfoRefresh,omitempty"`
	UserInfoFlushTime int      `xml:"UserInfoFlushTime,omitempty"` //Valid values are:0-120
}

// DHCPIfMode table contains DHCP interface mode information.
type DHCPIfMode struct {
	XMLName    xml.Name `xml:"DHCPIfMode"`
	Interfaces []IfMode `xml:"IfMode"`
}

type IfMode struct {
	XMLName xml.Name `xml:"IfMode"`
	IfIndex int      `xml:"IfIndex"`
	Mode    int      `xml:"Mode"`
}

// DHCPServerIPPool table contains DHCP server IP pool information.
type DHCPServerIPPool struct {
	XMLName xml.Name `xml:"DHCPServerIpPool"`
	IPPools []IPPool `xml:"IpPool"`
}

type IPPool struct {
	XMLName            xml.Name `xml:"IpPool"`
	PoolIndex          int      `xml:"PoolIndex"`
	PoolName           string   `xml:"PoolName,omitempty"`
	NetworkIpv4Address string   `xml:"NetworkIpv4Address,omitempty"`
	NetworkIpv4Mask    string   `xml:"NetworkIpv4Mask,omitempty"`
	GatewayIpv4Address string   `xml:"GatewayIpv4Address,omitempty"`
	DNSIpv4Address     string   `xml:"DNSIpv4Address,omitempty"`
	StartIpv4Address   string   `xml:"StartIpv4Address,omitempty"`
	EndIpv4Address     string   `xml:"EndIpv4Address,omitempty"`
	LeaseDay           int      `xml:"LeaseDay,omitempty"`
	LeaseHour          int      `xml:"LeaseHour,omitempty"`
	LeaseMinute        int      `xml:"LeaseMinute,omitempty"`
	LeaseSecond        int      `xml:"LeaseSecond,omitempty"`
	LeaseUnlimit       bool     `xml:"LeaseUnlimit,omitempty"`
}

// DHCPServerPoolStatic table contains static binding information of DHCP server IP pool.
type DHCPServerPoolStatic struct {
	XMLName     xml.Name       `xml:"DHCPServerPoolStatic"`
	StaticHosts []IpPoolStatic `xml:"IpPoolStatic"`
}

type IpPoolStatic struct {
	XMLName     xml.Name `xml:"IpPoolStatic"`
	PoolIndex   int      `xml:"PoolIndex"`
	Ipv4Address string   `xml:"Ipv4Address,omitempty"`
	Ipv4Mask    string   `xml:"Ipv4Mask,omitempty"`
	HAddress    string   `xml:"HAddress,omitempty"`
	HType       int      `xml:"HType,omitempty"`
}

// DHCPServerIPInUse table contains binding information about assigned IP addresses
type DHCPServerIPInUse struct {
	XMLName xml.Name  `xml:"DHCPServerIpInUse"`
	IPInUse []IPInUse `xml:"IpInUse"`
}

type IPInUse struct {
	XMLName     xml.Name `xml:"IpInUse"`
	PoolIndex   int      `xml:"PoolIndex"`
	VLANID      int      `xml:"VLANID"`
	HType       int      `xml:"HType"`
	Type        int      `xml:"Type"`
	IfIndex     int      `xml:"IfIndex"`
	Ipv4Address string   `xml:"Ipv4Address"`
	CID         string   `xml:"CID"`
	HAddress    string   `xml:"HAddress"`
	EndLease    string   `xml:"EndLease"`
}

type Ifmgr struct {
	/* top level
	   Ifmgr
	     EthInterfaces
	       []EthInterface
	     Interfaces
	       []Interface
	     PortIsolation
	       Capabilities
	       Interfaces
	         []Interface
	     Ports                                     ***ReadOnly***
	       []Port                                  ***ReadOnly***
	     StormConstrain
	       Interfaces
	         []Interface
	       Interval
	     Statistics                                ***ReadOnly***
	       []InterfaceStatistics                   ***ReadOnly***
	     TrafficStatistics                         ***ReadOnly***
	       InterfacesTrafficStatistics             ***ReadOnly***
	         []InterfaceTrafficStatistics          ***ReadOnly***
	       Interval                                ***ReadOnly***
	     EthPortStatistics                         ***ReadOnly***
	       []InterfacesEthPortStatistics           ***ReadOnly***
	*/
	EthInterfaces     *EthInterfaces     `xml:"EthInterfaces"`
	Interfaces        *Interfaces        `xml:"Interfaces"`
	Ports             *Ports             `xml:"Ports"`
	Statistics        *Statistics        `xml:"Statistics"`
	TrafficStatistics *TrafficStatistics `xml:"TrafficStatistics"`
	EthPortStatistics *EthPortStatistics `xml:"EthPortStatistics"`
}

// EthInterfaces table contains Ethernet interfaces information.
type EthInterfaces struct {
	Interfaces []EthInterface `xml:"Interface"`
}

type EthInterface struct {
	XMLName                   xml.Name                    `xml:"Interface"`
	IfIndex                   int                         `xml:"IfIndex"`
	FlowControl               int                         `xml:"FlowControl,omitempty"`
	Jumboframe                int                         `xml:"Jumboframe,omitempty"`
	AutoPowerDown             bool                        `xml:"AutoPowerDown,omitempty"`
	BPDUDrop                  bool                        `xml:"BPDUDrop,omitempty"`
	EEE                       bool                        `xml:"EEE,omitempty"`
	BroadcastSuppression      []BroadcastSuppression      `xml:"BroadcastSuppression,omitempty"`
	MulticastSuppression      []MulticastSuppression      `xml:"MulticastSuppression,omitempty"`
	UnknownUnicastSuppression []UnknownUnicastSuppression `xml:"UnknownUnicastSuppression,omitempty"`
}

type BroadcastSuppression struct {
	ConfigValue int `xml:"ConfigValue"`
	ActualValue int `xml:"ActualValue"`
	Unit        int `xml:"Unit"`
}

type UnknownUnicastSuppression struct {
	ConfigValue int `xml:"ConfigValue"`
	ActualValue int `xml:"ActualValue"`
	Unit        int `xml:"Unit"`
}

type MulticastSuppression struct {
	ConfigValue int `xml:"ConfigValue"`
	ActualValue int `xml:"ActualValue"`
	Unit        int `xml:"Unit"`
	Flag        int `xml:"Flag,omitempty"`
}

// Interfaces table contains basic interface information.
type Interfaces struct {
	Interfaces []Interface `xml:"Interface"`
}

type Interface struct {
	XMLName             xml.Name `xml:"Interface"`
	Name                string   `xml:"Name,omitempty"`
	AbbreviatedName     string   `xml:"AbbreviatedName,omitempty"`
	InetAddressIPV4     string   `xml:"InetAddressIPV4,omitempty"`
	InetAddressIPV4Mask string   `xml:"InetAddressIPV4Mask,omitempty"`
	MAC                 string   `xml:"MAC,omitempty"`
	IfIndex             int      `xml:"IfIndex"`
	PortIndex           int      `xml:"PortIndex,omitempty"`
	IfTypeExt           int      `xml:"ifTypeExt,omitempty"`
	//IfType - Interface type, according to IANAifType
	//  https://www.iana.org/assignments/ianaiftype-mib/ianaiftype-mib
	//Examples
	//  6-Physical ethernet Interface,
	//  161-Bridge-Aggregation,
	//  136-Vlan-interface
	IfType      int    `xml:"ifType,omitempty"`
	Description string `xml:"Description,omitempty"`
	//AdminStatus - Interface administration status
	// 1 - Admin Up  	2 - Admin Down
	AdminStatus int `xml:"AdminStatus,omitempty"`
	//OperStatus - Interface operation status
	// 1 - up
	// 2 - down
	// 3 - testing
	// 4 - unknown
	// 5 - dormant
	// 6 - notPresent
	// 7 - lowerLayerDown
	OperStatus int `xml:"OperStatus,omitempty"` //1-UP, 2-DOWN
	//Configured speed of an interface
	// 1 - auto					8 - 155Mbps
	// 2 - 10Mbps				16 - 622Mbps
	// 4 - 100M                 64 - 2Gbps
	// 32 - 1Gbps               128 - 2.5Gbps
	// 1024 - 10Gbps            256 - 4Gbps
	// 8192 - 40Gbps            512 - 8Gbps
	// 16384 - 100Gbps          2048 - 16Gbps
	//                          4096 - 20Gbps
	//                          32768 - 5Gbps
	//
	// Example: 37—Auto-negotiation 	mode, and negotiation values are 100Mbps and 1000Mbps.
	ConfigSpeed  int `xml:"ConfigSpeed,omitempty"`
	ActualSpeed  int `xml:"ActualSpeed,omitempty"`
	ConfigDuplex int `xml:"ConfigDuplex,omitempty"` //1-full, 2-half, 3-auto
	ActualDuplex int `xml:"ActualDuplex,omitempty"` //1-full, 2-half, 3-auto
	PortLayer    int `xml:"PortLayer,omitempty"`
	// LinkType - VLAN type of an interface:
	// 1 - Access,   2 - Trunk,   3 - Hybrid
	LinkType             int `xml:"LinkType,omitempty"`
	PVID                 int `xml:"PVID,omitempty"`
	PhysicalIndex        int `xml:"PhysicalIndex,omitempty"`
	ForwardingAttributes int `xml:"ForwardingAttributes,omitempty"`
	ConfigMTU            int `xml:"ConfigMTU,omitempty"`
	ActualMTU            int `xml:"ActualMTU,omitempty"`
	Loopback             int `xml:"Loopback,omitempty"`
	//MDI mode of an interface
	// 1 - MDI-II (straight-through cable)
	// 2 - MDI-X (crossover cable)
	// 3 - MDI-AUTO (auto-sensing)
	MDI               int  `xml:"MDI,omitempty"`
	ActualBandwidth   int  `xml:"ActualBandwidth,omitempty"`
	Interval          int  `xml:"Interval,omitempty"`          //absent in HP5130
	Actual64Bandwidth int  `xml:"Actual64Bandwidth,omitempty"` //absent in HP5130
	ForceUP           bool `xml:"ForceUP,omitempty"`
	SubPort           bool `xml:"SubPort,omitempty"`
}

type Ports struct {
	XMLName xml.Name `xml:"Ports"`
	Ports   []Port   `xml:"Port"`
}

type Port struct {
	XMLName   xml.Name `xml:"Port"`
	PortIndex int      `xml:"PortIndex,omitempty"`
	Name      string   `xml:"Name,omitempty"`
	IfIndex   int      `xml:"IfIndex,omitempty"`
}

type Statistics struct {
	Interfaces []InterfaceStatistics `xml:"Interface"`
}

//ReadOnly struct
type InterfaceStatistics struct {
	XMLName         xml.Name `xml:"Interface"`
	IfIndex         int      `xml:"IfIndex"`
	Name            string   `xml:"Name"`
	AbbreviatedName string   `xml:"AbbreviatedName"`
	InOctets        uint64   `xml:"InOctets"`
	InUcastPkts     uint64   `xml:"InUcastPkts"`
	InNUcastPkts    uint64   `xml:"InNUcastPkts"`
	InDiscards      uint64   `xml:"InDiscards"`
	InErrors        uint64   `xml:"InErrors"`
	InUnknownProtos uint64   `xml:"InUnknownProtos"`
	InRate          uint64   `xml:"InRate"`
	OutOctets       uint64   `xml:"OutOctets"`
	OutUcastPkts    uint64   `xml:"OutUcastPkts"`
	OutNUcastPkts   uint64   `xml:"OutNUcastPkts"`
	OutDiscards     uint64   `xml:"OutDiscards"`
	OutErrors       uint64   `xml:"OutErrors"`
	OutRate         uint64   `xml:"OutRate"`
	LastClear       string   `xml:"LastClear"`
}

type TrafficStatistics struct {
	TrafficStatistics *InterfacesTrafficStatistics `xml:"Interfaces"`
}

type InterfacesTrafficStatistics struct {
	Interfaces []InterfaceTrafficStatistics `xml:"Interface"`
}

//ReadOnly struct
type InterfaceTrafficStatistics struct {
	XMLName   xml.Name `xml:"Interface"`
	IfIndex   int      `xml:"IfIndex"`
	Name      string   `xml:"Name"`
	Interval  int      `xml:"Interval"`
	InPkts    uint64   `xml:"InPkts"`
	OutPkts   uint64   `xml:"OutPkts"`
	InOctets  uint64   `xml:"InOctets"`
	OutOctets uint64   `xml:"OutOctets"`
	InBits    uint64   `xml:"InBits"`
	OutBits   uint64   `xml:"OutBits"`
}

type EthPortStatistics struct {
	Interfaces []InterfaceEthPortStatistics `xml:"Interface"`
}

//ReadOnly struct
type InterfaceEthPortStatistics struct {
	IfIndex         int    `xml:"IfIndex"`
	Name            string `xml:"Name"`
	InBytes         uint64 `xml:"InBytes"`
	InPkts          uint64 `xml:"InPkts"`
	InUcastPkts     uint64 `xml:"InUcastPkts"`
	InBrdcastPkts   uint64 `xml:"InBrdcastPkts"`
	InMulticastPkts uint64 `xml:"InMulticastPkts"`
	//InPauses - Number of inbound pause frames on an interface.
	InPauses               uint64 `xml:"InPauses"`
	InNormalUnicastBytes   uint64 `xml:"InNormalUnicastBytes"`
	InNormalBrdcastBytes   uint64 `xml:"InNormalBrdcastBytes"`
	InNormalMulticastBytes uint64 `xml:"InNormalMulticastBytes"`
	InUnknownUnicastBytes  uint64 `xml:"InUnknownUnicastBytes"`
	InNormalPkts           uint64 `xml:"InNormalPkts"`
	InNormalUnicastPkts    uint64 `xml:"InNormalUnicastPkts"`
	InNormalBrdcastPkts    uint64 `xml:"InNormalBrdcastPkts"`
	InNormalMulticastPkts  uint64 `xml:"InNormalMulticastPkts"`
	InUnknownUnicastPkts   uint64 `xml:"InUnknownUnicastPkts"`
	InNormalPauses         uint64 `xml:"InNormalPauses"`
	InErrorPkts            uint64 `xml:"InErrorPkts"`
	//InPktSpeed - Rate of inbound packages on an interface.
	InPktSpeed  uint64 `xml:"InPktSpeed"`
	InByteSpeed uint64 `xml:"InByteSpeed"`
	InRunts     uint64 `xml:"InRunts"`
	InGiants    uint64 `xml:"InGiants"`
	//InThrottles  - Number of inbound frames that had a non-integer number of bytes
	InThrottles uint64 `xml:"InThrottles"`
	//InErrCRCFrames - Total number of inbound frames that had a normal length, but contained CRC errors.
	InErrCRCFrames uint64 `xml:"InErrCRCFrames"`
	//InErrFrames - Total number of inbound frames that contained CRC errors and a non-integer number of bytes.
	InErrFrames uint64 `xml:"InErrFrames"`
	//InAbortPkts - Total number of illegal inbound packets.
	InAbortPkts uint64 `xml:"InAbortPkts"`
	//InSpeedPeakBytes - Peak rate of inbound traffic in Bps.
	InSpeedPeakBytes uint64 `xml:"InSpeedPeakBytes"`
	//InSpeedPeakTime - The time when the peak inbound traffic rate occurred.
	InSpeedPeakTime  string `xml:"InSpeedPeakTime"`
	OutBytes         uint64 `xml:"OutBytes"`
	OutPkts          uint64 `xml:"OutPkts"`
	OutUcastPkts     uint64 `xml:"OutUcastPkts"`
	OutBrdcastPkts   uint64 `xml:"OutBrdcastPkts"`
	OutMulticastPkts uint64 `xml:"OutMulticastPkts"`
	//OutPauses - Number of outbound pause frames on an interface.
	OutPauses              uint64 `xml:"OutPauses"`
	OutNormalUnicastBytes  uint64 `xml:"OutNormalUnicastBytes"`
	OutNormalPkts          uint64 `xml:"OutNormalPkts"`
	OutNormalUnicastPkts   uint64 `xml:"OutNormalUnicastPkts"`
	OutNormalBrdcastPkts   uint64 `xml:"OutNormalBrdcastPkts"`
	OutNormalMulticastPkts uint64 `xml:"OutNormalMulticastPkts"`
	OutUnknownUnicastPkts  uint64 `xml:"OutUnknownUnicastPkts"`
	OutNormalPauses        uint64 `xml:"OutNormalPauses"`
	OutErrorPkts           uint64 `xml:"OutErrorPkts"`
	//OutPktSpeed - Rate of outbound packages on an interface.
	OutPktSpeed uint64 `xml:"OutPktSpeed"`
	//OutByteSpeed - Rate of outbound bytes on an interface.
	OutByteSpeed uint64 `xml:"OutByteSpeed"`
	//OutAbortPkts - Number of packets that failed to be transmitted, for example, because of Ethernet collisions.
	OutAbortPkts uint64 `xml:"OutAbortPkts"`
	//OutDeferedFrames - Number of frames that the interface deferred to transmit because of detected collisions.
	OutDeferedFrames uint64 `xml:"OutDeferedFrames"`
	//OutCollisionFrames - Number of frames that the interface stopped transmitting because Ethernet collisions were detected during transmission.
	OutCollisionFrames uint64 `xml:"OutCollisionFrames"`
	//OutLateCollisionFrames - Number of frames that the interface deferred to transmit after transmitting their first 512 bits because of detected collisions.
	OutLateCollisionFrames uint64 `xml:"OutLateCollisionFrames"`
	//OutLostCarriers - Number of carrier losses during transmission.
	OutLostCarriers   uint64 `xml:"OutLostCarriers"`
	OutSpeedPeakBytes uint64 `xml:"OutSpeedPeakBytes"`
	OutSpeedPeakTime  string `xml:"OutSpeedPeakTime"`
}

type IPCIM struct {
	/* top level
	   IPCIM
	     IpSourceBindingInterface
	       []SourceBinding
	     IpVerifySource
	       []VerifySource
	*/
	IPSourceBindingInterface *IPSourceBindingInterface `xml:"IpSourceBindingInterface"`
	IPVerifySource           *IPVerifySource           `xml:"IpVerifySource"`
}

// IPSourceBindingInterface table contains Ip Source Binding table information.
type IPSourceBindingInterface struct {
	SourceBindings []SourceBinding `xml:"SourceBinding"`
}

type SourceBinding struct {
	XMLName     xml.Name `xml:"SourceBinding"`
	IfIndex     string   `xml:"IfIndex"`
	Ipv4Address string   `xml:"Ipv4Address"`
	MacAddress  string   `xml:"MacAddress"`
	VLANID      string   `xml:"VLANID,omitempty"`
}

// IPVerifySource table contains Ip Verify Source table information.
type IPVerifySource struct {
	VerifySourceInterfaces []VerifySource `xml:"VerifySource"`
}

type VerifySource struct {
	XMLName          xml.Name `xml:"VerifySource"`
	IfIndex          int      `xml:"IfIndex"`
	VerifyIPAddress  bool     `xml:"VerifyIpAddress,omitempty"`
	VerifyMacAddress bool     `xml:"VerifyMacAddress,omitempty"`
}

type MAC struct {
	/*top level
	MAC
	  MacAging
	  MacFwdSrcCheck
	    []FwdSrcCheck
	  MacUnicastTable
	    []MacTableEntry
	  MacPort
	    []PortLearn
	  MacAging
	  MacSpecification

	*/
	MacUnicastTable  *MacUnicastTable  `xml:"MacUnicastTable"`
	MacPort          *MacPort          `xml:"MacPort"`
	MacVLAN          *MacVLAN          `xml:"MacVLAN"`
	MacAging         *MacAging         `xml:"MacAging"`
	MacSpecification *MacSpecification `xml:"MacSpecification"`
}

// MacUnicastTable table contains unicast MAC address table information.
type MacUnicastTable struct {
	Unicast []MacTableEntry `xml:"Unicast"`
}

type MacTableEntry struct {
	XMLName    xml.Name `xml:"Unicast"`
	VLANID     string   `xml:"VLANID"`
	MacAddress string   `xml:"MacAddress"`
	PortIndex  int      `xml:"PortIndex"`
	Status     string   `xml:"Status"`
	Aging      string   `xml:"Aging"`
}

// MacPort table contains the information of MAC learning on an interface.
type MacPort struct {
	XMLName    xml.Name    `xml:"MacPort"`
	PortsLearn []PortLearn `xml:"PortLearn"`
}

type PortLearn struct {
	XMLName              xml.Name `xml:"PortLearn"`
	PortIndex            int      `xml:"PortIndex"`
	LearnEnable          bool     `xml:"LearnEnable,omitempty"`
	PortForwardingEnable bool     `xml:"PortForwardingEnable,omitempty"`
}

// MacVLAN table contains the information of MAC learning for a VLAN.
type MacVLAN struct {
	XMLName    xml.Name    `xml:"MacVLAN"`
	VLANsLearn []VLANLearn `xml:"VLANLearn"`
}

type VLANLearn struct {
	XMLName     xml.Name `xml:"VLANLearn"`
	VLANID      int      `xml:"VLANID"`
	LearnEnable bool     `xml:"LearnEnable"`
}

type MacAging struct {
	XMLName      xml.Name `xml:"MacAging"`
	AgingTimeMin int      `xml:"AgingTimeMin,omitempty"`
	AgingTimeMax int      `xml:"AgingTimeMax,omitempty"`
	AgingTime    int      `xml:"AgingTime,omitempty"`
}

type MacSpecification struct {
	XMLName                  xml.Name `xml:"MacSpecification"`
	PortLearnMaxNumLimit     int      `xml:"PortLearnMaxNumLimit,omitempty"`
	SupportMacGroup          bool     `xml:"SupportMacGroup,omitempty"`
	SupportMacVLANLearnLimit bool     `xml:"SupportMacVLANLearnLimit,omitempty"`
	SupportPortBridgeEnable  bool     `xml:"SupportPortBridgeEnable,omitempty"`
}

// Port mirroring configuration
type MGROUP struct {
	/* top level
	   MGROUP
		 Capabilities            ***ReadOnly***
	     Groups
	       []Group
	     MonitorPort
	       []Group
		 ReflectorPort
	       []Group
	     EgressPort
	       []Group
		 ProbeVlan
	       []Group
	     SourcePorts
	       []SourcePort
	*/
	Capabilities  *PortMirroringCapabilities   `xml:"Capabilities"`
	Groups        *MirrorGroups                `xml:"Groups"`
	MonitorPort   *PortMirroringMonitorPorts   `xml:"MonitorPort"`
	ReflectorPort *PortMirroringReflectorPorts `xml:"ReflectorPort"`
	EgressPort    *PortMirroringEgressPorts    `xml:"EgressPort"`
	ProbeVlan     *PortMirroringProbeVlans     `xml:"ProbeVlan"`
	SourcePorts   *PortMirroringSourcePorts    `xml:"SourcePorts"`
}

// MirrorGroups table contains information about mirroring groups.
type MirrorGroups struct {
	XMLName      xml.Name      `xml:"Groups"`
	MirrorGroups []MirrorGroup `xml:"Group"`
}

type MirrorGroup struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`             // Valid values are: 1-256
	Type    int      `xml:"Type,omitempty"` // 1 - Local, 2 - Remote-source, 3 - Remote-destination
	Status  int      `xml:"Status,omitempty"`
}

// PortMirroringMonitorPorts table contains information about monitor port of mirroring groups.
type PortMirroringMonitorPorts struct {
	XMLName      xml.Name                   `xml:"MonitorPort"`
	MonitorPorts []PortMirroringMonitorPort `xml:"Group"`
}

type PortMirroringMonitorPort struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`
	Port    int      `xml:"Port"`
}

// PortMirroringReflectorPorts table contains information about reflector port of mirroring groups
type PortMirroringReflectorPorts struct {
	XMLName        xml.Name                     `xml:"ReflectorPort"`
	ReflectorPorts []PortMirroringReflectorPort `xml:"Group"`
}

type PortMirroringReflectorPort struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`
	Port    int      `xml:"Port"`
}

// PortMirroringEgressPorts table contains information about egress port of mirroring groups.
type PortMirroringEgressPorts struct {
	XMLName     xml.Name                  `xml:"EgressPort"`
	EgressPorts []PortMirroringEgressPort `xml:"Group"`
}

type PortMirroringEgressPort struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`   // Mirroring group ID. Value range: 1 to 256
	Port    int      `xml:"Port"` // Index of egress	port.
}

// PortMirroringProbeVlans table contains information about remote probe VLAN of mirroring groups.
type PortMirroringProbeVlans struct {
	XMLName    xml.Name                 `xml:"ProbeVlan"`
	ProbeVlans []PortMirroringProbeVlan `xml:"Group"`
}

// PortMirroringProbeVlan Dedicated VLAN that sends packets from the source device to the destination device.
type PortMirroringProbeVlan struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`     // Mirroring group ID. Value range: 1 to 256
	VlanID  int      `xml:"VlanID"` // Probe VLAN ID
}

// PortMirroringSourcePorts table contains information about source ports of the mirroring group.
type PortMirroringSourcePorts struct {
	XMLName     xml.Name                  `xml:"SourcePorts"`
	SourcePorts []PortMirroringSourcePort `xml:"SourcePort"`
}

type PortMirroringSourcePort struct {
	XMLName   xml.Name `xml:"SourcePort"`
	ID        int      `xml:"ID"`                  // Mirroring group ID. Value range: 1 to 256
	IfIndex   int      `xml:"IfIndex,omitempty"`   // Interface index
	Direction int      `xml:"Direction,omitempty"` // Direction of source port: 1—Inbound, 2—Outbound, 3—Both.
}

type PortMirroringCapabilities struct {
	XMLName                xml.Name `xml:"Capabilities"`
	MaxGroupNum            int      `xml:"MaxGroupNum"`
	SourceType             int      `xml:"SourceType"`
	VlanTagMode            int      `xml:"VlanTagMode"`
	Sampler                int      `xml:"Sampler"`
	MultiMonitor           bool     `xml:"MultiMonitor"`
	MultiMonitorInOneGroup bool     `xml:"MultiMonitorInOneGroup"`
}

type ResourceMonitor struct {
	/* top level
	   ResourceMonitor
		 Configuration
	     Monitors				***ReadOnly***
		   []Monitor			***ReadOnly***
		 Resources
		   []Resource
	*/
	Configuration *ResourceMonitorConfiguration `xml:"Configuration"`
	Monitors      *Monitors                     `xml:"Monitors"`
	Resources     *Resources                    `xml:"Resources"`
}

type ResourceMonitorConfiguration struct {
	XMLName                xml.Name `xml:"Configuration"`
	MinorResendEnable      bool     `xml:"MinorResendEnable,omitempty"`
	SyslogOutput           bool     `xml:"SyslogOutput,omitempty"`           //Output enable for syslog.
	SnmpNotificationOutput bool     `xml:"SnmpNotificationOutput,omitempty"` //Output enable for SNMP notification.
	NetconfEventOutput     bool     `xml:"NetconfEventOutput,omitempty"`     //Output enable for NETCONF event.
}

// Monitors table contains usage information about the resource monitor.
type Monitors struct {
	XMLName  xml.Name  `xml:"Monitors"`
	Monitors []Monitor `xml:"Monitor"`
}

type Monitor struct {
	XMLName    xml.Name    `xml:"Monitor"`
	DeviceNode *DeviceNode `xml:"DeviceNode"`
	Name       string      `xml:"Name"`
	Unit       string      `xml:"Unit"` // absolute | percentage
	Used       int         `xml:"Used"`
	Free       int         `xml:"Free"`
	Total      int         `xml:"Total"`
}

type DeviceNode struct {
	XMLName xml.Name `xml:"DeviceNode"`
	Chassis int      `xml:"Chassis"`
	Slot    int      `xml:"Slot"`
	CPUID   int      `xml:"CPUID"`
}

// Resources table contains thresholds information about the resource monitor.
type Resources struct {
	XMLName   xml.Name   `xml:"Resources"`
	Resources []Resource `xml:"Resource"`
}

type Resource struct {
	XMLName         xml.Name    `xml:"Resource"`
	DeviceNode      *DeviceNode `xml:"DeviceNode"`
	Name            string      `xml:"Name,omitempty"`
	Unit            string      `xml:"Unit,omitempty"`            //Unit of resource monitor threshold. Valid values are: absolute, percentage
	MinorThreshold  int         `xml:"MinorThreshold,omitempty"`  //The free resource threshold for minor condition.
	SevereThreshold int         `xml:"SevereThreshold,omitempty"` //The free resource threshold for severe condition.
}

type VLAN struct {
	/* top level
	   VLAN
	     Interfaces
	       []Interface
	     AccessInterfaces
	       []AccessInterface
	     TrunkInterfaces
	       []TrunkInterface
	     HybridInterfaces
	       []HybridInterface
	     VLANs
	       VLANID
	     VoicePorts
	       []VoicePort
	*/
	Interfaces       *VLANInterfaces   `xml:"Interfaces"`
	AccessInterfaces *AccessInterfaces `xml:"AccessInterfaces"`
	TrunkInterfaces  *TrunkInterfaces  `xml:"TrunkInterfaces"`
	HybridInterfaces *HybridInterfaces `xml:"HybridInterfaces"`
	VLANs            *VLANs            `xml:"VLANs"`
}

// VLANInterfaces table contains VLAN information for a port.
type VLANInterfaces struct {
	XMLName    xml.Name        `xml:"Interfaces"`
	Interfaces []VLANInterface `xml:"Interface"`
}

// AccessInterfaces table contains information about Access ports.
type AccessInterfaces struct {
	XMLName          xml.Name          `xml:"AccessInterfaces"`
	AccessInterfaces []AccessInterface `xml:"Interface"`
}

// TrunkInterfaces table contains information about Trunk ports.
type TrunkInterfaces struct {
	XMLName         xml.Name         `xml:"TrunkInterfaces"`
	TrunkInterfaces []TrunkInterface `xml:"Interface"`
}

// HybridInterfaces table contains information about Hybrid ports.
type HybridInterfaces struct {
	XMLName          xml.Name          `xml:"HybridInterfaces"`
	HybridInterfaces []HybridInterface `xml:"Interface"`
}

type VLANInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfIndex int      `xml:"IfIndex"`
	// LinkType:
	// 1 - Access, 2 - Trunk, 3 - Hybrid
	LinkType int `xml:"LinkType,omitempty"`
	// PVID value range: 1 to 4094.
	PVID int `xml:"PVID,omitempty"`
	// Name - full name of the interface, including the interface type	and number.
	// String. Length: up to 47	characters.
	Name string `xml:"Name,omitempty"`
	// UntaggedVlanList is a comma-separated list of VLAN items.
	// An item can be an individual VLAN ID or a VLAN ID range.
	// Use a hyphen (-) to separate the start and end VLANs of a range.
	// The string cannot end with a comma or hyphen, or include any whitespace characters.
	//
	// Example:
	// "1,2,3,5-8,10-20"
	UntaggedVlanList string `xml:"UntaggedVlanList"`
	// TaggedVlanList is a comma-separated list of VLAN items.
	// An item can be an individual VLAN ID or a VLAN ID range.
	// Use a hyphen (-) to separate the start and end VLANs of a range.
	// The string cannot end with a comma or hyphen, or include any whitespace characters.
	// The column is available only for	trunk and hybrid ports.
	TaggedVlanList string `xml:"TaggedVlanList"`
	// PermitVlanList is a comma-separated list of VLAN items.
	// An item can be an individual VLAN ID or a VLAN ID range.
	// Use a hyphen (-) to separate the start and end VLANs of a range.
	// The string cannot end with a comma or hyphen, or include any whitespace characters.
	// The column is available only for	trunk ports.
	PermitVlanList string `xml:"PermitVlanList"`
}

type AccessInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfIndex int      `xml:"IfIndex"`
	PVID    int      `xml:"PVID"`
}

type TrunkInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfIndex int      `xml:"IfIndex"`
	// PermitVlanList is a comma-separated list of VLAN items.
	// An item can be an individual VLAN ID or a VLAN ID range.
	// Use a hyphen (-) to separate the start and end VLANs of a range.
	// The string cannot end with a comma or hyphen, or include any whitespace characters.
	//
	// Examples: "1,300", "300-302", "1,301-302", "1-4094" (all)
	PermitVlanList string `xml:"PermitVlanList"`
	// PVID value range: 1 to 4094.
	PVID int `xml:"PVID,omitempty"`
}

type HybridInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfIndex int      `xml:"IfIndex"`
	// UntaggedVlanList is a comma-separated list of VLAN items.
	// An item can be an individual VLAN ID or a VLAN ID range.
	// Use a hyphen (-) to separate the start and end VLANs of a range.
	// The string cannot end with a comma or hyphen, or include any whitespace characters.
	UntaggedVlanList string `xml:"UntaggedVlanList"`
	// TaggedVlanList is a comma-separated list of VLAN items.
	// An item can be an individual VLAN ID or a VLAN ID range.
	// Use a hyphen (-) to separate the start and end VLANs of a range.
	// The string cannot end with a comma or hyphen, or include any whitespace characters.
	// The column is available only for	trunk and hybrid ports.
	TaggedVlanList string `xml:"TaggedVlanList"`
	// PVID value range: 1 to 4094.
	PVID int `xml:"PVID"`
}

// VLANs table contains basic VLAN information.
type VLANs struct {
	XMLName xml.Name `xml:"VLANs"`
	VLANs   []VLANID `xml:"VLANs"`
}

type VLANID struct {
	XMLName          xml.Name `xml:"VLANID"`
	ID               int      `xml:"ID"`
	Name             string   `xml:"Name,omitempty"`
	Description      string   `xml:"Description,omitempty"`
	RouteIfIndex     string   `xml:"RouteIfIndex,omitempty"`
	UntaggedPortList string   `xml:"UntaggedPortList,omitempty"`
	TaggedPortList   string   `xml:"TaggedPortList,omitempty"`
	AccessPortList   string   `xml:"AccessPortList,omitempty"`
	Ipv4             *Ipv4    `xml:"Ipv4"`
	Shared           bool     `xml:"Shared,omitempty"`
}

// IPv4 address of the VLAN interface
type Ipv4 struct {
	XMLName     xml.Name `xml:"Ipv4"`
	Ipv4Address string   `xml:"Ipv4Address"`
	Ipv4Mask    string   `xml:"Ipv4Mask"`
}

type STP struct {
	/* top level
	   STP
	     Base
	     Region
	     Interfaces
	       []Interface
	*/
	Base       *STPBase                    `xml:"Base"`
	Interfaces *STPInterfacesConfiguration `xml:"Interfaces"`
}

type STPBase struct {
	XMLName xml.Name `xml:"Base"`
	//The spanning-tree working modes:
	//0 - STP; 2 - RSTP; 3 - MSTP; 4 - PVST;
	Mode           int  `xml:"Mode,omitempty"`
	TcThreshold    int  `xml:"TcThreshold,omitempty"`    //values 1-255
	PathCostMethod int  `xml:"PathCostMethod,omitempty"` // 0 - Leagcy; 1 - IEEE 802.1D-1998; 2 - IEEE 802.1t;
	HelloTime      int  `xml:"HelloTime,omitempty"`      //The intervals in seconds. Valid values are: 1-10.
	MaxHops        int  `xml:"MaxHops,omitempty"`        //Valid values are: 1-40.
	MaxAge         int  `xml:"MaxAge,omitempty"`         //Valid values are: 6-40.
	ForwardDelay   int  `xml:"ForwardDelay,omitempty"`   //The forward delay timer in seconds. Valid values are: 4-30.
	TcSnooping     bool `xml:"TcSnooping,omitempty"`
	DigestSnooping bool `xml:"DigestSnooping"`
	BPDUProtect    bool `xml:"BPDUProtect,omitempty"`
	TcProtect      bool `xml:"TcProtect,omitempty"`
	Enable         bool `xml:"Enable,omitempty"`
}

// STPInterfacesConfiguration table contains information about interface-level STP functions.
type STPInterfacesConfiguration struct {
	XMLName                    xml.Name                    `xml:"Interfaces"`
	STPInterfacesConfiguration []STPInterfaceConfiguration `xml:"Interface"`
}

type STPInterfaceConfiguration struct {
	XMLName           xml.Name `xml:"Interface"`
	IfIndex           int      `xml:"IfIndex"`
	PointToPoint      int      `xml:"PointToPoint,omitempty"`      //1 - Force-true; 2 - Force-false; 3 - Auto-negotiated by the link;
	TransmitHoldCount int      `xml:"TransmitHoldCount,omitempty"` //Valid values are: 1-255.
	Enable            bool     `xml:"Enable,omitempty"`
	EdgedPort         bool     `xml:"EdgedPort,omitempty"`
	RootProtect       bool     `xml:"RootProtect,omitempty"`
	LoopProtect       bool     `xml:"LoopProtect,omitempty"`
	RoleRestrict      bool     `xml:"RoleRestrict,omitempty"`
	TcRestrict        bool     `xml:"TcRestrict,omitempty"`
	DigestSnooping    bool     `xml:"DigestSnooping,omitempty"`
}

type Syslog struct {
	/* top level
	   Syslog
	     Configuration
	     LogBuffer
	     LogHosts
	       []Host
	     Logs                  ***ReadOnly***
	       []Log               ***ReadOnly***
	     OutputRules
	       []OutputRule
	*/
	Configuration *SyslogConfiguration `xml:"Configuration"`
	LogBuffer     *LogBuffer           `xml:"LogBuffer"`
	LogHosts      *LogHosts            `xml:"LogHosts"`
	Logs          *Logs                `xml:"Logs"`
}

// SyslogConfiguration contains information about syslog (the information center).
type SyslogConfiguration struct {
	XMLName xml.Name `xml:"Configuration"`
	// Status of syslog (enable or disable)
	State string `xml:"State,omitempty"`
	// Status of the duplicate log suppression feature
	DuplicateLogSuppression string `xml:"DuplicateLogSuppression,omitempty"`
}

type LogBuffer struct {
	XMLName xml.Name `xml:"LogBuffer"`
	State   string   `xml:"State,omitempty"`
	// BufferSize - Maximum log buffer size configured by the user.
	BufferSize int `xml:"BufferSize,omitempty"`
	// BufferSizeLimit - Maximum log buffer size supported by the device.
	BufferSizeLimit int `xml:"BufferSizeLimit,omitempty"` //***ReadOnly***
	// LogsCount - Number of logs stored in the log buffer.
	LogsCount int `xml:"LogsCount,omitempty"` //***ReadOnly***
	// DroppedLogsCount - Number of dropped logs.
	DroppedLogsCount int `xml:"DroppedLogsCount,omitempty"` //***ReadOnly***
	// OverwrittenLogsCount - Number of overwritten logs.
	OverwrittenLogsCount int                   `xml:"OverwrittenLogsCount,omitempty"` //***ReadOnly***
	LogsCountPerSeverity *LogsCountPerSeverity `xml:"LogsCountPerSeverity"`           //***ReadOnly***
}

//ReadOnly struct
type LogsCountPerSeverity struct {
	XMLName       xml.Name `xml:"LogsCountPerSeverity"`
	Emergency     int      `xml:"Emergency"`
	Alert         int      `xml:"Alert"`
	Critical      int      `xml:"Critical"`
	Error         int      `xml:"Error"`
	Warning       int      `xml:"Warning"`
	Notice        int      `xml:"Notice"`
	Informational int      `xml:"Informational"`
	Debug         int      `xml:"Debug"`
}

// LogHosts table contains log hosts information.
type LogHosts struct {
	XMLName xml.Name  `xml:"LogHosts"`
	Hosts   []LogHost `xml:"Hosts"`
}

type LogHost struct {
	XMLName xml.Name `xml:"Host"`
	Address string   `xml:"Address,omitempty"`
	VRF     string   `xml:"VRF,omitempty"`
	Port    int      `xml:"Port,omitempty"`
	//Logging facility used by the log host, as follows:
	//128 - local0, 136 - local1, 144 - local2,  152 - local3, 160 - local4, 168 - local5, 176 - local6, 184 - local7.
	Facility int `xml:"Facility,omitempty"`
}

// Logs table contains information about the logs in the log buffer (syslog messages).
// ReadOnly struct
type Logs struct {
	XMLName xml.Name `xml:"Logs"`
	Logs    []Log    `xml:"Log"`
}

//ReadOnly struct
type Log struct {
	XMLName xml.Name `xml:"Log"`
	// Index - number of the log as table index.
	Index int `xml:"Index"`
	// Time when the log was generated.
	Time string `xml:"Time"`
	// Group - module that generated the log. String length constraints must be in range(1..8).
	Group string `xml:"Group"`
	// Digest - Brief description of the log, String length constraints must be in range(1..8).
	Digest string `xml:"Digest"`
	// Severity level of the log:
	// 0(Emergency); 1(Alert); 2(Critical); 3(Error); 4(Warning); 5(Notification); 6(Informational); 7(Debugging)
	Severity int `xml:"Severity"`
	// Content of the log. String length constraints must be in range(0..1023).
	Content string `xml:"Content"`
}
