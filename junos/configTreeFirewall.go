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
	Name    string      `xml:"name"`
	From    *FilterFrom `xml:"from"`
	Then    *FilterThen `xml:"then"`
}

type FilterFrom struct {
	XMLName               xml.Name                   `xml:"from"`
	Interfaces            []FilterInterface          `xml:"interface,omitempty"`
	Address               []FilterAddress            `xml:"address,omitempty"`
	SourceAddress         []FilterSourceAddress      `xml:"source-address,omitempty"`
	DestinationAddress    []FilterDestinationAddress `xml:"destination-address,omitempty"`
	PrefixList            []FilterPrefixList         `xml:"prefix-list,omitempty"`
	SourcePrefixList      []FilterPrefixList         `xml:"source-prefix-list,omitempty"`
	DestinationPrefixList []FilterPrefixList         `xml:"destination-prefix-list,omitempty"`
	TTL                   []string                   `xml:"ttl,omitempty"`
	TTLExcept             []string                   `xml:"ttl-except,omitempty"`
	Protocol              []string                   `xml:"protocol,omitempty"`
	ProtocolExcept        []string                   `xml:"protocol-except,omitempty"`
	Port                  []string                   `xml:"port,omitempty"`
	PortExcept            []string                   `xml:"port-except,omitempty"`
	SourcePort            []string                   `xml:"source-port,omitempty"`
	SourcePortExcept      []string                   `xml:"source-port-except,omitempty"`
	DestinationPort       []string                   `xml:"destination-port,omitempty"`
	DestinationPortExcept []string                   `xml:"destination-port-except,omitempty"`
	ICMPType              []string                   `xml:"icmp-type,omitempty"`
	ICMPTypeExcept        []string                   `xml:"icmp-type-except,omitempty"`
	ICMPCode              []string                   `xml:"icmp-code,omitempty"`
	ICMPCodeExcept        []string                   `xml:"icmp-code-except,omitempty"`
	IPOptions             []string                   `xml:"ip-options,omitempty"`
	IPOptionsExcept       []string                   `xml:"ip-options-except,omitempty"`
	DSCP                  []string                   `xml:"dscp,omitempty"`
	DSCPExcept            []string                   `xml:"dscp-except,omitempty"`
	SourceClass           []string                   `xml:"source-class,omitempty"`
	DestinationClass      []string                   `xml:"destination-class,omitempty"`
}

type FilterAddress struct {
	XMLName xml.Name `xml:"address"`
	Name    string   `xml:"name"`
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

type FilterPrefixList struct {
	XMLName xml.Name `xml:"prefix-list"`
	Name    string   `xml:"name"`
}

type FilterThen struct {
	XMLName         xml.Name                   `xml:"then"`
	Policer         string                     `xml:"policer,omitempty"`
	RoutingInstance *FilterThenRoutingInstance `xml:"routing-instance"`
}

type FilterThenRoutingInstance struct {
	XMLName             xml.Name `xml:"routing-instance"`
	RoutingInstanceName string   `xml:"routing-instance-name"`
}

type Policer struct {
	XMLName xml.Name `xml:"policer"`
	Name    string   `xml:"name"`
}
