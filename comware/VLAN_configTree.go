package comware

import "encoding/xml"

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

// VLANInterface - read-only struct
type VLANInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfIndex int      `xml:"IfIndex"`
	// LinkType:
	// 1 - Access, 2 - Trunk, 3 - Hybrid
	LinkType InterfaceLinkType `xml:"LinkType"`
	// PVID value range: 1 to 4094.
	PVID int `xml:"PVID"`
	// Name - full name of the interface, including the interface type	and number.
	// String. Length: up to 47	characters.
	Name string `xml:"Name"`
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
	VLANs   []VLANID `xml:"VLANID"`
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

// InterfaceLinkType
//
//  1 - Access,
//  2 - Trunk,
//  3 - Hybrid
type InterfaceLinkType int

func (linkType InterfaceLinkType) String() string {
	switch linkType {
	case InterfaceLinkTypeAccess:
		return InterfaceLinkTypeAccessString
	case InterfaceLinkTypeTrunk:
		return InterfaceLinkTypeTrunkString
	case InterfaceLinkTypeHybrid:
		return InterfaceLinkTypeHybridString
	}

	return UnknownString
}
