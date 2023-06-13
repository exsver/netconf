package comware

import "encoding/xml"

type Ifmgr struct {
	/* top level
	   Ifmgr
	     DataBuffer
	       IfBuffer
	         []Interface
	     DeviceCapabilities
	     EthInterfaceCapabilities
	       []Interface
	     InterfaceCapabilities
	       []Interface
	     TypeCapabilities
	       []Capability
	     EthInterfaces
	       []EthInterface
	     Interfaces
	       []Interface
	     PortIsolation
	       Capabilities
	       Interfaces
	         []Interface
	     Ports                                     ***ReadOnly***
	       []Port                                  ***ReadOnly***
	     PortIsolation
	       Capabilities
	       Interfaces
	         []Interface
	     StormConstrain
	       Interfaces
	         []Interface
	       Interval
	     Statistics                                ***ReadOnly***
	       []InterfaceStatistics                   ***ReadOnly***
	     TrafficStatistics                         ***ReadOnly***
	       InterfacesTrafficStatistics             ***ReadOnly***
	         []InterfaceTrafficStatistics          ***ReadOnly***
	       Interval                                ***ReadOnly***
	     EthPortStatistics                         ***ReadOnly***
	       []InterfacesEthPortStatistics           ***ReadOnly***
	*/
	InterfaceCapabilities *InterfacesCapabilities `xml:"InterfaceCapabilities"`
	EthInterfaces         *EthInterfaces          `xml:"EthInterfaces"`
	Interfaces            *Interfaces             `xml:"Interfaces"`
	Ports                 *Ports                  `xml:"Ports"`
	Statistics            *Statistics             `xml:"Statistics"`
	TrafficStatistics     *TrafficStatistics      `xml:"TrafficStatistics"`
	EthPortStatistics     *EthPortStatistics      `xml:"EthPortStatistics"`
}

// InterfacesCapabilities table contains port capability information.
type InterfacesCapabilities struct {
	Interfaces []InterfaceCapabilities `xml:"Interface"`
}

// ReadOnly
type InterfaceCapabilities struct {
	XMLName xml.Name `xml:"Interface"`
	IfIndex int      `xml:"IfIndex"`
	// Speed which can be configured.
	// Each bit represents a speed as follows:
	// 1:AuTo 2:10Mbps 4:100Mbps 8:155Mbps 16:622Mbps 32:1Gbps 64:2Gbps 128:2.5Gbps 256:4Gbps 512:8Gbps
	// 1024:10Gbps 2048:16Gbps 4096:20Gbps 8192:40Gbps 16384:100Gbps 32768:5Gbps 65536:25Gbps 131072:32Gbps
	//
	// For interface 10/100/1000 Speed = 39 (1:AuTo + 2:10Mbps + 4:100Mbps + 32:1Gbps)
	Speed        int  `xml:"Speed"`
	AutoSpeed    int  `xml:"AutoSpeed"`
	Duplex       int  `xml:"Duplex"`
	MDI          int  `xml:"MDI"`
	Configurable bool `xml:"Configurable"`
	Shutdown     bool `xml:"Shutdown"`
	Removable    bool `xml:"Removable"`
}

// EthInterfaces table contains Ethernet interfaces information.
type EthInterfaces struct {
	Interfaces []EthInterface `xml:"Interface"`
}

type EthInterface struct {
	XMLName                   xml.Name                    `xml:"Interface"`
	IfIndex                   int                         `xml:"IfIndex"`
	FlowControl               int                         `xml:"FlowControl,omitempty"`
	Jumboframe                int                         `xml:"Jumboframe,omitempty"`
	AutoPowerDown             bool                        `xml:"AutoPowerDown,omitempty"`
	BPDUDrop                  bool                        `xml:"BPDUDrop,omitempty"`
	EEE                       bool                        `xml:"EEE,omitempty"`
	BroadcastSuppression      []BroadcastSuppression      `xml:"BroadcastSuppression,omitempty"`
	MulticastSuppression      []MulticastSuppression      `xml:"MulticastSuppression,omitempty"`
	UnknownUnicastSuppression []UnknownUnicastSuppression `xml:"UnknownUnicastSuppression,omitempty"`
}

