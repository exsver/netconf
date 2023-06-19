package comware

import "encoding/xml"

type Route struct {
	/* top level
	   Route
	     Ipv4RouteECMPNumbers
	       []RouteECMPNumber
	     Ipv4RouteStatistics
	       []Statistics
	     Ipv4Routes
	       []RouteEntry
	     Ipv6RouteECMPNumbers
	       []RouteECMPNumber
	     Ipv6RouteStatistics
	       []Statistics
	     Ipv6Routes
	       []RouteEntry
	*/
	IPv4RouteECMPNumbers *IPv4RouteECMPNumbers `xml:"Ipv4RouteECMPNumbers"`
	IPv4RouteStatistics  *RouteStatistics      `xml:"Ipv4RouteStatistics"`
	IPv4Routes           *IPv4Routes           `xml:"Ipv4Routes"`
	IPv6RouteECMPNumbers *IPv6RouteECMPNumbers `xml:"Ipv6RouteECMPNumbers"`
	IPv6RouteStatistics  *RouteStatistics      `xml:"Ipv6RouteStatistics"`
	IPv6Routes           *IPv6Routes           `xml:"Ipv6Routes"`
}

type IPv4RouteECMPNumbers struct {
	Routes []IPv4RouteECMPNumber `xml:"RouteECMPNumber"`
}

type IPv4RouteECMPNumber struct {
	XMLName    xml.Name        `xml:"RouteECMPNumber"`
	IPv4       *RouteEntryIPv4 `xml:"Ipv4"`
	ECMPNumber int             `xml:"ECMPNumber"`
}

type RouteStatistics struct {
	Statistics []RouteStatistic `xml:"Statistics"`
}

type RouteStatistic struct {
	XMLName       xml.Name `xml:"Statistics"`
	ProtocolID    int      `xml:"ProtocolID"`
	TotalRoutes   int      `xml:"TotalRoutes"`
	ActiveRoutes  int      `xml:"ActiveRoutes"`
	AddedRoutes   int      `xml:"AddedRoutes"`
	DeletedRoutes int      `xml:"DeletedRoutes"`
}

type IPv4Routes struct {
	Routes []IPv4Route `xml:"RouteEntry"`
}

type IPv4Route struct {
	XMLName    xml.Name            `xml:"RouteEntry"`
	IPv4       *RouteEntryIPv4     `xml:"Ipv4"`
	Nexthop    string              `xml:"Nexthop,omitempty"`
	IfIndex    int                 `xml:"IfIndex,omitempty"`
	Protocol   *RouteEntryProtocol `xml:"Protocol"`
	Age        int                 `xml:"Age,omitempty"`
	Preference int                 `xml:"Preference,omitempty"`
	Metric     int                 `xml:"Metric,omitempty"`
	Tag        int                 `xml:"Tag,omitempty"`
	Neighbor   string              `xml:"Neighbor,omitempty"`
	ASNumber   *RouteEntryASNumber `xml:"ASNumber"`
}

type IPv6RouteECMPNumbers struct {
	Routes []IPv6RouteECMPNumber `xml:"RouteECMPNumber"`
}

type IPv6RouteECMPNumber struct {
	XMLName    xml.Name              `xml:"RouteECMPNumber"`
	IPv6       *RouteEntryIpv6Prefix `xml:"Ipv6Prefix"`
	ECMPNumber int                   `xml:"ECMPNumber"`
}

type IPv6Routes struct {
	Routes []IPv6Route `xml:"RouteEntry"`
}

type IPv6Route struct {
	XMLName    xml.Name              `xml:"RouteEntry"`
	IPv6       *RouteEntryIpv6Prefix `xml:"Ipv6Prefix"`
	Nexthop    string                `xml:"Nexthop,omitempty"`
	IfIndex    int                   `xml:"IfIndex,omitempty"`
	Protocol   *RouteEntryProtocol   `xml:"Protocol"`
	Age        int                   `xml:"Age,omitempty"`
	Preference int                   `xml:"Preference,omitempty"`
	Metric     int                   `xml:"Metric,omitempty"`
	Tag        int                   `xml:"Tag,omitempty"`
	Neighbor   string                `xml:"Neighbor,omitempty"`
	ASNumber   *RouteEntryASNumber   `xml:"ASNumber"`
}

type RouteEntryIPv4 struct {
	XMLName          xml.Name `xml:"Ipv4"`
	Ipv4Address      string   `xml:"Ipv4Address"`
	Ipv4PrefixLength int      `xml:"Ipv4PrefixLength"`
}

type RouteEntryIpv6Prefix struct {
	XMLName          xml.Name `xml:"Ipv6Prefix"`
	Ipv6Address      string   `xml:"Ipv6Address"`
	Ipv6PrefixLength int      `xml:"Ipv6PrefixLength"`
}

type RouteEntryProtocol struct {
	XMLName       xml.Name `xml:"Protocol"`
	ProtocolID    int      `xml:"ProtocolID"`
	SubProtocolID int      `xml:"SubProtocolID"`
}

type RouteEntryASNumber struct {
	XMLName  xml.Name `xml:"ASNumber"`
	OriginAS int      `xml:"OriginAS"`
	LastAS   int      `xml:"LastAS"`
}
