package comware

import "encoding/xml"

// Port mirroring configuration
type MGROUP struct {
	/* top level
	   MGROUP
		 Capabilities            ***ReadOnly***
	     Groups
	       []Group
	     MonitorPort
	       []Group
		 ReflectorPort
	       []Group
	     EgressPort
	       []Group
		 ProbeVlan
	       []Group
	     SourcePorts
	       []SourcePort
	*/
	Capabilities  *PortMirroringCapabilities   `xml:"Capabilities"`
	Groups        *MirrorGroups                `xml:"Groups"`
	MonitorPort   *PortMirroringMonitorPorts   `xml:"MonitorPort"`
	ReflectorPort *PortMirroringReflectorPorts `xml:"ReflectorPort"`
	EgressPort    *PortMirroringEgressPorts    `xml:"EgressPort"`
	ProbeVlan     *PortMirroringProbeVlans     `xml:"ProbeVlan"`
	SourcePorts   *PortMirroringSourcePorts    `xml:"SourcePorts"`
}

// MirrorGroups table contains information about mirroring groups.
type MirrorGroups struct {
	XMLName      xml.Name      `xml:"Groups"`
	MirrorGroups []MirrorGroup `xml:"Group"`
}

type MirrorGroup struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`             // Valid values are: 1-256
	Type    int      `xml:"Type,omitempty"` // 1 - Local, 2 - Remote-source, 3 - Remote-destination
	Status  int      `xml:"Status,omitempty"`
}

// PortMirroringMonitorPorts table contains information about monitor port of mirroring groups.
type PortMirroringMonitorPorts struct {
	XMLName      xml.Name                   `xml:"MonitorPort"`
	MonitorPorts []PortMirroringMonitorPort `xml:"Group"`
}

type PortMirroringMonitorPort struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`
	Port    int      `xml:"Port"`
}

// PortMirroringReflectorPorts table contains information about reflector port of mirroring groups
type PortMirroringReflectorPorts struct {
	XMLName        xml.Name                     `xml:"ReflectorPort"`
	ReflectorPorts []PortMirroringReflectorPort `xml:"Group"`
}

type PortMirroringReflectorPort struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`
	Port    int      `xml:"Port"`
}

// PortMirroringEgressPorts table contains information about egress port of mirroring groups.
type PortMirroringEgressPorts struct {
	XMLName     xml.Name                  `xml:"EgressPort"`
	EgressPorts []PortMirroringEgressPort `xml:"Group"`
}

type PortMirroringEgressPort struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`   // Mirroring group ID. Value range: 1 to 256
	Port    int      `xml:"Port"` // Index of egress	port.
}

// PortMirroringProbeVlans table contains information about remote probe VLAN of mirroring groups.
type PortMirroringProbeVlans struct {
	XMLName    xml.Name                 `xml:"ProbeVlan"`
	ProbeVlans []PortMirroringProbeVlan `xml:"Group"`
}

// PortMirroringProbeVlan Dedicated VLAN that sends packets from the source device to the destination device.
type PortMirroringProbeVlan struct {
	XMLName xml.Name `xml:"Group"`
	ID      int      `xml:"ID"`     // Mirroring group ID. Value range: 1 to 256
	VlanID  int      `xml:"VlanID"` // Probe VLAN ID
}

// PortMirroringSourcePorts table contains information about source ports of the mirroring group.
type PortMirroringSourcePorts struct {
	XMLName     xml.Name                  `xml:"SourcePorts"`
	SourcePorts []PortMirroringSourcePort `xml:"SourcePort"`
}

type PortMirroringSourcePort struct {
	XMLName   xml.Name `xml:"SourcePort"`
	ID        int      `xml:"ID"`                  // Mirroring group ID. Value range: 1 to 256
	IfIndex   int      `xml:"IfIndex,omitempty"`   // Interface index
	Direction int      `xml:"Direction,omitempty"` // Direction of source port: 1—Inbound, 2—Outbound, 3—Both.
}

type PortMirroringCapabilities struct {
	XMLName                xml.Name `xml:"Capabilities"`
	MaxGroupNum            int      `xml:"MaxGroupNum"`
	SourceType             int      `xml:"SourceType"`
	VlanTagMode            int      `xml:"VlanTagMode"`
	Sampler                int      `xml:"Sampler"`
	MultiMonitor           bool     `xml:"MultiMonitor"`
	MultiMonitorInOneGroup bool     `xml:"MultiMonitorInOneGroup"`
}