type BroadcastSuppression struct {
	ConfigValue int             `xml:"ConfigValue"`
	ActualValue int             `xml:"ActualValue,omitempty"`
	Unit        SuppressionUnit `xml:"Unit"`
}

type UnknownUnicastSuppression struct {
	ConfigValue int             `xml:"ConfigValue"`
	ActualValue int             `xml:"ActualValue,omitempty"`
	Unit        SuppressionUnit `xml:"Unit"`
}

type MulticastSuppression struct {
	ConfigValue int             `xml:"ConfigValue"`
	ActualValue int             `xml:"ActualValue,omitempty"`
	Unit        SuppressionUnit `xml:"Unit"`
	Flag        int             `xml:"Flag,omitempty"`
}

// Interfaces table contains basic interface information.
type Interfaces struct {
	Interfaces []Interface `xml:"Interface"`
}

type Interface struct {
	XMLName xml.Name `xml:"Interface"`
	// Name - Full name of an interface.
	// Examples:
	//  GigabitEthernet1/0/1
	//  Ten-GigabitEthernet1/0/25
	//  Bridge-Aggregation1
	//  Vlan-interface99
	Name string `xml:"Name,omitempty"`
	// AbbreviatedName - Abbreviated name of an interface.
	// Examples:
	//  GE1/0/1
	//  XGE1/0/25
	//  BAGG1
	//  Vlan99
	AbbreviatedName     string `xml:"AbbreviatedName,omitempty"`
	InetAddressIPV4     string `xml:"InetAddressIPV4,omitempty"`
	InetAddressIPV4Mask string `xml:"InetAddressIPV4Mask,omitempty"`
	// MAC - MAC address of an interface.
	MAC       string `xml:"MAC,omitempty"`
	IfIndex   int    `xml:"IfIndex"`
	PortIndex int    `xml:"PortIndex,omitempty"`
	IfTypeExt int    `xml:"ifTypeExt,omitempty"`
	// IfType - Interface type, according to IANAifType
	//  https://www.iana.org/assignments/ianaiftype-mib/ianaiftype-mib
	// Examples:
	//  6 - Physical ethernet Interface,
	//  24 - Loopback,
	//  136 - Vlan-interface
	//  161 - Bridge-Aggregation,
	IfType int `xml:"ifType,omitempty"`
	// Description - Interface description.
	// String length constraints must be in range(0..255).
	Description string `xml:"Description,omitempty"`
	// AdminStatus - Interface administration status
	//  1 - Admin Up
	//  2 - Admin Down
	AdminStatus InterfaceAdminStatus `xml:"AdminStatus,omitempty"`
	// OperStatus - Interface operation status
	//  1 - up
	//  2 - down
	//  3 - testing
	//  4 - unknown
	//  5 - dormant
	//  6 - notPresent
	//  7 - lowerLayerDown
	OperStatus InterfaceOperStatus `xml:"OperStatus,omitempty"`
	// Configured speed of an interface
	//  1 - auto				8 - 155Mbps
	//  2 - 10Mbps				16 - 622Mbps
	//  4 - 100Mbps             64 - 2Gbps
	//  32 - 1Gbps              128 - 2.5Gbps
	//  1024 - 10Gbps           256 - 4Gbps
	//  8192 - 40Gbps           512 - 8Gbps
	//  16384 - 100Gbps         2048 - 16Gbps
	//                          4096 - 20Gbps
	//                          32768 - 5Gbps
	//
	// Example: 37 - Auto-negotiation mode, and negotiation values are 100Mbps and 1000Mbps.
	ConfigSpeed int `xml:"ConfigSpeed,omitempty"`
	// Actual speed of an interface (units: kbps).
	ActualSpeed int `xml:"ActualSpeed,omitempty"`
	// Configured duplex mode of an interface:
	//  1 - full
	//  2 - half
	//  3 - auto
	ConfigDuplex InterfaceDuplex `xml:"ConfigDuplex,omitempty"`
	// Actual duplex mode of an interface:
	//  1 - full
	//  2 - half
	//  3 - auto
	ActualDuplex InterfaceDuplex `xml:"ActualDuplex,omitempty"`
	// PortLayer - Port layer of an interface as follows:
	//  1 - Layer 2
	//  2 - Layer 3
	PortLayer int `xml:"PortLayer,omitempty"`
	// LinkType - VLAN type of an interface:
	//  1 - Access
	//  2 - Trunk
	//  3 - Hybrid
	LinkType             InterfaceLinkType `xml:"LinkType,omitempty"`
	PVID                 int               `xml:"PVID,omitempty"`
	PhysicalIndex        int               `xml:"PhysicalIndex,omitempty"`
	ForwardingAttributes int               `xml:"ForwardingAttributes,omitempty"`
	ConfigMTU            int               `xml:"ConfigMTU,omitempty"`
	ActualMTU            int               `xml:"ActualMTU,omitempty"`
	Loopback             int               `xml:"Loopback,omitempty"`
	// MDI mode of an interface
	//  1 - MDI-II (straight-through cable)
	//  2 - MDI-X (crossover cable)
	//  3 - MDI-AUTO (auto-sensing)
	MDI               int  `xml:"MDI,omitempty"`
	ActualBandwidth   int  `xml:"ActualBandwidth,omitempty"`
	Interval          int  `xml:"Interval,omitempty"`          // absent in HP5130
	Actual64Bandwidth int  `xml:"Actual64Bandwidth,omitempty"` // absent in HP5130
	ForceUP           bool `xml:"ForceUP,omitempty"`
	SubPort           bool `xml:"SubPort,omitempty"`
}

