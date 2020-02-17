package junos

import "encoding/xml"

// Config used to configure device
type Config struct {
	XMLName       xml.Name       `xml:"config"`
	Configuration *Configuration `xml:"configuration"`
}

// Data used to collect information from device
type Data struct {
	XMLName       xml.Name       `xml:"data"`
	Configuration *Configuration `xml:"configuration"`
}

type Configuration struct {
	XMLName           xml.Name           `xml:"configuration"`
	Version           string             `xml:"version,omitempty"`
	System            *System            `xml:"system"`
	Chassis           *Chassis           `xml:"chassis"`
	Interfaces        *Interfaces        `xml:"interfaces"`
	SNMP              *SNMP              `xml:"snmp"`
	ForwardingOptions *ForwardingOptions `xml:"forwarding-options"`
	RoutingOptions    *RoutingOptions    `xml:"routing-options"`
	Protocols         *Protocols         `xml:"protocols"`
	PolicyOptions     *PolicyOptions     `xml:"policy-options"`
	Firewall          *Firewall          `xml:"firewall"`
	RoutingInstances  *RoutingInstances  `xml:"routing-instances"`
	Vlans             *Vlans             `xml:"vlans"`
}
