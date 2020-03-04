package junos

import "encoding/xml"

type Interfaces struct {
	XMLName    xml.Name `xml:"interfaces"`
	Interfaces []Iface  `xml:"interface"`
}

type Iface struct {
	XMLName                xml.Name                `xml:"interface"`
	NetconfConfigOperation string                  `xml:"operation,attr,omitempty"`
	Name                   string                  `xml:"name"`
	Description            string                  `xml:"description,omitempty"`
	Encapsulation          string                  `xml:"encapsulation,omitempty"`
	Mtu                    int                     `xml:"mtu,omitempty"`
	AggregatedEtherOptions *AggregatedEtherOptions `xml:"aggregated-ether-options"`
	GigetherOptions        *GigetherOptions        `xml:"gigether-options"`
	OpticsOptions          *OpticsOptions          `xml:"optics-options"`
	Units                  []Unit                  `xml:"unit"`
}

type Unit struct {
	XMLName                xml.Name    `xml:"unit"`
	NetconfConfigOperation string      `xml:"operation,attr,omitempty"`
	Name                   string      `xml:"name"`
	Description            string      `xml:"description,omitempty"`
	Encapsulation          string      `xml:"encapsulation,omitempty"`
	VlanIDList             string      `xml:"vlan-id-list,omitempty"`
	VlanID                 int         `xml:"vlan-id,omitempty"`
	NativeInnerVlanID      int         `xml:"native-inner-vlan-id,omitempty"`
	Disable                bool        `xml:"disable,omitempty"`
	Traps                  bool        `xml:"traps,omitempty"`
	NoTraps                bool        `xml:"no-traps,omitempty"`
	Family                 *UnitFamily `xml:"family"`
	Filter                 *UnitFilter `xml:"filter"` // Filters to apply to all families configured under this logical interface
}

type UnitFamily struct {
	XMLName xml.Name         `xml:"family"`
	Inet    *UnitFamilyInet  `xml:"inet"`
	Inet6   *UnitFamilyInet6 `xml:"inet6"`
}

type UnitFamilyInet struct {
	XMLName xml.Name                `xml:"inet"`
	Address []UnitFamilyInetAddress `xml:"address"`
	Filter  *UnitFamilyFilter       `xml:"filter"`
}

type UnitFamilyInetAddress struct {
	XMLName                xml.Name                   `xml:"address"`
	NetconfConfigOperation string                     `xml:"operation,attr,omitempty"`
	Name                   string                     `xml:"name"`
	Broadcast              string                     `xml:"broadcast,omitempty"`
	Primary                bool                       `xml:"primary,omitempty"`
	Preferred              bool                       `xml:"preferred,omitempty"`
	ARP                    []UnitFamilyInetAddressARP `xml:"arp,omitempty"`
	VRRPGroup              []VRRPGroup                `xml:"vrrp-group,omitempty"`
}

type UnitFamilyInetAddressARP struct {
	XMLName      xml.Name `xml:"arp"`
	Name         string   `xml:"name"`
	MAC          string   `xml:"mac,omitempty"`
	MulticastMAC string   `xml:"multicast-mac,omitempty"`
	Publish      bool     `xml:"publish,omitempty"`
}

type VRRPGroup struct {
	XMLName            xml.Name          `xml:"vrrp-group"`
	Name               string            `xml:"name"`
	AuthenticationType string            `xml:"authentication-type,omitempty"`
	AuthenticationKey  string            `xml:"authentication-key,omitempty"`
	VirtualAddress     []string          `xml:"virtual-address,omitempty"`
	Priority           int               `xml:"priority,omitempty"`      // Virtual router election priority (0..255)
	FastInterval       int               `xml:"fast-interval,omitempty"` // Fast advertisement interval (10..40950 milliseconds)
	AcceptData         bool              `xml:"accept-data,omitempty"`
	NoAcceptData       bool              `xml:"no-accept-data,omitempty"`
	NoPreempt          bool              `xml:"no-preempt,omitempty"`
	Preempt            *VRRPGroupPreempt `xml:"preempt"`
}

type VRRPGroupPreempt struct {
	XMLName  xml.Name `xml:"preempt"`
	HoldTime int      `xml:"hold-time"` // Preemption hold time (0..3600 seconds)
}

type UnitFilter struct {
	XMLName xml.Name `xml:"filter"`
	Input   string   `xml:"input,omitempty"`  // Name of filter applied to received packets
	Output  string   `xml:"output,omitempty"` // Name of filter applied to transmitted packets
}

type UnitFamilyFilter struct {
	XMLName    xml.Name            `xml:"filter"`
	InputList  []string            `xml:"input-list,omitempty"`
	OutputList []string            `xml:"output-list,omitempty"`
	Input      *FamilyFilterInput  `xml:"input"`
	Output     *FamilyFilterOutput `xml:"output"`
}

type FamilyFilterInput struct {
	XMLName    xml.Name `xml:"input"`
	FilterName string   `xml:"filter-name,omitempty"`
}

type FamilyFilterOutput struct {
	XMLName    xml.Name `xml:"output"`
	FilterName string   `xml:"filter-name,omitempty"`
}