type Ports struct {
	XMLName xml.Name `xml:"Ports"`
	Ports   []Port   `xml:"Port"`
}

type Port struct {
	XMLName   xml.Name `xml:"Port"`
	PortIndex int      `xml:"PortIndex,omitempty"`
	Name      string   `xml:"Name,omitempty"`
	IfIndex   int      `xml:"IfIndex,omitempty"`
}

// Statistics table contains interface statistics.
type Statistics struct {
	Interfaces []InterfaceStatistics `xml:"Interface"`
}

// ReadOnly struct
type InterfaceStatistics struct {
	XMLName         xml.Name `xml:"Interface"`
	IfIndex         int      `xml:"IfIndex"`
	Name            string   `xml:"Name"`
	AbbreviatedName string   `xml:"AbbreviatedName"`
	InOctets        uint64   `xml:"InOctets"`
	InUcastPkts     uint64   `xml:"InUcastPkts"`
	InNUcastPkts    uint64   `xml:"InNUcastPkts"`
	InDiscards      uint64   `xml:"InDiscards"`
	InErrors        uint64   `xml:"InErrors"`
	InUnknownProtos uint64   `xml:"InUnknownProtos"`
	InRate          uint64   `xml:"InRate"`
	OutOctets       uint64   `xml:"OutOctets"`
	OutUcastPkts    uint64   `xml:"OutUcastPkts"`
	OutNUcastPkts   uint64   `xml:"OutNUcastPkts"`
	OutDiscards     uint64   `xml:"OutDiscards"`
	OutErrors       uint64   `xml:"OutErrors"`
	OutRate         uint64   `xml:"OutRate"`
	// LastClear - Local time when the statistics on an interface were cleared most recently.
	//  0000-00-00T00:00:00 mean newer
	LastClear string `xml:"LastClear"`
}

type TrafficStatistics struct {
	TrafficStatistics *InterfacesTrafficStatistics `xml:"Interfaces"`
}

type InterfacesTrafficStatistics struct {
	Interfaces []InterfaceTrafficStatistics `xml:"Interface"`
}

