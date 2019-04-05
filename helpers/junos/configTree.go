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
	ForwardingOptions *ForwardingOptions `xml:"forwarding-options"`
	RoutingOptions    *RoutingOptions    `xml:"routing-options"`
	Protocols         *Protocols         `xml:"protocols"`
	Vlans             *Vlans             `xml:"vlans"`
}

type System struct {
	XMLName            xml.Name            `xml:"system"`
	HostName           string              `xml:"host-name,omitempty"`
	ARP                *ARP                `xml:"arp"`
	RootAuthentication *RootAuthentication `xml:"root-authentication"`
	Login              *Login              `xml:"login"`
	Services           *Services           `xml:"services"`
	Syslog             *Syslog             `xml:"syslog"`
	NTP                *NTP                `xml:"ntp"`
}

type ARP struct {
	XMLName    xml.Name `xml:"arp"`
	AgingTimer int      `xml:"aging-timer"`
}

type RootAuthentication struct {
	XMLName           xml.Name `xml:"root-authentication"`
	EncryptedPassword string   `xml:"encrypted-password,omitempty"`
}

type Login struct {
	XMLName xml.Name `xml:"login"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName        xml.Name        `xml:"user"`
	Name           string          `xml:"name,omitempty"`
	UID            int             `xml:"uid,omitempty"`
	Class          string          `xml:"class,omitempty"`
	Authentication *Authentication `xml:"authentication"`
}

type Authentication struct {
	XMLName           xml.Name `xml:"authentication"`
	EncryptedPassword string   `xml:"encrypted-password,omitempty"`
}

type Services struct {
	XMLName xml.Name `xml:"services"`
}

type Syslog struct {
	XMLName xml.Name `xml:"syslog"`
}

type NTP struct {
	XMLName xml.Name `xml:"ntp"`
}

type Chassis struct {
	XMLName           xml.Name           `xml:"chassis"`
	AggregatedDevices *AggregatedDevices `xml:"aggregated-devices"`
}

type AggregatedDevices struct {
	XMLName  xml.Name  `xml:"aggregated-devices"`
	Ethernet *Ethernet `xml:"ethernet"`
}

type Ethernet struct {
	XMLName     xml.Name `xml:"ethernet"`
	DeviceCount int      `xml:"device-count"`
}

type Interfaces struct {
	XMLName    xml.Name `xml:"interfaces"`
	Interfaces []Iface  `xml:"interface"`
}

type Iface struct {
	XMLName xml.Name `xml:"interface"`
	Name    string   `xml:"name"`
	Units   []Unit   `xml:"unit"`
}

type Unit struct {
	XMLName xml.Name `xml:"unit"`
	Name    int      `xml:"name"`
}

type ForwardingOptions struct {
	XMLName xml.Name `xml:"forwarding-options"`
}

type RoutingOptions struct {
	XMLName xml.Name              `xml:"routing-options"`
	Static  *RoutingOptionsStatic `xml:"static"`
}

type RoutingOptionsStatic struct {
	XMLName xml.Name `xml:"static"`
	Routes  []Route  `xml:"route"`
}

type Route struct {
	XMLName xml.Name `xml:"route"`
	Name    string   `xml:"name"`
	NextHop string   `xml:"next-hop"`
}

type Protocols struct {
	XMLName xml.Name `xml:"protocols"`
}

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
