package comware

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