// ReadOnly struct
type InterfaceTrafficStatistics struct {
	XMLName   xml.Name `xml:"Interface"`
	IfIndex   int      `xml:"IfIndex"`
	Name      string   `xml:"Name"`
	Interval  int      `xml:"Interval"`
	InPkts    uint64   `xml:"InPkts"`
	OutPkts   uint64   `xml:"OutPkts"`
	InOctets  uint64   `xml:"InOctets"`
	OutOctets uint64   `xml:"OutOctets"`
	InBits    uint64   `xml:"InBits"`
	OutBits   uint64   `xml:"OutBits"`
}

type EthPortStatistics struct {
	Interfaces []InterfaceEthPortStatistics `xml:"Interface"`
}

// ReadOnly struct
type InterfaceEthPortStatistics struct {
	IfIndex         int    `xml:"IfIndex"`
	Name            string `xml:"Name"`
	InBytes         uint64 `xml:"InBytes"`
	InPkts          uint64 `xml:"InPkts"`
	InUcastPkts     uint64 `xml:"InUcastPkts"`
	InBrdcastPkts   uint64 `xml:"InBrdcastPkts"`
	InMulticastPkts uint64 `xml:"InMulticastPkts"`
	// InPauses - Number of inbound pause frames on an interface.
	InPauses               uint64 `xml:"InPauses"`
	InNormalUnicastBytes   uint64 `xml:"InNormalUnicastBytes"`
	InNormalBrdcastBytes   uint64 `xml:"InNormalBrdcastBytes"`
	InNormalMulticastBytes uint64 `xml:"InNormalMulticastBytes"`
	InUnknownUnicastBytes  uint64 `xml:"InUnknownUnicastBytes"`
	InNormalPkts           uint64 `xml:"InNormalPkts"`
	InNormalUnicastPkts    uint64 `xml:"InNormalUnicastPkts"`
	InNormalBrdcastPkts    uint64 `xml:"InNormalBrdcastPkts"`
	InNormalMulticastPkts  uint64 `xml:"InNormalMulticastPkts"`
	InUnknownUnicastPkts   uint64 `xml:"InUnknownUnicastPkts"`
	InNormalPauses         uint64 `xml:"InNormalPauses"`
	InErrorPkts            uint64 `xml:"InErrorPkts"`
	// InPktSpeed - Rate of inbound packages on an interface.
	InPktSpeed  uint64 `xml:"InPktSpeed"`
	InByteSpeed uint64 `xml:"InByteSpeed"`
	InRunts     uint64 `xml:"InRunts"`
	InGiants    uint64 `xml:"InGiants"`
	// InThrottles  - Number of inbound frames that had a non-integer number of bytes
	InThrottles uint64 `xml:"InThrottles"`
	// InErrCRCFrames - Total number of inbound frames that had a normal length, but contained CRC errors.
	InErrCRCFrames uint64 `xml:"InErrCRCFrames"`
	// InErrFrames - Total number of inbound frames that contained CRC errors and a non-integer number of bytes.
	InErrFrames uint64 `xml:"InErrFrames"`
	// InAbortPkts - Total number of illegal inbound packets.
	InAbortPkts uint64 `xml:"InAbortPkts"`
	// InSpeedPeakBytes - Peak rate of inbound traffic in Bps.
	InSpeedPeakBytes uint64 `xml:"InSpeedPeakBytes"`
	// InSpeedPeakTime - The time when the peak inbound traffic rate occurred.
	InSpeedPeakTime  string `xml:"InSpeedPeakTime"`
	OutBytes         uint64 `xml:"OutBytes"`
	OutPkts          uint64 `xml:"OutPkts"`
	OutUcastPkts     uint64 `xml:"OutUcastPkts"`
	OutBrdcastPkts   uint64 `xml:"OutBrdcastPkts"`
	OutMulticastPkts uint64 `xml:"OutMulticastPkts"`
	// OutPauses - Number of outbound pause frames on an interface.
	OutPauses              uint64 `xml:"OutPauses"`
	OutNormalUnicastBytes  uint64 `xml:"OutNormalUnicastBytes"`
	OutNormalPkts          uint64 `xml:"OutNormalPkts"`
	OutNormalUnicastPkts   uint64 `xml:"OutNormalUnicastPkts"`
	OutNormalBrdcastPkts   uint64 `xml:"OutNormalBrdcastPkts"`
	OutNormalMulticastPkts uint64 `xml:"OutNormalMulticastPkts"`
	OutUnknownUnicastPkts  uint64 `xml:"OutUnknownUnicastPkts"`
	OutNormalPauses        uint64 `xml:"OutNormalPauses"`
	OutErrorPkts           uint64 `xml:"OutErrorPkts"`
	// OutPktSpeed - Rate of outbound packages on an interface.
	OutPktSpeed uint64 `xml:"OutPktSpeed"`
	// OutByteSpeed - Rate of outbound bytes on an interface.
	OutByteSpeed uint64 `xml:"OutByteSpeed"`
	// OutAbortPkts - Number of packets that failed to be transmitted, for example, because of Ethernet collisions.
	OutAbortPkts uint64 `xml:"OutAbortPkts"`
	// OutDeferedFrames - Number of frames that the interface deferred to transmit because of detected collisions.
	OutDeferedFrames uint64 `xml:"OutDeferedFrames"`
	// OutCollisionFrames - Number of frames that the interface stopped transmitting because Ethernet collisions were detected during transmission.
	OutCollisionFrames uint64 `xml:"OutCollisionFrames"`
	// OutLateCollisionFrames - Number of frames that the interface deferred to transmit after transmitting their first 512 bits because of detected collisions.
	OutLateCollisionFrames uint64 `xml:"OutLateCollisionFrames"`
	// OutLostCarriers - Number of carrier losses during transmission.
	OutLostCarriers   uint64 `xml:"OutLostCarriers"`
	OutSpeedPeakBytes uint64 `xml:"OutSpeedPeakBytes"`
	OutSpeedPeakTime  string `xml:"OutSpeedPeakTime"`
}

