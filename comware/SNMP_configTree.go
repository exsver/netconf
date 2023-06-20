package comware

import "encoding/xml"

type SNMP struct {
	/* top level
	   SNMP
	     Agent
	     Capabilities
	     Communities
	       []Community
	     MIBViews
	       []MIBView
	     Statistics
	     System
	*/
	Agent        *SNMPAgent        `xml:"Agent"`
	Capabilities *SNMPCapabilities `xml:"Capabilities"`
	Communities  *SNMPCommunities  `xml:"Communities"`
	MIBViews     *SNMPMIBViews     `xml:"MIBViews"`
	Statistics   *SNMPStatistics   `xml:"Statistics"`
	System       *SNMPSystem       `xml:"System"`
}

type SNMPAgent struct {
	XMLName       xml.Name `xml:"Agent"`
	LocalEngineId string   `xml:"LocalEngineId"`
	AgentUdpPort  int      `xml:"AgentUdpPort"`
	// PacketMaxSize - Maximum size of packets that the agent can receive or send in bytes.
	// Min value: 484, max value: 17940
	PacketMaxSize int `xml:"PacketMaxSize"`
	// GetLog - Logging status of get
	GetLog bool `xml:"GetLog"`
	// SetLog - Logging status of set
	SetLog bool `xml:"SetLog"`
	// NotificationLog - Logging status of notification
	NotificationLog bool `xml:"NotificationLog"`
}

type SNMPCapabilities struct {
	XMLName                 xml.Name `xml:"Capabilities"`
	MaxCommunities          int      `xml:"MaxCommunities"`
	MaxContexts             int      `xml:"MaxContexts"`
	MaxMaxCommunityMappings int      `xml:"MaxCommunityMappings"`
	MaxMIBViews             int      `xml:"MaxMIBViews"`
	MaxGroups               int      `xml:"MaxGroups"`
	MaxUsers                int      `xml:"MaxUsers"`
	MaxRemoteManages        int      `xml:"MaxRemoteManages"`
	MaxTargetHosts          int      `xml:"MaxTargetHosts"`
	MaxV3Roles              int      `xml:"MaxV3Roles"`
}

type SNMPCommunities struct {
	Communities []SNMPCommunity `xml:"Community"`
}

type SNMPCommunity struct {
	XMLName xml.Name `xml:"Community"`
	Name    string   `xml:"Name"`
	// Community type as follows:
	//  0 - read-only
	//  1 - read-write
	Type            int                           `xml:"Type"`
	MIBView         string                        `xml:"MIBView,omitempty"`
	Context         string                        `xml:"Context,omitempty"`
	IPv4BasicACL    *SNMPCommunityIPv4BasicACL    `xml:"IPv4BasicACL,omitempty"`
	IPv4AdvancedACL *SNMPCommunityIPv4AdvancedACL `xml:"IPv4AdvancedACL,omitempty"`
	IPv6BasicACL    *SNMPCommunityIPv6BasicACL    `xml:"IPv6BasicACL,omitempty"`
	IPv6AdvancedACL *SNMPCommunityIPv6AdvancedACL `xml:"IPv6AdvancedACL,omitempty"`
}

type SNMPCommunityIPv4BasicACL struct {
	XMLName xml.Name `xml:"IPv4BasicACL"`
	Number  int      `xml:"Number"`
}

type SNMPCommunityIPv4AdvancedACL struct {
	XMLName xml.Name `xml:"IPv4AdvancedACL"`
	Number  int      `xml:"Number"`
}

type SNMPCommunityIPv6BasicACL struct {
	XMLName xml.Name `xml:"IPv6BasicACL"`
	Number  int      `xml:"Number"`
}

type SNMPCommunityIPv6AdvancedACL struct {
	XMLName xml.Name `xml:"IPv6AdvancedACL"`
	Number  int      `xml:"Number"`
}

type SNMPMIBViews struct {
	MIBView []SNMPMIBView `xml:"View"`
}

