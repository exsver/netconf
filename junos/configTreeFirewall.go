package junos

import "encoding/xml"

type Firewall struct {
	XMLName  xml.Name      `xml:"firewall"`
	Filters  []Filter      `xml:"filter,omitempty"`
	Policers []Policer     `xml:"policer,omitempty"`
	Family   *FilterFamily `xml:"family,omitempty"`
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
	XMLName               xml.Name                      `xml:"from"`
	Interfaces            []FilterInterface             `xml:"interface,omitempty"`
	Address               []FilterAddress               `xml:"address,omitempty"`
	SourceAddress         []FilterSourceAddress         `xml:"source-address,omitempty"`
	DestinationAddress    []FilterDestinationAddress    `xml:"destination-address,omitempty"`
	PrefixList            []FilterPrefixList            `xml:"prefix-list,omitempty"`
	SourcePrefixList      []FilterSourcePrefixList      `xml:"source-prefix-list,omitempty"`
	DestinationPrefixList []FilterDestinationPrefixList `xml:"destination-prefix-list,omitempty"`
	TTL                   []string                      `xml:"ttl,omitempty"`
	TTLExcept             []string                      `xml:"ttl-except,omitempty"`
	Protocol              []string                      `xml:"protocol,omitempty"`
	ProtocolExcept        []string                      `xml:"protocol-except,omitempty"`
	Port                  []string                      `xml:"port,omitempty"`
	PortExcept            []string                      `xml:"port-except,omitempty"`
	SourcePort            []string                      `xml:"source-port,omitempty"`
	SourcePortExcept      []string                      `xml:"source-port-except,omitempty"`
	DestinationPort       []string                      `xml:"destination-port,omitempty"`
	DestinationPortExcept []string                      `xml:"destination-port-except,omitempty"`
	ICMPType              []string                      `xml:"icmp-type,omitempty"`
	ICMPTypeExcept        []string                      `xml:"icmp-type-except,omitempty"`
	ICMPCode              []string                      `xml:"icmp-code,omitempty"`
	ICMPCodeExcept        []string                      `xml:"icmp-code-except,omitempty"`
	IPOptions             []string                      `xml:"ip-options,omitempty"`
	IPOptionsExcept       []string                      `xml:"ip-options-except,omitempty"`
	DSCP                  []string                      `xml:"dscp,omitempty"`
	DSCPExcept            []string                      `xml:"dscp-except,omitempty"`
	SourceClass           []string                      `xml:"source-class,omitempty"`
	DestinationClass      []string                      `xml:"destination-class,omitempty"`
	IsFragment            bool                          `xml:"is-fragment,omitempty"`
	FirstFragment         bool                          `xml:"first-fragment,omitempty"`
	TCPEstablished        bool                          `xml:"tcp-established,omitempty"`
	TCPInitial            bool                          `xml:"tcp-initial,omitempty"`
}

type FilterAddress struct {
	XMLName xml.Name `xml:"address"`
	Name    string   `xml:"name"`
	Except  bool     `xml:"except,omitempty"`
}

type FilterSourceAddress struct {
	XMLName xml.Name `xml:"source-address"`
	Name    string   `xml:"name"`
	Except  bool     `xml:"except,omitempty"`
}

type FilterDestinationAddress struct {
	XMLName xml.Name `xml:"destination-address"`
	Name    string   `xml:"name"`
	Except  bool     `xml:"except,omitempty"`
}

type FilterInterface struct {
	XMLName xml.Name `xml:"interface"`
	Name    string   `xml:"name"`
}

type FilterPrefixList struct {
	XMLName xml.Name `xml:"prefix-list"`
	Name    string   `xml:"name"`
	Except  bool     `xml:"except,omitempty"`
}

type FilterSourcePrefixList struct {
	XMLName xml.Name `xml:"source-prefix-list"`
	Name    string   `xml:"name"`
	Except  bool     `xml:"except,omitempty"`
}

type FilterDestinationPrefixList struct {
	XMLName xml.Name `xml:"destination-prefix-list"`
	Name    string   `xml:"name"`
	Except  bool     `xml:"except,omitempty"`
}