type UnitFamilyInet6 struct {
	XMLName xml.Name                 `xml:"inet6"`
	Address []UnitFamilyInet6Address `xml:"address"`
}

type UnitFamilyInet6Address struct {
	XMLName   xml.Name         `xml:"address"`
	Name      string           `xml:"name"`
	VRRPGroup []VRRPInet6Group `xml:"vrrp-inet6-group,omitempty"`
}

type VRRPInet6Group struct {
	XMLName                 xml.Name `xml:"vrrp-inet6-group"`
	Name                    string   `xml:"name"`
	VirtualLinkLocalAddress string   `xml:"virtual-link-local-address,omitempty"`
	VirtualAddress          []string `xml:"virtual-inet6-address,omitempty"`
	Priority                int      `xml:"priority,omitempty"`
	AcceptData              bool     `xml:"accept-data,omitempty"`
}

type AggregatedEtherOptions struct {
	XMLName                           xml.Name                    `xml:"aggregated-ether-options"`
	MinimumLinks                      int                         `xml:"minimum-links,omitempty"` // Minimum number of aggregated links (1..64)
	LinkSpeed                         string                      `xml:"link-speed,omitempty"`
	LACP                              *AggregatedEtherOptionsLACP `xml:"lacp"`
	Loopback                          bool                        `xml:"loopback,omitempty"`
	NoLoopback                        bool                        `xml:"no-loopback,omitempty"`
	FlowControl                       bool                        `xml:"flow-control,omitempty"`
	NoFlowControl                     bool                        `xml:"no-flow-control,omitempty"`
	NoSourceFiltering                 bool                        `xml:"no-source-filtering,omitempty"`
	PadToMinimumFrameSize             bool                        `xml:"pad-to-minimum-frame-size,omitempty"`
	LogicalInterfaceChassisRedundancy bool                        `xml:"logical-interface-chassis-redundancy,omitempty"`
	LogicalInterfaceFpcRedundancy     bool                        `xml:"logical-interface-fpc-redundancy,omitempty"`
}

type AggregatedEtherOptionsLACP struct {
	XMLName      xml.Name `xml:"lacp"`
	Periodic     string   `xml:"periodic,omitempty"`   // "fast"
	SyncReset    string   `xml:"sync-reset,omitempty"` // "enable" | "disable"
	Active       bool     `xml:"active,omitempty"`
	Passive      bool     `xml:"passive,omitempty"`
	AcceptData   bool     `xml:"accept-data,omitempty"`
	FastFailover bool     `xml:"fast-failover,omitempty"`
	ForceUp      bool     `xml:"force-up,omitempty"`
}

type GigetherOptions struct {
	XMLName                  xml.Name                        `xml:"gigether-options"`
	Options8023ad            *GigetherOptionsIEEE8023ad      `xml:"ieee-802.3ad"`
	AutoNegotiation          *GigetherOptionsAutoNegotiation `xml:"auto-negotiation"`
	IgnoreL3Incompletes      bool                            `xml:"ignore-l3-incompletes,omitempty"`
	AsynchronousNotification bool                            `xml:"asynchronous-notification,omitempty"`
	Loopback                 bool                            `xml:"loopback,omitempty"`
	NoLoopback               bool                            `xml:"no-loopback,omitempty"`
	NoFlowControl            bool                            `xml:"no-flow-control,omitempty"`
	NoSourceFiltering        bool                            `xml:"no-source-filtering,omitempty"`
	NoAutoNegotiation        bool                            `xml:"no-auto-negotiation,omitempty"`
	NoAutoMdix               bool                            `xml:"no-auto-mdix,omitempty"`
	PadToMinimumFrameSize    bool                            `xml:"pad-to-minimum-frame-size,omitempty"`
}

type GigetherOptionsIEEE8023ad struct {
	XMLName   xml.Name             `xml:"ieee-802.3ad"`
	Bundle    string               `xml:"bundle,omitempty"`
	LinkIndex int                  `xml:"link-index,omitempty"` // Desired child link index within the Aggregated Interface (0..63)
	Primary   bool                 `xml:"primary,omitempty"`
	Backup    bool                 `xml:"backup,omitempty"`
	LACP      *GigetherOptionsLACP `xml:"lacp"`
}

type GigetherOptionsLACP struct {
	XMLName      xml.Name `xml:"lacp"`
	PortPriority int      `xml:"port-priority,omitempty"` // Priority of the port (0 ... 65535)
}

type GigetherOptionsAutoNegotiation struct {
	XMLName     xml.Name `xml:"auto-negotiation"`
	RemoteFault string   `xml:"remote-fault,omitempty"` // "local-interface-online" | "local-interface-offline"
}

type OpticsOptions struct {
	XMLName xml.Name `xml:"optics-options"`
}

func (iface *Iface) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			Interfaces: &Interfaces{
				Interfaces: []Iface{*iface},
			},
		},
	}
}

func (unit *Unit) ConvertToConfig(ifaceName string) *Config {
	return &Config{
		Configuration: &Configuration{
			Interfaces: &Interfaces{
				Interfaces: []Iface{
					{
						Name:  ifaceName,
						Units: []Unit{*unit},
					},
				},
			},
		},
	}
}
