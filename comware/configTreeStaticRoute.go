package comware

import "encoding/xml"

type StaticRoute struct {
	/* top level
	   StaticRoute
	     Ipv4StaticRouteConfigurations
	       []RouteEntry
	     Ipv6StaticRouteConfigurations
	       []RouteEntry
	*/
	Ipv4StaticRouteConfigurations *Ipv4StaticRouteConfigurations `xml:"Ipv4StaticRouteConfigurations"`
	Ipv6StaticRouteConfigurations *Ipv6StaticRouteConfigurations `xml:"Ipv6StaticRouteConfigurations"`
}

type Ipv4StaticRouteConfigurations struct {
	Routes []IPv4StaticRoute `xml:"RouteEntry"`
}

type IPv4StaticRoute struct {
	XMLName            xml.Name `xml:"RouteEntry"`
	Ipv4Address        string   `xml:"Ipv4Address"`
	NexthopIpv4Address string   `xml:"NexthopIpv4Address,omitempty"`
	Description        string   `xml:"Description,omitempty"`
	DestVrfIndex       int      `xml:"DestVrfIndex"`
	DestTopologyIndex  int      `xml:"DestTopologyIndex"`
	Ipv4PrefixLength   int      `xml:"Ipv4PrefixLength,omitempty"`
	NexthopVrfIndex    int      `xml:"NexthopVrfIndex"`
	IfIndex            int      `xml:"IfIndex"`
	Tag                int      `xml:"Tag,omitempty"`
	Preference         int      `xml:"Preference,omitempty"`
	Permanent          bool     `xml:"Permanent"`
	// RecursiveHost - Recursive look up ARP host route.
	RecursiveHost      bool     `xml:"RecursiveHost"`
}

type Ipv6StaticRouteConfigurations struct {
	Routes []IPv6StaticRoute `xml:"RouteEntry"`
}

type IPv6StaticRoute struct {
	XMLName            xml.Name `xml:"RouteEntry"`
	Ipv6Address        string   `xml:"Ipv6Address"`
	NexthopIpv6Address string   `xml:"NexthopIpv6Address,omitempty"`
	Description        string   `xml:"Description,omitempty"`
	DestVrfIndex       int      `xml:"DestVrfIndex"`
	DestTopologyIndex  int      `xml:"DestTopologyIndex"`
	Ipv6PrefixLength   int      `xml:"Ipv6PrefixLength,omitempty"`
	NexthopVrfIndex    int      `xml:"NexthopVrfIndex"`
	IfIndex            int      `xml:"IfIndex"`
	Tag                int      `xml:"Tag,omitempty"`
	Preference         int      `xml:"Preference,omitempty"`
	Permanent          bool     `xml:"Permanent"`
}
