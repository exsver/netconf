package junos

import "encoding/xml"

type RoutingOptions struct {
	XMLName xml.Name              `xml:"routing-options"`
	Static  *RoutingOptionsStatic `xml:"static"`
}

type RoutingOptionsStatic struct {
	XMLName xml.Name      `xml:"static"`
	Routes  []StaticRoute `xml:"route"`
}

type StaticRoute struct {
	XMLName                xml.Name `xml:"route"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	NextTable              string   `xml:"next-table,omitempty"`
	NextHop                []string `xml:"next-hop"`
	Active                 bool     `xml:"active,omitempty"`
	Passive                bool     `xml:"passive,omitempty"`
	Discard                bool     `xml:"discard,omitempty"`
	Readvertise            bool     `xml:"readvertise,omitempty"`
	NoReadvertise          bool     `xml:"no-readvertise,omitempty"`
	Receive                bool     `xml:"receive,omitempty"`
	Reject                 bool     `xml:"reject,omitempty"`
	Resolve                bool     `xml:"resolve,omitempty"`
	NoResolve              bool     `xml:"no-resolve,omitempty"`
	Retain                 bool     `xml:"retain,omitempty"`
	NoRetain               bool     `xml:"no-retain,omitempty"`
	Install                bool     `xml:"install,omitempty"`
	NoInstall              bool     `xml:"no-install,omitempty"`
	LongestMatch           bool     `xml:"longest-match,omitempty"`
	NoLongestMatch         bool     `xml:"no-longest-match,omitempty"`
}

func (staticRoute *StaticRoute) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			RoutingOptions: &RoutingOptions{
				Static: &RoutingOptionsStatic{
					Routes: []StaticRoute{*staticRoute},
				},
			},
		},
	}
}
