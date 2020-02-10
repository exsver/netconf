package junos

import (
	"encoding/xml"

	"github.com/exsver/netconf/netconf"
)

type PhysicalInterfaces struct {
	XMLName                xml.Name            `xml:"interface-information"`
	PhysicalInterfacesList []PhysicalInterface `xml:"physical-interface"`
}

type PhysicalInterface struct {
	Name              string             `xml:"name"`
	AdminStatus       string             `xml:"admin-status"`      // up
	OperStatus        string             `xml:"oper-status"`       // up|down
	LocalIndex        int                `xml:"local-index"`       //
	SNMPIndex         int                `xml:"snmp-index"`        //
	Description       string             `xml:"description"`       //
	IfType            string             `xml:"if-type"`           // Ethernet|Software-Pseudo|GRE|Mgmt-VLAN
	LinkLevelType     string             `xml:"link-level-type"`   // Ethernet|VxLAN-Tunnel-Endpoint|Interface-Specific|Mgmt-VLAN
	MTU               string             `xml:"mtu"`               // Unlimited|1500|1514|1600|9512
	Speed             string             `xml:"speed"`             // Auto|800mbps|1000mbps|10Gbps|Unlimited
	LinkType          string             `xml:"link-type"`         // Full-Duplex|
	IfMediaType       string             `xml:"if-media-type"`     // Fiber
	InterfaceFlapped  string             `xml:"interface-flapped"` // 2018-05-22 19:15:36 UTC (1d 04:25 ago) | Never
	TrafficStatistics TrafficStatistics  `xml:"traffic-statistics"`
	LogicalInterfaces []LogicalInterface `xml:"logical-interface"`
}

type LogicalInterface struct {
	Name              string            `xml:"name"`
	LocalIndex        int               `xml:"local-index"`
	SNMPIndex         int               `xml:"snmp-index"`
	Encapsulation     string            `xml:"encapsulation"`
	TrafficStatistics TrafficStatistics `xml:"traffic-statistics"`
}

type TrafficStatistics struct {
	InputBPS      int64 `xml:"input-bps"`
	InputPPS      int   `xml:"input-pps"`
	OutputBPS     int64 `xml:"output-bps"`
	OutputPPS     int   `xml:"output-pps"`
	InputPackets  int   `xml:"input-packets"`
	OutputPackets int   `xml:"output-packets"`
}

func (targetDevice *TargetDevice) GetInterfaceInformation() ([]PhysicalInterface, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get-interface-information/>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return []PhysicalInterface{}, err
	}

	if rpcReply.Error() != nil {
		return []PhysicalInterface{}, rpcReply.Error()
	}

	rpcReply.Content = netconf.Normalize(rpcReply.Content)

	var interfacesList PhysicalInterfaces

	err = xml.Unmarshal(rpcReply.Content, &interfacesList)

	return interfacesList.PhysicalInterfacesList, err
}

type FiberInterfaces struct {
	XMLName          xml.Name            `xml:"interface-information"`
	Name             string              `xml:"name"`
	InterfacesOptics []OpticsDiagnostics `xml:"physical-interface>optics-diagnostics"`
}

