package junos

import "encoding/xml"

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
