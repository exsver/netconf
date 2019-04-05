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

func (acl *Group) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			Groups: &Groups{
				Groups: []Group{*acl},
			},
		},
	}
}

func (acl *NamedGroup) ConvertToTop() *Top {
	return &Top{
		ACL: &ACL{
			NamedGroups: &NamedGroups{
				Groups: []NamedGroup{*acl},
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