// InterfaceAdminStatus
//
//	1 - AdmUp,
//	2 - AdmDown,
type InterfaceAdminStatus int

func (status InterfaceAdminStatus) String() string {
	switch status {
	case InterfaceAdminStatusUP:
		return InterfaceAdminStatusUPString
	case InterfaceAdminStatusDown:
		return InterfaceAdminStatusDownString
	}

	return UnknownString
}

// InterfaceDuplex
//
//	1 - Full,
//	2 - Half,
//	3 - Auto
type InterfaceDuplex int

func (duplex InterfaceDuplex) String() string {
	switch duplex {
	case InterfaceDuplexFull:
		return InterfaceDuplexFullString
	case InterfaceDuplexHalf:
		return InterfaceDuplexHalfString
	case InterfaceDuplexAuto:
		return InterfaceDuplexAutoString
	}

	return UnknownString
}

// InterfaceOperStatus
//
//	1 - Up,
//	2 - Down,
//	3 - Testing,
//	4 - Unknown,
//	5 - Dormant,
//	6 - NotPresent,
//	7 - LowerLayerDown
type InterfaceOperStatus int

func (status InterfaceOperStatus) String() string {
	switch status {
	case InterfaceStatusUp:
		return InterfaceStatusUpString
	case InterfaceStatusDown:
		return InterfaceStatusDownString
	case InterfaceStatusTesting:
		return InterfaceStatusTestingString
	case InterfaceStatusUnknown:
		return InterfaceStatusUnknownString
	case InterfaceStatusDormant:
		return InterfaceStatusDormantString
	case InterfaceStatusNotPresent:
		return InterfaceStatusNotPresentString
	case InterfaceStatusLowerLayerDown:
		return InterfaceStatusLowerLayerDownString
	}

	return UnknownString
}

// SuppressionUnit
//
//	1 - ratio (suppression threshold in percentage),
//	2 - pps (suppression threshold in pps),
//	3 - kbps (suppression threshold in kbps).
type SuppressionUnit int

func (unit SuppressionUnit) String() string {
	switch unit {
	case SuppressionUnitRatio:
		return SuppressionUnitRatioString
	case SuppressionUnitPps:
		return SuppressionUnitPpsString
	case SuppressionUnitKbps:
		return SuppressionUnitKbpsString
	}

	return UnknownString
}
