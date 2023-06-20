package comware

func (vlans *VLANs) ConvertToTop() *Top {
	return &Top{
		VLAN: &VLAN{
			VLANs: vlans,
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
