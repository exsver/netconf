package comware

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strings"
)

type LAGG struct {
	/* top level
	LAGG
	  LAGGBase
	  LAGGGroups
	    []LAGGGroup
	  LAGGMembers
	    []LAGGMember
	*/
	Base        *LAGGBase    `xml:"Base"`
	LAGGGroups  *LAGGGroups  `xml:"LAGGGroups"`
	LAGGMembers *LAGGMembers `xml:"LAGGMembers"`
}

type LAGGBase struct {
	XMLName xml.Name `xml:"Base"`
	// SystemID - The Actor's System ID.
	SystemID string `xml:"SystemID,omitempty"`
	// SystemPriority - The priority assigned to the Actor System.
	SystemPriority int `xml:"SystemPriority,omitempty"`
	// LoadSharingMode - Global link-aggregation load sharing mode.
	LoadSharingMode int `xml:"LoadSharingMode"`
}

type LAGGGroups struct {
	Groups []LAGGGroup `xml:"LAGGGroup"`
}

type LAGGGroup struct {
	XMLName            xml.Name       `xml:"LAGGGroup"`
	GroupId            int            `xml:"GroupId"`
	LinkMode           LAGGLinkMode   `xml:"LinkMode,omitempty"`
	IfIndex            int            `xml:"IfIndex,omitempty"`
	MemberList         LAGGMemberList `xml:"MemberList,omitempty"`
	SelectedMemberList LAGGMemberList `xml:"SelectedMemberList,omitempty"`
	// LoadSharingMode - link-aggregation load sharing mode.
	//  2 - destination-mac,
	//  4 - source-mac,
	//  6 - destination-mac source-mac,
	//  8 - destination-ip,
	//  16 - source-ip,
	//  24 - destination-ip source-ip.
	LoadSharingMode       int    `xml:"LoadSharingMode,omitempty"`
	ManagementPort        int    `xml:"ManagementPort,omitempty"`
	PartnerSystemPriority int    `xml:"PartnerSystemPriority,omitempty"`
	PartnerSystemID       string `xml:"PartnerSystemID,omitempty"`
}

type LAGGMembers struct {
	Members []LAGGMember `xml:"LAGGMember"`
}

type LAGGMember struct {
	XMLName xml.Name `xml:"LAGGMember"`
	// IfIndex - Interface index of member port
	IfIndex int `xml:"IfIndex"`
	// GroupID - Aggregation group identifier
	GroupID int `xml:"GroupId"`
	// SelectedStatus - Selected status of a member port:
	//   1 - Selected
	//   2 - Unselected
	//   3 - Individual
	SelectedStatus LAGGMemberSelectedStatus `xml:"SelectedStatus,omitempty"`
	// UnSelectedReason - Unselected reason of a member port:
	//   0 - The port is attached to this aggregator.
	//       Indicate that selected status of a member port is Selected or Individual.
	//   1 - The current number of active ports has reached the upper limit.
	//   2 - All aggregation resources are already in-use.
	//   3 - The port's configuration is improper for being attached.
	//   4 - The port's partner is improper for being attached.
	//   5 - The number of current active ports has not reached the lower limit.
	//   6 - The port's physical state (down) is improper for being attached.
	//   7 - The port is not selected for an aggregator.
	//   8 - The port's hardware restriction is improper for being attached.
	//   9 - The port's speed is improper for being attached.
	//   10 - The port's duplex mode is improper for being attached.
	UnSelectedReason int `xml:"UnSelectedReason,omitempty"`
	// LacpMode - LACP mode of a member port:
	//   1 - Active
	//   2 - Passive
	LacpMode int `xml:"LacpMode,omitempty"`
	// ActorPortPriority - The priority assigned to this port by the Actor
	ActorPortPriority int `xml:"ActorPortPriority,omitempty"`
	// ActorState - The Actor's state variables for the port
	ActorState int `xml:"ActorState,omitempty"`
	// ActorOperKey - The operational Key value assigned to the port by the Actor
	ActorOperKey int `xml:"ActorOperKey,omitempty"`
	// ActorPort - The port number assigned to the port by the Actor
	ActorPort int `xml:"ActorPort,omitempty"`
	// PartnerSystemPriority - The priority assigned to the Partner System
	PartnerSystemPriority int `xml:"PartnerSystemPriority,omitempty"`
	// PartnerPortPriority - The priority assigned to this port by the Partner
	PartnerPortPriority int `xml:"PartnerPortPriority,omitempty"`
	// PartnerState - The Partner's state variables for the port
	PartnerState int `xml:"PartnerState,omitempty"`
	// PartnerOperKey - The operational Key value assigned to the port by the Partner
	PartnerOperKey int `xml:"PartnerOperKey,omitempty"`
	// PartnerPort - The port number assigned to the port by the Partner
	PartnerPort int `xml:"PartnerPort,omitempty"`
	// PartnerSystemID - The Partner's System ID
	PartnerSystemID       string `xml:"PartnerSystemID,omitempty"`
	ManagementPortEnable  *bool  `xml:"ManagementPortEnable,omitempty"`
	LacpEnable            *bool  `xml:"LacpEnable,omitempty"`
	LacpShortPeriodEnable *bool  `xml:"LacpShortPeriodEnable,omitempty"`
}

// LAGGLinkMode
//
//	1 - Static,
//	2 - Dynamic.
type LAGGLinkMode int

func (mode LAGGLinkMode) String() string {
	switch mode {
	case LAGGLinkModeStatic:
		return LAGGLinkModeStaticString
	case LAGGLinkModeDynamic:
		return LAGGLinkModeDynamicString
	}

	return UnknownString
}

type LAGGMemberList string

func (list LAGGMemberList) List() []int {
	return getLAGGPortListByMap(string(list))
}

func (list LAGGMemberList) String() string {
	ports := getLAGGPortListByMap(string(list))

	str := ""
	for _, port := range ports {
		str = fmt.Sprintf("%s %v", str, port)
	}

	return strings.TrimSpace(str)
}

func getLAGGPortListByMap(sMap string) []int {
	aMap, _ := base64.StdEncoding.DecodeString(sMap)

	var members []int

	for i := 0; i < len(aMap); i++ {
		iNum := aMap[i]
		for k := 7; k >= 0; k-- {
			if iNum&(1<<k) != 0 {
				members = append(members, i*8+8-k)
			}
		}
	}

	return members
}

// LAGGMemberSelectedStatus
//
//	1 - Selected,
//	2 - Unselected,
//	3 - Individual.
type LAGGMemberSelectedStatus int

func (status LAGGMemberSelectedStatus) String() string {
	switch status {
	case LAGGMemberSelectedStatusSelected:
		return LAGGMemberSelectedStatusSelectedString
	case LAGGMemberSelectedStatusUnselected:
		return LAGGMemberSelectedStatusUnselectedString
	case LAGGMemberSelectedStatusIndividual:
		return LAGGMemberSelectedStatusIndividualString
	}

	return UnknownString
}
