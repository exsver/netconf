package junos

import "encoding/xml"

type Vlans struct {
	XMLName xml.Name `xml:"vlans"`
	Vlans   []Vlan   `xml:"vlan"`
}

type Vlan struct {
	XMLName     xml.Name `xml:"vlan"`
	Name        string   `xml:"name,omitempty"`
	Description string   `xml:"description,omitempty"`
	VlanID      int      `xml:"vlan-id"`
	L3Interface string   `xml:"l3-interface,omitempty"`
}
