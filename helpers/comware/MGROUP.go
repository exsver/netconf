package comware

import "github.com/exsver/netconf"

func (targetDevice *TargetDevice) GetDataMGROUP() (*MGROUP, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><MGROUP/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}
	return data.Top.MGROUP, nil
}

func (targetDevice *TargetDevice) NewMirroringGroupLocal(id, monitorPortIfIndex int, sourcePorts []PortMirroringSourcePort) error {
	// Set or replace MirroringGroup ID
	for i, _ := range sourcePorts {
		sourcePorts[i].ID = id
	}

	conf := MGROUP{
		Groups: &MirrorGroups{
			MirrorGroups: []MirrorGroup{
				{
					ID:   id,
					Type: 1,
				},
			},
		},
		SourcePorts: &PortMirroringSourcePorts{
			SourcePorts: sourcePorts,
		},
		MonitorPort: &PortMirroringMonitorPorts{
			MonitorPorts: []PortMirroringMonitorPort{
				{
					ID:   id,
					Port: monitorPortIfIndex,
				},
			},
		},
	}

	return targetDevice.Configure(*conf.ConvertToTop(), "create")
}

func (targetDevice *TargetDevice) NewMirroringGroupRemoteSource(id, egressPortIfIndex, remoteProbeVlan int, sourcePorts []PortMirroringSourcePort) error {
	// Set or replace MirroringGroup ID
	for i, _ := range sourcePorts {
		sourcePorts[i].ID = id
	}

	conf := MGROUP{
		Groups: &MirrorGroups{
			MirrorGroups: []MirrorGroup{
				{
					ID:   id,
					Type: 2,
				},
			},
		},
		SourcePorts: &PortMirroringSourcePorts{
			SourcePorts: sourcePorts,
		},
		EgressPort: &PortMirroringEgressPorts{
			EgressPorts: []PortMirroringEgressPort{
				{
					ID:   id,
					Port: egressPortIfIndex,
				},
			},
		},
		ProbeVlan: &PortMirroringProbeVlans{
			ProbeVlans: []PortMirroringProbeVlan{
				{
					ID:     id,
					VlanID: remoteProbeVlan,
				},
			},
		},
	}

	return targetDevice.Configure(*conf.ConvertToTop(), "create")
}

func (targetDevice *TargetDevice) NewMirroringGroupRemoteDestination(id, monitorPortIfIndex, remoteProbeVlan int) error {

	conf := MGROUP{
		Groups: &MirrorGroups{
			MirrorGroups: []MirrorGroup{
				{
					ID:   id,
					Type: 3,
				},
			},
		},
		MonitorPort: &PortMirroringMonitorPorts{
			MonitorPorts: []PortMirroringMonitorPort{
				{
					ID:   id,
					Port: monitorPortIfIndex,
				},
			},
		},
		ProbeVlan: &PortMirroringProbeVlans{
			ProbeVlans: []PortMirroringProbeVlan{
				{
					ID:     id,
					VlanID: remoteProbeVlan,
				},
			},
		},
	}

	return targetDevice.Configure(*conf.ConvertToTop(), "create")
}

func (targetDevice *TargetDevice) MirroringGroupRemove(id int) error {
	mirrorGroup := MirrorGroup{
		ID: id,
	}

	return targetDevice.Configure(*mirrorGroup.ConvertToTop(), "remove")
}
