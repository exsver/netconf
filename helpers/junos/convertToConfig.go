package junos

func (vlan *Vlan) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			Vlans: &Vlans{
				Vlans: []Vlan{*vlan},
			},
		},
	}
}
