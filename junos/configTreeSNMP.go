package junos

import "encoding/xml"

type SNMP struct {
	XMLName     xml.Name        `xml:"snmp"`
	SystemName  string          `xml:"system-name,omitempty"`
	Description string          `xml:"description,omitempty"`
	Location    string          `xml:"location,omitempty"`
	Contact     string          `xml:"contact,omitempty"`
	Communities []SNMPCommunity `xml:"community,omitempty"`
}

type SNMPCommunity struct {
	XMLName       xml.Name `xml:"community"`
	Name          string   `xml:"name"`                    // community string
	Authorization string   `xml:"authorization,omitempty"` // read-only | read-write
}
