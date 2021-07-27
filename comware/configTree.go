package comware

import "encoding/xml"

// Data used to collect information from device
type Data struct {
	XMLName   xml.Name   `xml:"data"`
	Top       *Top       `xml:"top"`
	SavePoint *SavePoint `xml:"save-point"`
}

type Filter struct {
	XMLName xml.Name `xml:"filter"`
	Top     Top      `xml:"top"`
}

// Top used to configure device
type Top struct {
	XMLName         xml.Name         `xml:"top"`
	ACL             *ACL             `xml:"ACL"`
	ARP             *ARP             `xml:"ARP"`
	Device          *Device          `xml:"Device"`
	DHCP            *DHCP            `xml:"DHCP"`
	DHCPSP          *DHCPSP          `xml:"DHCPSP"`
	Diagnostic      *Diagnostic      `xml:"Diagnostic"`
	DNS             *DNS             `xml:"DNS"`
	FileSystem      *FileSystem      `xml:"FileSystem"`
	IPCIM           *IPCIM           `xml:"IPCIM"` // IP Source Guard
	Ifmgr           *Ifmgr           `xml:"Ifmgr"` // Interfaces
	IRF             *IRF             `xml:"IRF"`
	LAGG            *LAGG            `xml:"LAGG"` // Link Aggregation
	LR              *LR              `xml:"LR"`   // Interface Rate Limit
	MAC             *MAC             `xml:"MAC"`
	MGROUP          *MGROUP          `xml:"MGROUP"` // Port Mirroring
	PoE             *PoE             `xml:"PoE"`
	PortSecurity    *PortSecurity    `xml:"PortSecurity"`
	ResourceMonitor *ResourceMonitor `xml:"ResourceMonitor"`
	Route           *Route           `xml:"Route"`
	SNMP            *SNMP            `xml:"SNMP"`
	StaticRoute     *StaticRoute     `xml:"StaticRoute"`
	STP             *STP             `xml:"STP"`
	Syslog          *Syslog          `xml:"Syslog"`
	VLAN            *VLAN            `xml:"VLAN"`
}
