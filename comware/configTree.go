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
	IPCIM           *IPCIM           `xml:"IPCIM"`
	Ifmgr           *Ifmgr           `xml:"Ifmgr"`
	MAC             *MAC             `xml:"MAC"`
	MGROUP          *MGROUP          `xml:"MGROUP"`
	PortSecurity    *PortSecurity    `xml:"PortSecurity"`
	ResourceMonitor *ResourceMonitor `xml:"ResourceMonitor"`
	STP             *STP             `xml:"STP"`
	Syslog          *Syslog          `xml:"Syslog"`
	VLAN            *VLAN            `xml:"VLAN"`
	IRF             *IRF             `xml:"IRF"`
}
