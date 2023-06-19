package comware

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
