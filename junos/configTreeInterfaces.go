package junos

import "encoding/xml"

type Interfaces struct {
	XMLName    xml.Name `xml:"interfaces"`
	Interfaces []Iface  `xml:"interface"`
}

type Iface struct {
	XMLName                xml.Name                `xml:"interface"`
	Name                   string                  `xml:"name"`
	Description            string                  `xml:"description,omitempty"`
	Encapsulation          string                  `xml:"encapsulation,omitempty"`
	Mtu                    int                     `xml:"mtu,omitempty"`
	AggregatedEtherOptions *AggregatedEtherOptions `xml:"aggregated-ether-options"`
	GigetherOptions        *GigetherOptions        `xml:"gigether-options"`
	OpticsOptions          *OpticsOptions          `xml:"optics-options"`
	Units                  []Unit                  `xml:"unit"`
}

type Unit struct {
	XMLName xml.Name `xml:"unit"`
	Name    int      `xml:"name"`
}

type AggregatedEtherOptions struct {
	XMLName      xml.Name `xml:"aggregated-ether-options"`
	MinimumLinks int      `xml:"minimum-links,omitempty"` // Minimum number of aggregated links (1..64)
	LinkSpeed    string   `xml:"link-speed,omitempty"`
	Lacp         *LACP    `xml:"lacp"`
}

type LACP struct {
	XMLName  xml.Name `xml:"lacp"`
	Periodic string   `xml:"periodic,omitempty"`
}

type GigetherOptions struct {
	XMLName xml.Name `xml:"gigether-options"`
}

type OpticsOptions struct {
	XMLName xml.Name `xml:"optics-options"`
}
