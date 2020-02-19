package junos

import "encoding/xml"

type Protocols struct {
	XMLName xml.Name      `xml:"protocols"`
	BGP     *ProtocolBGP  `xml:"bgp"`
	IsIs    *ProtocolIsIs `xml:"isis"`
	OSPF    *ProtocolOSPF `xml:"ospf"`
	Lacp    *ProtocolLacp `xml:"lacp"`
	VRRP    *ProtocolVRRP `xml:"vrrp"`
}

type ProtocolBGP struct {
	XMLName      xml.Name   `xml:"bgp"`
	Description  string     `xml:"description,omitempty"`
	Groups       []BGPGroup `xml:"group,omitempty"`
	Disable      bool       `xml:"disable,omitempty"`
	Passive      bool       `xml:"passive,omitempty"`
	MTUDiscovery bool       `xml:"mtu-discovery,omitempty"`
}

type BGPGroup struct {
	XMLName           xml.Name `xml:"group"`
	Name              string   `xml:"name"`
	Type              string   `xml:"type,omitempty"`
	Description       string   `xml:"description,omitempty"`
	Damping           bool     `xml:"damping,omitempty"`
	EnforceFirstAS    bool     `xml:"enforce-first-as,omitempty"`
	NoAdvertisePeerAS bool     `xml:"no-advertise-peer-as,omitempty"`
	NoAggregatorID    bool     `xml:"no-aggregator-id,omitempty"`
	NoClientReflect   bool     `xml:"no-client-reflect,omitempty"`
}

type ProtocolIsIs struct {
	XMLName xml.Name `xml:"isis"`
	Disable bool     `xml:"disable,omitempty"`
}

type ProtocolOSPF struct {
	XMLName xml.Name `xml:"ospf"`
	Disable bool     `xml:"disable,omitempty"`
}

type ProtocolLacp struct {
	XMLName       xml.Name `xml:"lacp"`
	FastHelloIssu bool     `xml:"fast-hello-issu,omitempty"`
}

type ProtocolVRRP struct {
	XMLName            xml.Name `xml:"vrrp"`
	AsymmetricHoldTime bool     `xml:"asymmetric-hold-time,omitempty"`
	SkewTimerDisable   bool     `xml:"skew-timer-disable,omitempty"`
	Version3           bool     `xml:"version-3,omitempty"`
}
