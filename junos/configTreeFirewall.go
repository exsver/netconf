package junos

import "encoding/xml"

type Firewall struct {
	XMLName  xml.Name      `xml:"firewall"`
	Filters  []Filter      `xml:"filter,omitempty"`
	Policers []Policer     `xml:"policer,omitempty"`
	Family   *FilterFamily `xml:"family,omitempty"`
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
	XMLName xml.Name      `xml:"inet6"`
	Filters []FilterInet6 `xml:"filter,omitempty"`
}

type Filter struct {
	XMLName                xml.Name `xml:"filter"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Terms                  []Term   `xml:"term"`
}

type FilterInet6 struct {
	XMLName                xml.Name    `xml:"filter"`
	NetconfConfigOperation string      `xml:"operation,attr,omitempty"`
	Name                   string      `xml:"name"`
	Terms                  []TermInet6 `xml:"term"`
}

type Term struct {
	XMLName                   xml.Name    `xml:"term"`
	NetconfConfigOperation    string      `xml:"operation,attr,omitempty"`
	NetconfInsertPosition     string      `xml:"insert,attr,omitempty"` // "first" | "after" | "before"
	NetconfInsertPositionName string      `xml:"name,attr,omitempty"`   // referent-value for InsertPosition
	Name                      string      `xml:"name"`
	Filter                    string      `xml:"filter,omitempty"` // filter to include
	From                      *FilterFrom `xml:"from"`             // match criteria
	Then                      *FilterThen `xml:"then"`             // action
}

type TermInet6 struct {
	XMLName                   xml.Name         `xml:"term"`
	NetconfConfigOperation    string           `xml:"operation,attr,omitempty"`
	NetconfInsertPosition     string           `xml:"insert,attr,omitempty"` // "first" | "after" | "before"
	NetconfInsertPositionName string           `xml:"name,attr,omitempty"`   // referent-value for InsertPosition
	Name                      string           `xml:"name"`
	Filter                    string           `xml:"filter,omitempty"` // filter to include
	From                      *FilterFromInet6 `xml:"from"`             // match criteria
	Then                      *FilterThenInet6 `xml:"then"`             // action
}

type FilterFrom struct {
	XMLName                xml.Name                      `xml:"from"`
	NetconfConfigOperation string                        `xml:"operation,attr,omitempty"`
	Interfaces             []FilterInterface             `xml:"interface,omitempty"`
	Address                []FilterAddress               `xml:"address,omitempty"`
	SourceAddress          []FilterSourceAddress         `xml:"source-address,omitempty"`
	DestinationAddress     []FilterDestinationAddress    `xml:"destination-address,omitempty"`
	PrefixList             []FilterPrefixList            `xml:"prefix-list,omitempty"`
	SourcePrefixList       []FilterSourcePrefixList      `xml:"source-prefix-list,omitempty"`
	DestinationPrefixList  []FilterDestinationPrefixList `xml:"destination-prefix-list,omitempty"`
	Protocol               []string                      `xml:"protocol,omitempty"`
	ProtocolExcept         []string                      `xml:"protocol-except,omitempty"`
	Port                   []string                      `xml:"port,omitempty"`
	PortExcept             []string                      `xml:"port-except,omitempty"`
	SourcePort             []string                      `xml:"source-port,omitempty"`
	SourcePortExcept       []string                      `xml:"source-port-except,omitempty"`
	DestinationPort        []string                      `xml:"destination-port,omitempty"`
	DestinationPortExcept  []string                      `xml:"destination-port-except,omitempty"`
	TTL                    []string                      `xml:"ttl,omitempty"`
	TTLExcept              []string                      `xml:"ttl-except,omitempty"`
	ICMPType               []string                      `xml:"icmp-type,omitempty"`
	ICMPTypeExcept         []string                      `xml:"icmp-type-except,omitempty"`
	ICMPCode               []string                      `xml:"icmp-code,omitempty"`
	ICMPCodeExcept         []string                      `xml:"icmp-code-except,omitempty"`
	IPOptions              []string                      `xml:"ip-options,omitempty"`
	IPOptionsExcept        []string                      `xml:"ip-options-except,omitempty"`
	DSCP                   []string                      `xml:"dscp,omitempty"`
	DSCPExcept             []string                      `xml:"dscp-except,omitempty"`
	SourceClass            []string                      `xml:"source-class,omitempty"`
	SourceClassExcept      []string                      `xml:"source-class-except,omitempty"`
	DestinationClass       []string                      `xml:"destination-class,omitempty"`
	DestinationClassExcept []string                      `xml:"destination-class-except,omitempty"`
	ForwardingClass        []string                      `xml:"forwarding-class,omitempty"`
	ForwardingClassExcept  []string                      `xml:"forwarding-class-except,omitempty"`
	TCPFlags               string                        `xml:"tcp-flags,omitempty"`
	IsFragment             bool                          `xml:"is-fragment,omitempty"`
	FirstFragment          bool                          `xml:"first-fragment,omitempty"`
	TCPInitial             bool                          `xml:"tcp-initial,omitempty"`
	TCPEstablished         bool                          `xml:"tcp-established,omitempty"`
}

type FilterFromInet6 struct {
	XMLName                xml.Name                      `xml:"from"`
	NetconfConfigOperation string                        `xml:"operation,attr,omitempty"`
	Interfaces             []FilterInterface             `xml:"interface,omitempty"`
	Address                []FilterAddress               `xml:"address,omitempty"`
	SourceAddress          []FilterSourceAddress         `xml:"source-address,omitempty"`
	DestinationAddress     []FilterDestinationAddress    `xml:"destination-address,omitempty"`
	PrefixList             []FilterPrefixList            `xml:"prefix-list,omitempty"`
	SourcePrefixList       []FilterSourcePrefixList      `xml:"source-prefix-list,omitempty"`
	DestinationPrefixList  []FilterDestinationPrefixList `xml:"destination-prefix-list,omitempty"`
	NextHeader             []string                      `xml:"next-header,omitempty"`
	NextHeaderExcept       []string                      `xml:"next-header-except,omitempty"`
	ExtensionHeader        []string                      `xml:"extension-header,omitempty"`
	ExtensionHeaderExcept  []string                      `xml:"extension-header-except,omitempty"`
	Port                   []string                      `xml:"port,omitempty"`
	PortExcept             []string                      `xml:"port-except,omitempty"`
	SourcePort             []string                      `xml:"source-port,omitempty"`
	SourcePortExcept       []string                      `xml:"source-port-except,omitempty"`
	DestinationPort        []string                      `xml:"destination-port,omitempty"`
	DestinationPortExcept  []string                      `xml:"destination-port-except,omitempty"`
	HopLimit               []string                      `xml:"hop-limit,omitempty"`
	HopLimitExcept         []string                      `xml:"hop-limit-except,omitempty"`
	ICMPType               []string                      `xml:"icmp-type,omitempty"`
	ICMPTypeExcept         []string                      `xml:"icmp-type-except,omitempty"`
	ICMPCode               []string                      `xml:"icmp-code,omitempty"`
	ICMPCodeExcept         []string                      `xml:"icmp-code-except,omitempty"`
	SourceClass            []string                      `xml:"source-class,omitempty"`
	SourceClassExcept      []string                      `xml:"source-class-except,omitempty"`
	DestinationClass       []string                      `xml:"destination-class,omitempty"`
	DestinationClassExcept []string                      `xml:"destination-class-except,omitempty"`
	ForwardingClass        []string                      `xml:"forwarding-class,omitempty"`
	ForwardingClassExcept  []string                      `xml:"forwarding-class-except,omitempty"`
	TrafficClass           []string                      `xml:"traffic-class,omitempty"`
	TrafficClassExcept     []string                      `xml:"traffic-class-except,omitempty"`
	TCPFlags               string                        `xml:"tcp-flags,omitempty"`
	TCPInitial             bool                          `xml:"tcp-initial,omitempty"`
	TCPEstablished         bool                          `xml:"tcp-established,omitempty"`
}

type FilterThen struct {
	XMLName                xml.Name                   `xml:"then"`
	NetconfConfigOperation string                     `xml:"operation,attr,omitempty"`
	Policer                string                     `xml:"policer,omitempty"`
	DontFragment           string                     `xml:"dont-fragment,omitempty"` // set | clear
	Next                   string                     `xml:"next,omitempty"`          // term
	Count                  string                     `xml:"count,omitempty"`         // counter name
	Discard                *FilterThenDiscard         `xml:"discard"`
	Reject                 *FilterThenReject          `xml:"reject"`
	RoutingInstance        *FilterThenRoutingInstance `xml:"routing-instance"`
	Accept                 bool                       `xml:"accept,omitempty"`
	Log                    bool                       `xml:"log,omitempty"`
	Syslog                 bool                       `xml:"syslog,omitempty"`
	Sample                 bool                       `xml:"sample,omitempty"`
}

type FilterThenInet6 struct {
	XMLName                xml.Name                   `xml:"then"`
	NetconfConfigOperation string                     `xml:"operation,attr,omitempty"`
	Policer                string                     `xml:"policer,omitempty"`
	Next                   string                     `xml:"next,omitempty"`  // term
	Count                  string                     `xml:"count,omitempty"` // counter name
	Reject                 *FilterThenRejectInet6     `xml:"reject"`
	RoutingInstance        *FilterThenRoutingInstance `xml:"routing-instance"`
	Accept                 bool                       `xml:"accept,omitempty"`
	Discard                bool                       `xml:"discard,omitempty"`
	Log                    bool                       `xml:"log,omitempty"`
	Syslog                 bool                       `xml:"syslog,omitempty"`
	Sample                 bool                       `xml:"sample,omitempty"`
}

type FilterAddress struct {
	XMLName                xml.Name `xml:"address"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Except                 bool     `xml:"except,omitempty"`
}