type OpticsDiagnostics struct {
	ModuleTemperature                   string                        `xml:"module-temperature"`                      // QSFP|XFP
	ModuleVoltage                       string                        `xml:"module-voltage"`                          // QSFP| -
	ModuleNotReadyAlarm                 string                        `xml:"module-not-ready-alarm"`                  //   - |XFP
	ModulePowerDownAlarm                string                        `xml:"module-power-down-alarm"`                 //   - |XFP
	ModuleVoltageHighAlarm              string                        `xml:"module-voltage-high-alarm"`               // QSFP| -
	ModuleVoltageLowAlarm               string                        `xml:"module-voltage-low-alarm"`                // QSFP| -
	ModuleVoltageHighWarn               string                        `xml:"module-voltage-high-warn"`                // QSFP| -
	ModuleVoltageLowWarn                string                        `xml:"module-voltage-low-warn"`                 // QSFP| -
	ModuleTemperatureHighAlarm          string                        `xml:"module-temperature-high-alarm"`           // QSFP|XFP
	ModuleTemperatureHighAlarmThreshold string                        `xml:"module-temperature-high-alarm-threshold"` // QSFP|XFP
	ModuleTemperatureLowAlarm           string                        `xml:"module-temperature-low-alarm"`            // QSFP|XFP
	ModuleTemperatureLowAlarmThreshold  string                        `xml:"module-temperature-low-alarm-threshold"`  // QSFP|XFP
	ModuleTemperatureHighWarn           string                        `xml:"module-temperature-high-warn"`            // QSFP|XFP
	ModuleTemperatureHighWarnThreshold  string                        `xml:"module-temperature-high-warn-threshold"`  // QSFP|XFP
	ModuleTemperatureLowWarn            string                        `xml:"module-temperature-low-warn"`             // QSFP|XFP
	ModuleTemperatureLowWarnThreshold   string                        `xml:"module-temperature-low-warn-threshold"`   // QSFP|XFP
	ModuleVoltageHighAlarmThreshold     string                        `xml:"module-voltage-high-alarm-threshold"`     // QSFP| -
	ModuleVoltageLowAlarmThreshold      string                        `xml:"module-voltage-low-alarm-threshold"`      // QSFP| -
	ModuleVoltageHighWarnThreshold      string                        `xml:"module-voltage-high-warn-threshold"`      // QSFP| -
	ModuleVoltageLowWarnThreshold       string                        `xml:"module-voltage-low-warn-threshold"`       // QSFP| -
	LaserBiasCurrent                    string                        `xml:"laser-bias-current"`                      //   - |XFP
	LaserOutputPower                    string                        `xml:"laser-output-power"`                      //   - |XFP
	LaserOutputPowerDbm                 string                        `xml:"laser-output-power-dbm"`                  //   - |XFP
	LaserBiasCurrentHighAlarm           string                        `xml:"laser-bias-current-high-alarm"`           //   - |XFP
	LaserBiasCurrentHighAlarmThreshold  string                        `xml:"laser-bias-current-high-alarm-threshold"` // QSFP|XFP
	LaserBiasCurrentLowAlarm            string                        `xml:"laser-bias-current-low-alarm"`            //   - |XFP
	LaserBiasCurrentLowAlarmThreshold   string                        `xml:"laser-bias-current-low-alarm-threshold"`  // QSFP|XFP
	LaserBiasCurrentHighWarn            string                        `xml:"laser-bias-current-high-warn"`            //   - |XFP
	LaserBiasCurrentHighWarnThreshold   string                        `xml:"laser-bias-current-high-warn-threshold"`  // QSFP|XFP
	LaserBiasCurrentLowWarn             string                        `xml:"laser-bias-current-low-warn"`             //   - |XFP
	LaserBiasCurrentLowWarnThreshold    string                        `xml:"laser-bias-current-low-warn-threshold"`   // QSFP|XFP
	LaserTXPowerHighAlarm               string                        `xml:"laser-tx-power-high-alarm"`               //   - |XFP
	LaserTXPowerHighAlarmThreshold      string                        `xml:"laser-tx-power-high-alarm-threshold"`     // QSFP|XFP
	LaserTXPowerHighAlarmThresholdDbm   string                        `xml:"laser-tx-power-high-alarm-threshold-dbm"` // QSFP|XFP
	LaserTXPowerLowAlarm                string                        `xml:"laser-tx-power-low-alarm"`                //   - |XFP
	LaserTXPowerLowAlarmThreshold       string                        `xml:"laser-tx-power-low-alarm-threshold"`      // QSFP|XFP
	LaserTXPowerLowAlarmThresholdDbm    string                        `xml:"laser-tx-power-low-alarm-threshold-dbm"`  // QSFP|XFP
	LaserTXPowerHighWarn                string                        `xml:"laser-tx-power-high-warn"`                //   - |XFP
	LaserTXPowerHighWarnThreshold       string                        `xml:"laser-tx-power-high-warn-threshold"`      // QSFP|XFP
	LaserTXPowerHighWarnThresholdDbm    string                        `xml:"laser-tx-power-high-warn-threshold-dbm"`  // QSFP|XFP
	LaserTXPowerLowWarn                 string                        `xml:"laser-tx-power-low-warn"`                 //   - |XFP
	LaserTXPowerLowWarnThreshold        string                        `xml:"laser-tx-power-low-warn-threshold"`       // QSFP|XFP
	LaserTXPowerLowWarnThresholdDbm     string                        `xml:"laser-tx-power-low-warn-threshold-dbm"`   // QSFP|XFP
	LaserRXOpticalPower                 string                        `xml:"laser-rx-optical-power"`                  //   - |XFP
	LaserRXOpticalPowerDbm              string                        `xml:"laser-rx-optical-power-dbm"`              //   - |XFP
	LaserRXPowerHighAlarm               string                        `xml:"laser-rx-power-high-alarm"`               //   - |XFP
	LaserRXPowerHighAlarmThreshold      string                        `xml:"laser-rx-power-high-alarm-threshold"`     // QSFP|XFP
	LaserRXPowerHighAlarmThresholdDbm   string                        `xml:"laser-rx-power-high-alarm-threshold-dbm"` // QSFP|XFP
	LaserRXPowerLowAlarm                string                        `xml:"laser-rx-power-low-alarm"`                //   - |XFP
	LaserRXPowerLowAlarmThreshold       string                        `xml:"laser-rx-power-low-alarm-threshold"`      // QSFP|XFP
	LaserRXPowerLowAlarmThresholdDbm    string                        `xml:"laser-rx-power-low-alarm-threshold-dbm"`  // QSFP|XFP
	LaserRXPowerHighWarn                string                        `xml:"laser-rx-power-high-warn"`                //   - |XFP
	LaserRXPowerHighWarnThreshold       string                        `xml:"laser-rx-power-high-warn-threshold"`      // QSFP|XFP
	LaserRXPowerHighWarnThresholdDbm    string                        `xml:"laser-rx-power-high-warn-threshold-dbm"`  // QSFP|XFP
	LaserRXPowerLowWarn                 string                        `xml:"laser-rx-power-low-warn"`                 //   - |XFP
	LaserRXPowerLowWarnThreshold        string                        `xml:"laser-rx-power-low-warn-threshold"`       // QSFP|XFP
	LaserRXPowerLowWarnThresholdDbm     string                        `xml:"laser-rx-power-low-warn-threshold-dbm"`   // QSFP|XFP
	TXDataNotReadyAlarm                 string                        `xml:"tx-data-not-ready-alarm"`                 //   - |XFP
	TXNotReadyAlarm                     string                        `xml:"tx-not-ready-alarm"`                      //   - |XFP
	TXLaserFaultAlarm                   string                        `xml:"tx-laser-fault-alarm"`                    //   - |XFP
	TXCDRLossOfLockAlarm                string                        `xml:"tx-cdr-loss-of-lock-alarm"`               //   - |XFP
	RXNotReadyAlarm                     string                        `xml:"rx-not-ready-alarm"`                      //   - |XFP
	RXLossOfSignalAlarm                 string                        `xml:"rx-loss-of-signal-alarm"`                 //   - |XFP
	RXCDRLossOfLockAlarm                string                        `xml:"rx-cdr-loss-of-lock-alarm"`               //   - |XFP
	OpticsDiagnosticsLaneValues         []OpticsDiagnosticsLaneValues `xml:"optics-diagnostics-lane-values"`          // QSFP| -
}