type SNMPMIBView struct {
	XMLName xml.Name `xml:"View"`
	Name    string   `xml:"Name"`
	SubTree string   `xml:"SubTree"`
	Mask    string   `xml:"Mask"`
	// Type - The type of subtree:
	//  1 - Included
	//  2 - Excluded
	Type int `xml:"Type"`
}

type SNMPStatistics struct {
	InPkts  int `xml:"InPkts"`
	OutPkts int `xml:"OutPkts"`
	// InBadVersions - The total number of SNMP Messages which were delivered to the SNMP protocol entity and were for an unsupported SNMP version.
	InBadVersions int `xml:"InBadVersions"`
	// InBadCommunityNames - The total number of SNMP Messages delivered to the SNMP protocol entity which used a SNMP community name not known to said entity.
	InBadCommunityNames int `xml:"InBadCommunityNames"`
	// InBadCommunityUses - The total number of SNMP Messages delivered to the SNMP protocol entity which represented an SNMP operation which was not allowed by the SNMP community named in the Message.
	InBadCommunityUses int `xml:"InBadCommunityUses"`
	// InASNParseErrs - The total number of ASN.1 or BER errors encountered by the SNMP protocol entity when decoding received SNMP Messages.
	InASNParseErrs  int `xml:"InASNParseErrs"`
	InTooBigs       int `xml:"InTooBigs"`
	InNoSuchNames   int `xml:"InNoSuchNames"`
	InBadValues     int `xml:"InBadValues"`
	InReadOnlys     int `xml:"InReadOnlys"`
	InGenErrs       int `xml:"InGenErrs"`
	InTotalReqVars  int `xml:"InTotalReqVars"`
	InTotalSetVars  int `xml:"InTotalSetVars"`
	InGetRequests   int `xml:"InGetRequests"`
	InGetNexts      int `xml:"InGetNexts"`
	InSetRequests   int `xml:"InSetRequests"`
	InGetResponses  int `xml:"InGetResponses"`
	InTraps         int `xml:"InTraps"`
	OutTooBigs      int `xml:"OutTooBigs"`
	OutNoSuchNames  int `xml:"OutNoSuchNames"`
	OutBadValues    int `xml:"OutBadValues"`
	OutGenErrs      int `xml:"OutGenErrs"`
	OutGetRequests  int `xml:"OutGetRequests"`
	OutGetNexts     int `xml:"OutGetNexts"`
	OutSetRequests  int `xml:"OutSetRequests"`
	OutGetResponses int `xml:"OutGetResponses"`
	// OutTraps - The total number of SNMP Trap PDUs which have been generated by the SNMP protocol entity.
	OutTraps int `xml:"OutTraps"`
	// SilentDrops - The total number of Confirmed Class PDUs (such as GetRequest-PDUs, GetNextRequest-PDUs, GetBulkRequest-PDUs, SetRequest-PDUs, and InformRequest-PDUs)
	// delivered to the SNMP entity which were silently dropped because the size of a reply containing an alternate Response Class PDU (such as a Response-PDU)
	// with an empty variable-bindings field was greater than either a local constraint or the maximum message size associated with the originator of the request.
	SilentDrops int `xml:"SilentDrops"`
}

type SNMPSystem struct {
	XMLName xml.Name `xml:"System"`
	// AgentStatus - SNMP status as follows: enable, disable.
	AgentStatus string `xml:"AgentStatus"`
	// SNMP  version information.
	Version *SNMPVersion `xml:"Version"`
	// Location - SNMP physical location information.
	// String length constraints must be in range(0..255).
	Location string `xml:"Location,omitempty"`
	// Contact - SNMP contact information.
	// String length constraints must be in range(0..255).
	Contact string `xml:"Contact,omitempty"`
}

type SNMPVersion struct {
	XMLName xml.Name `xml:"Version"`
	V1      string   `xml:"V1"`
	V2C     string   `xml:"V2C"`
	V3      string   `xml:"V3"`
}
