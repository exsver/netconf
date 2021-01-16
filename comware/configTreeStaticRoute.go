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
	Routes []IPv4RouteEntry `xml:"RouteEntry"`
}

type IPv4RouteEntry struct {
	XMLName            xml.Name `xml:"RouteEntry"`
	Ipv4Address        string   `xml:"Ipv4Address,omitempty"`
	NexthopIpv4Address string   `xml:"NexthopIpv4Address,omitempty"`
	Description        string   `xml:"Description,omitempty"`
	DestVrfIndex       int      `xml:"DestVrfIndex,omitempty"`
	DestTopologyIndex  int      `xml:"DestTopologyIndex,omitempty"`
	Ipv4PrefixLength   int      `xml:"Ipv4PrefixLength,omitempty"`
	NexthopVrfIndex    int      `xml:"NexthopVrfIndex,omitempty"`
	IfIndex            int      `xml:"IfIndex,omitempty"`
	Tag                int      `xml:"Tag,omitempty"`
	Preference         int      `xml:"Preference,omitempty"`
	Permanent          bool     `xml:"Permanent"`
	RecursiveHost      bool     `xml:"RecursiveHost"`
}

type Ipv6StaticRouteConfigurations struct {
	Routes []IPv6RouteEntry `xml:"RouteEntry"`
}

type IPv6RouteEntry struct {
	XMLName            xml.Name `xml:"RouteEntry"`
	Ipv6Address        string   `xml:"Ipv6Address,omitempty"`
	NexthopIpv6Address string   `xml:"NexthopIpv6Address,omitempty"`
	Description        string   `xml:"Description,omitempty"`
	DestVrfIndex       int      `xml:"DestVrfIndex,omitempty"`
	DestTopologyIndex  int      `xml:"DestTopologyIndex,omitempty"`
	Ipv6PrefixLength   int      `xml:"Ipv6PrefixLength,omitempty"`
	NexthopVrfIndex    int      `xml:"NexthopVrfIndex,omitempty"`
	IfIndex            int      `xml:"IfIndex,omitempty"`
	Tag                int      `xml:"Tag,omitempty"`
	Preference         int      `xml:"Preference,omitempty"`
	Permanent          bool     `xml:"Permanent"`
	RecursiveHost      bool     `xml:"RecursiveHost"`
}
