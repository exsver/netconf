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
	IPv4Routes *IPv4Routes `xml:"Ipv4Routes"`
}

type IPv4Routes struct {
	Routes []IPv4Route `xml:"RouteEntry"`
}

type IPv4Route struct {
	XMLName xml.Name `xml:"RouteEntry"`
	Nexthop string   `xml:"Nexthop,omitempty"`
	IfIndex int      `xml:"IfIndex,omitempty"`
	IPv4    *IPv4    `xml:"Ipv4"`
}

type IPv4 struct {
	XMLName          xml.Name `xml:"Ipv4"`
	Ipv4Address      string   `xml:"Ipv4Address"`
	Ipv4PrefixLength int      `xml:"Ipv4PrefixLength"`
}
