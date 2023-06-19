package comware

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
