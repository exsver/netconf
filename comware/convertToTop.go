package comware

func (base *Base) ConvertToTop() *Top {
	return &Top{
		Device: &Device{
			Base: base,
		},
	}
}

func (iface *Ifmgr) ConvertToTop() *Top {
	return &Top{
		Ifmgr: iface,
	}
}

func (iface *Interface) ConvertToTop() *Top {
	return &Top{
		Ifmgr: &Ifmgr{
			Interfaces: &Interfaces{
				[]Interface{*iface},
			},
		},
	}
}

func (iface *EthInterface) ConvertToTop() *Top {
	return &Top{
		Ifmgr: &Ifmgr{
			EthInterfaces: &EthInterfaces{
				[]EthInterface{*iface},
			},
		},
	}
}

func (port *Port) ConvertToTop() *Top {
	return &Top{
		Ifmgr: &Ifmgr{
			Ports: &Ports{
				Ports: []Port{*port},
			},
		},
	}
}

func (binding *SourceBinding) ConvertToTop() *Top {
	return &Top{
		IPCIM: &IPCIM{
			IPSourceBindingInterface: &IPSourceBindingInterface{
				SourceBindings: []SourceBinding{*binding},
			},
		},
	}
}

func (verifySource *VerifySource) ConvertToTop() *Top {
	return &Top{
		IPCIM: &IPCIM{
			IPVerifySource: &IPVerifySource{
				[]VerifySource{*verifySource},
			},
		},
	}
}

func (iface *AccessInterface) ConvertToTop() *Top {
	return &Top{
		VLAN: &VLAN{
			AccessInterfaces: &AccessInterfaces{
				AccessInterfaces: []AccessInterface{*iface},
			},
		},
	}
}

func (iface *TrunkInterface) ConvertToTop() *Top {
	return &Top{
		VLAN: &VLAN{
			TrunkInterfaces: &TrunkInterfaces{
				TrunkInterfaces: []TrunkInterface{*iface},
			},
		},
	}
}

func (iface *HybridInterface) ConvertToTop() *Top {
	return &Top{
		VLAN: &VLAN{
			HybridInterfaces: &HybridInterfaces{
				HybridInterfaces: []HybridInterface{*iface},
			},
		},
	}
}

func (mGroup *MGROUP) ConvertToTop() *Top {
	return &Top{
		MGROUP: mGroup,
	}
}

func (mirrorGroup *MirrorGroup) ConvertToTop() *Top {
	return &Top{
		MGROUP: &MGROUP{
			Groups: &MirrorGroups{
				MirrorGroups: []MirrorGroup{*mirrorGroup},
			},
		},
	}
}

func (monitorPort *PortMirroringMonitorPort) ConvertToTop() *Top {
	return &Top{
		MGROUP: &MGROUP{
			MonitorPort: &PortMirroringMonitorPorts{
				MonitorPorts: []PortMirroringMonitorPort{*monitorPort},
			},
		},
	}
}

func (egressPort *PortMirroringEgressPort) ConvertToTop() *Top {
	return &Top{
		MGROUP: &MGROUP{
			EgressPort: &PortMirroringEgressPorts{
				EgressPorts: []PortMirroringEgressPort{*egressPort},
			},
		},
	}
}

func (probeVlan *PortMirroringProbeVlan) ConvertToTop() *Top {
	return &Top{
		MGROUP: &MGROUP{
			ProbeVlan: &PortMirroringProbeVlans{
				ProbeVlans: []PortMirroringProbeVlan{*probeVlan},
			},
		},
	}
}

func (acl *Group) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			Groups: &Groups{
				Groups: []Group{*acl},
			},
		},
	}
}

func (namedGroup *NamedGroup) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			NamedGroups: &NamedGroups{
				Groups: []NamedGroup{*namedGroup},
			},
		},
	}
}

func (rules *IPv4NamedAdvanceRules) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv4NamedAdvanceRules: rules,
		},
	}
}

func (rule *IPv4NamedAdvanceRule) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv4NamedAdvanceRules: &IPv4NamedAdvanceRules{
				IPv4NamedAdvanceRules: []IPv4NamedAdvanceRule{*rule},
			},
		},
	}
}

func (rule *IPv4NamedBasicRule) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			IPv4NamedBasicRules: &IPv4NamedBasicRules{
				IPv4NamedBasicRules: []IPv4NamedBasicRule{*rule},
			},
		},
	}
}

func (pfilter *Pfilter) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			PfilterApply: &PfilterApply{
				Pfilters: []Pfilter{*pfilter},
			},
		},
	}
}

func (filterSource *FilterSource) ConvertToTop() *Top {
	return &Top{
		ARP: &ARP{
			ArpFilterSource: &ArpFilterSource{
				FilterSources: []FilterSource{*filterSource},
			},
		},
	}
}

func (arpRateLimitLog *ArpRateLimitLog) ConvertToTop() *Top {
	return &Top{
		ARP: &ARP{
			ArpRateLimitLog: arpRateLimitLog,
		},
	}
}

func (dhcp *DHCP) ConvertToTop() *Top {
	return &Top{
		DHCP: dhcp,
	}
}

func (dhcpConfig *DHCPConfig) ConvertToTop() *Top {
	return &Top{
		DHCP: &DHCP{
			DHCPConfig: dhcpConfig,
		},
	}
}

func (dhcpIfMode *DHCPIfMode) ConvertToTop() *Top {
	return &Top{
		DHCP: &DHCP{
			DHCPIfMode: dhcpIfMode,
		},
	}
}

func (dhcpServerIPPool *DHCPServerIPPool) ConvertToTop() *Top {
	return &Top{
		DHCP: &DHCP{
			DHCPServerIPPool: dhcpServerIPPool,
		},
	}
}

func (dhcpServerPoolStatic *DHCPServerPoolStatic) ConvertToTop() *Top {
	return &Top{
		DHCP: &DHCP{
			DHCPServerPoolStatic: dhcpServerPoolStatic,
		},
	}
}

func (portSecurity *PortSecurity) ConvertToTop() *Top {
	return &Top{
		PortSecurity: portSecurity,
	}
}

func (portSecurityCommon *PortSecurityCommon) ConvertToTop() *Top {
	return &Top{
		PortSecurity: &PortSecurity{
			Common: portSecurityCommon,
		},
	}
}

func (iface *PortSecurityInterface) ConvertToTop() *Top {
	return &Top{
		PortSecurity: &PortSecurity{
			Interfaces: &PortSecurityInterfaces{
				[]PortSecurityInterface{*iface},
			},
		},
	}
}

func (lagg *LAGG) ConvertToTop() *Top {
	return &Top{
		LAGG: lagg,
	}
}

func (base *LAGGBase) ConvertToTop() *Top {
	return &Top{
		LAGG: &LAGG{
			Base: base,
		},
	}
}

func (member *LAGGMember) ConvertToTop() *Top {
	return &Top{
		LAGG: &LAGG{
			LAGGMembers: &LAGGMembers{
				[]LAGGMember{*member},
			},
		},
	}
}

func (route *IPv4RouteEntry) ConvertToTop() *Top {
	return &Top{
		StaticRoute: &StaticRoute{
			Ipv4StaticRouteConfigurations: &Ipv4StaticRouteConfigurations{
				[]IPv4RouteEntry{*route},
			},
		},
	}
}

func (route *IPv6RouteEntry) ConvertToTop() *Top {
	return &Top{
		StaticRoute: &StaticRoute{
			Ipv6StaticRouteConfigurations: &Ipv6StaticRouteConfigurations{
				[]IPv6RouteEntry{*route},
			},
		},
	}
}