type FilterThen struct {
	XMLName         xml.Name                   `xml:"then"`
	Policer         string                     `xml:"policer,omitempty"`
	DontFragment    string                     `xml:"dont-fragment,omitempty"` // set | clear
	Next            string                     `xml:"next,omitempty"`          // term
	Count           string                     `xml:"count,omitempty"`         // counter name
	Discard         *FilterThenDiscard         `xml:"discard"`
	Reject          *FilterThenReject          `xml:"reject"`
	RoutingInstance *FilterThenRoutingInstance `xml:"routing-instance"`
	Accept          bool                       `xml:"accept,omitempty"`
	Log             bool                       `xml:"log,omitempty"`
	Syslog          bool                       `xml:"syslog,omitempty"`
	Sample          bool                       `xml:"sample,omitempty"`
}

type FilterThenDiscard struct {
	XMLName    xml.Name `xml:"discard"`
	Accounting string   `xml:"accounting,omitempty"`
}

type FilterThenRoutingInstance struct {
	XMLName             xml.Name `xml:"routing-instance"`
	RoutingInstanceName string   `xml:"routing-instance-name"`
}

type FilterThenReject struct {
	XMLName xml.Name `xml:"reject"`
	// Select one bool
	AdministrativelyProhibited bool `xml:"administratively-prohibited,omitempty"`
	BadHostTos                 bool `xml:"bad-host-tos,omitempty"`
	BadNetworkTos              bool `xml:"bad-network-tos,omitempty"`
	FragmentationNeeded        bool `xml:"fragmentation-needed,omitempty"`
	HostProhibited             bool `xml:"host-prohibited,omitempty"`
	HostUnknown                bool `xml:"host-unknown,omitempty"`
	HostUnreachable            bool `xml:"host-unreachable,omitempty"`
	NetworkProhibited          bool `xml:"network-prohibited,omitempty"`
	NetworkUnknown             bool `xml:"network-unknown,omitempty"`
	NetworkUnreachable         bool `xml:"network-unreachable,omitempty"`
	PortUnreachable            bool `xml:"port-unreachable,omitempty"`
	PrecedenceCutoff           bool `xml:"precedence-cutoff,omitempty"`
	PrecedenceViolation        bool `xml:"precedence-violation,omitempty"`
	ProtocolUnreachable        bool `xml:"protocol-unreachable,omitempty"`
	SourceHostIsolated         bool `xml:"source-host-isolated,omitempty"`
	SourceRouteFailed          bool `xml:"source-route-failed,omitempty"`
	TCPReset                   bool `xml:"tcp-reset,omitempty"`
}

type Policer struct {
	XMLName xml.Name `xml:"policer"`
	Name    string   `xml:"name"`
}

type FilterFamily struct {
	XMLName xml.Name           `xml:"family"`
	Inet    *FilterFamilyInet  `xml:"inet"`
	Inet6   *FilterFamilyInet6 `xml:"inet6"`
}

type FilterFamilyInet struct {
	XMLName xml.Name `xml:"inet"`
	Filters []Filter `xml:"filter,omitempty"`
}

type FilterFamilyInet6 struct {
	XMLName xml.Name `xml:"inet6"`
	Filters []Filter `xml:"filter,omitempty"`
}

func (filter *Filter) ConvertToConfig(family string) *Config {
	var conf Config

	switch family {
	case "inet":
		conf = Config{
			Configuration: &Configuration{
				Firewall: &Firewall{
					Family: &FilterFamily{
						Inet: &FilterFamilyInet{
							Filters: []Filter{*filter},
						},
					},
				},
			},
		}
	case "inet6":
		conf = Config{
			Configuration: &Configuration{
				Firewall: &Firewall{
					Family: &FilterFamily{
						Inet6: &FilterFamilyInet6{
							Filters: []Filter{*filter},
						},
					},
				},
			},
		}
	default:
		conf = Config{
			Configuration: &Configuration{
				Firewall: &Firewall{
					Filters: []Filter{*filter},
				},
			},
		}
	}

	return &conf
}
