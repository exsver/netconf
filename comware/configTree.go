package comware

import (
	"encoding/xml"

	"github.com/exsver/netconf/netconf"
)

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
	RBAC            *RBAC            `xml:"RBAC"`
	ResourceMonitor *ResourceMonitor `xml:"ResourceMonitor"`
	Route           *Route           `xml:"Route"`
	SNMP            *SNMP            `xml:"SNMP"`
	StaticRoute     *StaticRoute     `xml:"StaticRoute"`
	STP             *STP             `xml:"STP"`
	Syslog          *Syslog          `xml:"Syslog"`
	VLAN            *VLAN            `xml:"VLAN"`
}

func (targetDevice *TargetDevice) GetData() (*Top, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top, nil
}
