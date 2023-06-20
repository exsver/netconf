package comware

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