type FilterSourceAddress struct {
	XMLName                xml.Name `xml:"source-address"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Except                 bool     `xml:"except,omitempty"`
}

type FilterDestinationAddress struct {
	XMLName                xml.Name `xml:"destination-address"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Except                 bool     `xml:"except,omitempty"`
}

type FilterInterface struct {
	XMLName                xml.Name `xml:"interface"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
}

type FilterPrefixList struct {
	XMLName                xml.Name `xml:"prefix-list"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Except                 bool     `xml:"except,omitempty"`
}

type FilterSourcePrefixList struct {
	XMLName                xml.Name `xml:"source-prefix-list"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Except                 bool     `xml:"except,omitempty"`
}

type FilterDestinationPrefixList struct {
	XMLName                xml.Name `xml:"destination-prefix-list"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Except                 bool     `xml:"except,omitempty"`
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

type FilterThenRejectInet6 struct {
	XMLName xml.Name `xml:"reject"`
	// Select one bool
	AddressUnreachable         bool `xml:"address-unreachable,omitempty"`
	AdministrativelyProhibited bool `xml:"administratively-prohibited,omitempty"`
	BeyondScope                bool `xml:"beyond-scope,omitempty"`
	NoRoute                    bool `xml:"no-route,omitempty"`
	PortUnreachable            bool `xml:"port-unreachable,omitempty"`
	TCPReset                   bool `xml:"tcp-reset,omitempty"`
}

type Policer struct {
	XMLName                xml.Name `xml:"policer"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
}

func (filter *Filter) ConvertToConfig(isFamilyInet bool) *Config {
	if isFamilyInet {
		return &Config{
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
	}

	return &Config{
		Configuration: &Configuration{
			Firewall: &Firewall{
				Filters: []Filter{*filter},
			},
		},
	}
}

func (filter *FilterInet6) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			Firewall: &Firewall{
				Family: &FilterFamily{
					Inet6: &FilterFamilyInet6{
						Filters: []FilterInet6{*filter},
					},
				},
			},
		},
	}
}

// family: ""|"inet"
func (term *Term) ConvertToConfig(isFamilyInet bool, filterName string) *Config {
	if isFamilyInet {
		return &Config{
			Configuration: &Configuration{
				Firewall: &Firewall{
					Family: &FilterFamily{
						Inet: &FilterFamilyInet{
							Filters: []Filter{
								{
									Name:  filterName,
									Terms: []Term{*term},
								},
							},
						},
					},
				},
			},
		}
	}

	return &Config{
		Configuration: &Configuration{
			Firewall: &Firewall{
				Filters: []Filter{
					{
						Name:  filterName,
						Terms: []Term{*term},
					},
				},
			},
		},
	}
}

func (term *TermInet6) ConvertToConfig(filterName string) *Config {
	return &Config{
		Configuration: &Configuration{
			Firewall: &Firewall{
				Family: &FilterFamily{
					Inet6: &FilterFamilyInet6{
						Filters: []FilterInet6{
							{
								Name:  filterName,
								Terms: []TermInet6{*term},
							},
						},
					},
				},
			},
		},
	}
}

func (filter *Filter) AppendToConfig(isFamilyInet bool, config *Config) *Config {
	if config.Configuration == nil {
		config.Configuration = &Configuration{}
	}

	if config.Configuration.Firewall == nil {
		config.Configuration.Firewall = &Firewall{}
	}

	if isFamilyInet {
		if config.Configuration.Firewall.Family == nil {
			config.Configuration.Firewall.Family = &FilterFamily{}
		}

		if config.Configuration.Firewall.Family.Inet == nil {
			config.Configuration.Firewall.Family.Inet = &FilterFamilyInet{}
		}

		if config.Configuration.Firewall.Family.Inet.Filters == nil {
			config.Configuration.Firewall.Family.Inet.Filters = make([]Filter, 0, 1)
		}

		config.Configuration.Firewall.Family.Inet.Filters = append(config.Configuration.Firewall.Family.Inet.Filters, *filter)

		return config
	}

	if config.Configuration.Firewall.Filters == nil {
		config.Configuration.Firewall.Filters = make([]Filter, 0, 1)
	}

	config.Configuration.Firewall.Filters = append(config.Configuration.Firewall.Filters, *filter)

	return config
}