type OpticsDiagnosticsLaneValues struct {
	LaneIndex                        int    `xml:"lane-index"`
	LaserBiasCurrent                 string `xml:"laser-bias-current"`
	LaserOutputPower                 string `xml:"laser-output-power"`
	LaserOutputPowerDbm              string `xml:"laser-output-power-dbm"`
	LaserRXPpticalPower              string `xml:"laser-rx-optical-power"`
	LaserRXOpticalPowerDbm           string `xml:"laser-rx-optical-power-dbm"`
	LaserBiasCurrentHighAlarm        string `xml:"laser-bias-current-high-alarm"`
	LaserBiasCurrentLowAlarm         string `xml:"laser-bias-current-low-alarm"`
	LaserBiasCurrentHighWarn         string `xml:"laser-bias-current-high-warn"`
	LaserBiasCurrentLowWarn          string `xml:"laser-bias-current-low-warn"`
	LaserRXPowerHighAlarm            string `xml:"laser-rx-power-high-alarm"`
	LaserRXPowerLowAlarm             string `xml:"laser-rx-power-low-alarm"`
	LaserRXPowerHighWarn             string `xml:"laser-rx-power-high-warn"`
	LaserRXPowerLowWarn              string `xml:"laser-rx-power-low-warn"`
	TXLossOfSignalFunctionalityAlarm string `xml:"tx-loss-of-signal-functionality-alarm"`
	RXLossOfSignalAlarm              string `xml:"rx-loss-of-signal-alarm"`
	TXLaserDisabledAlarm             string `xml:"tx-laser-disabled-alarm"`
}

func (targetDevice *TargetDevice) GetDiagnosticsOptics() ([]OpticsDiagnostics, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get-interface-optics-diagnostics-information/>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return []OpticsDiagnostics{}, err
	}

	if rpcReply.Error() != nil {
		return []OpticsDiagnostics{}, rpcReply.Error()
	}

	rpcReply.Content = netconf.Normalize(rpcReply.Content)

	var interfacesList FiberInterfaces

	err = xml.Unmarshal(rpcReply.Content, &interfacesList)

	return interfacesList.InterfacesOptics, err
}
