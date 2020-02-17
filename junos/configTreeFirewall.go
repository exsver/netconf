package junos

import "encoding/xml"

type Firewall struct {
	XMLName  xml.Name  `xml:"firewall"`
	Filters  []Filter  `xml:"filter,omitempty"`
	Policers []Policer `xml:"policer,omitempty"`
}

type Filter struct {
	XMLName xml.Name `xml:"filter"`
	Name    string   `xml:"name"`
	Terms   []Term   `xml:"term"`
}

type Term struct {
	XMLName xml.Name    `xml:"term"`
	From    *FilterFrom `xml:"from"`
	Then    *FilterThen `xml:"then"`
}

type FilterFrom struct {
	XMLName            xml.Name                   `xml:"from"`
	Interfaces         []FilterInterface          `xml:"interface,omitempty"`
	SourceAddress      []FilterSourceAddress      `xml:"source-address,omitempty"`
	DestinationAddress []FilterDestinationAddress `xml:"destination-address,omitempty"`
	TTL                []string                   `xml:"ttl,omitempty"`
	Protocol           []string                   `xml:"protocol,omitempty"`
	ProtocolExcept     []string                   `xml:"protocol-except,omitempty"`
	Port               []string                   `xml:"port,omitempty"`
	PortExcept         []string                   `xml:"port-except,omitempty"`
	SourcePort         []string                   `xml:"source-port,omitempty"`
	DestinationPort    []string                   `xml:"destination-port,omitempty"`
}

type FilterSourceAddress struct {
	XMLName xml.Name `xml:"source-address"`
	Name    string   `xml:"name"`
}

type FilterDestinationAddress struct {
	XMLName xml.Name `xml:"destination-address"`
	Name    string   `xml:"name"`
}

type FilterInterface struct {
	XMLName xml.Name `xml:"interface"`
	Name    string   `xml:"name"`
}

type FilterThen struct {
	XMLName xml.Name `xml:"then"`
}

type Policer struct {
	XMLName xml.Name `xml:"policer"`
	Name    string   `xml:"name"`
}
