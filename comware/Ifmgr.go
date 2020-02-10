package comware

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) GetDataIfmgr() (*Ifmgr, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Ifmgr/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Ifmgr, nil
}

//GetPortsList returns а list of all physical interfaces ([]Port), Bridge-aggregation interfaces, and Management interfaces. Exclude SVI.
func (targetDevice *TargetDevice) GetPorts() ([]Port, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Ifmgr><Ports/></Ifmgr></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Ifmgr.Ports.Ports, nil
}

//RegExp Examples:
// "^GigabitEthernet"       -- all GigabitEthernet ports
// "^Ten-GigabitEthernet"   -- all Ten-GigabitEthernet ports
func (targetDevice *TargetDevice) GetPortsRegExp(regExp string) ([]Port, error) {
	request := netconf.RPCMessage{
		InnerXML:    []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Ifmgr><Ports><Port><Name nc:regExp="*"/><IfIndex/></Port></Ports></Ifmgr></top></filter></get>`),
		Xmlns:       []string{netconf.BaseURI},
		CustomAttrs: []string{`xmlns:nc="http://www.hp.com/netconf/base:1.0"`},
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte(`nc:regExp="*"`), append([]byte(`nc:regExp="`), append([]byte(regExp), []byte(`"`)...)...), 1)

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Ifmgr.Ports.Ports, nil
}

//Filters examples:
// all items -                    nil
// all BAGG interfaces -          []string{`<ifType>161</ifType>`}
// all ethernet Interfaces -      []string{`<ifType>6</ifType>`}
// all Vlan-interfaces -          []string{`<ifType>136</ifType>`}
// Port with ifIndex 10 -         []string{`<IfIndex>10</IfIndex>`}
// Port with index 10 -           []string{`<PortIndex>10</PortIndex>`}
// All Ports in Up state -        []string{`<OperStatus>1</OperStatus>`}
// Port with Description "test" - []string{`<Description>test</Description>`}
// all ethernet Interfaces in UP
// state -                        []string{`<ifType>6</ifType>`, `<OperStatus>1</OperStatus>`}
func (targetDevice *TargetDevice) GetInterfacesInfo(filters []string) ([]Interface, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Ifmgr><Interfaces/></Ifmgr></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	if filters != nil {
		request.InnerXML = []byte(
			`<get>
          <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
              <Ifmgr>
                <Interfaces>
                  <Interface>
                    <IfIndex/>
                    <Name/>
                    <AbbreviatedName/>
                    <PortIndex/>
                    <ifTypeExt/>
                    <ifType/>
                    <Description/>
                    <AdminStatus/>
                    <OperStatus/>
                    <ConfigSpeed/>
                    <ActualSpeed/>
                    <ConfigDuplex/>
                    <ActualDuplex/>
                    <PortLayer/>
                    <InetAddressIPV4/>
                    <InetAddressIPV4Mask/>
                    <LinkType/>
                    <PVID/>
                    <PhysicalIndex/>
                    <MAC/>
                    <ForwardingAttributes/>
                    <ConfigMTU/>
                    <ActualMTU/>
                    <Loopback/>
                    <MDI/>
                    <ActualBandwidth/>
                    <SubPort/>
                    <ForceUP/>
                  </Interface>
                </Interfaces>
              </Ifmgr>
            </top>
          </filter>
        </get>`)
		for _, filter := range filters {
			request.InnerXML = bytes.Replace(request.InnerXML, convertToEmptyTag([]byte(filter)), []byte(filter), 1)
		}
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Ifmgr.Interfaces.Interfaces, nil
}

type IfIdentity struct {
	Name            string
	AbbreviatedName string
	Description     string
}

func (targetDevice *TargetDevice) GetIfIdentity() (map[int]IfIdentity, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
        <get>
          <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
              <Ifmgr>
                <Interfaces>
                  <Interface>
                    <IfIndex/>
                    <Name/>
                    <AbbreviatedName/>
                    <Description/>
                  </Interface>
                </Interfaces>
              </Ifmgr>
            </top>
          </filter>
        </get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	ifIdentity := make(map[int]IfIdentity)

	if data.Top != nil {
		for _, v := range data.Top.Ifmgr.Interfaces.Interfaces {
			ifIdentity[v.IfIndex] = IfIdentity{
				Name:            v.Name,
				AbbreviatedName: v.AbbreviatedName,
				Description:     v.Description,
			}
		}
	}

	return ifIdentity, nil
}

//CLI equivalent "port link-type [access|trunk|hybrid]"
func (targetDevice *TargetDevice) SetInterfaceLinkType(ifIndex int, linkType string) error {
	iface := Interface{IfIndex: ifIndex}

	switch linkType {
	case "access", "Access", "ACCESS", "A", "1":
		iface.LinkType = 1
	case "trunk", "Trunk", "TRUNK", "T", "2":
		iface.LinkType = 2
	case "hybrid", "Hybrid", "HYBRID", "H", "3":
		iface.LinkType = 3
	default:
		return fmt.Errorf(`invalid linkType string. Correct values are: "access", "Access", "ACCESS", "A", "1", "trunk", "Trunk", "TRUNK", "T", "2", "hybrid", "Hybrid", "HYBRID", "H", "3"`)
	}

	return targetDevice.Configure(*iface.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) SetInterfaceSpeed(ifIndex int, linkSpeed string) error {
	iface := Interface{IfIndex: ifIndex}

	switch linkSpeed {
	case "auto", "Auto", "AUTO":
		iface.ConfigSpeed = 1
	case "10":
		iface.ConfigSpeed = 2
	case "100":
		iface.ConfigSpeed = 4
	case "1000", "1G":
		iface.ConfigSpeed = 32
	case "10000", "10G":
		iface.ConfigSpeed = 1024
	case "40000", "40G":
		iface.ConfigSpeed = 8192
	case "100000", "100G":
		iface.ConfigSpeed = 16384
	default:
		return fmt.Errorf(`invalid linkSpeed string. Correct values are: "10", "100", "1000", "1G", "10000", "10G", "40000", "40G", "100000", "100G" ,"auto", "Auto", "AUTO"`)
	}

	return targetDevice.Configure(*iface.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) SetInterfaceDesription(ifIndex int, description string) error {
	iface := Interface{
		IfIndex:     ifIndex,
		Description: description,
	}

	return targetDevice.Configure(*iface.ConvertToTop(), "merge")
}

//RestoreInterfaceConfiguration sets following port settings to Default values:
//Description
//AdminStatus
//Speed/Duplex/MDI
//BPDUDrop
//FlowControl
//Broadcast/Multicast/UnknownUnicast Suppression
//Jumboframe
func (targetDevice *TargetDevice) RestoreInterfaceConfiguration(ifIndex int) error {
	iface := &Ifmgr{
		EthInterfaces: &EthInterfaces{
			[]EthInterface{
				{IfIndex: ifIndex},
			},
		},
		Interfaces: &Interfaces{
			[]Interface{
				{IfIndex: ifIndex},
			},
		},
	}

	return targetDevice.Configure(*iface.ConvertToTop(), "replace")
}

//RestoreInterfaceDefaultConfiguration sets all port settings to default values.
//CLI equivalent "default" in interface view
func (targetDevice *TargetDevice) RestoreInterfaceDefaultConfiguration(ifIndex int) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
<action>
  <top xmlns="http://www.hp.com/netconf/action:1.0">
    <Ifmgr>
      <Interfaces>
        <Interface>
          <IfIndex/>
          <Default/>
        </Interface>
      </Interfaces>
    </Ifmgr>
  </top>
</action>`),
		Xmlns: []string{netconf.BaseURI},
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("<IfIndex/>"), append([]byte("<IfIndex>"), append([]byte(strconv.Itoa(ifIndex)), []byte("</IfIndex>")...)...), 1)
	err := targetDevice.PerformAction(request)

	return err
}

func (targetDevice *TargetDevice) GetTrafficStatistics() ([]InterfaceTrafficStatistics, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Ifmgr><TrafficStatistics/></Ifmgr></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Ifmgr.TrafficStatistics.TrafficStatistics.Interfaces, nil
}

func (targetDevice *TargetDevice) GetInterfaceStatistics() ([]InterfaceStatistics, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Ifmgr><Statistics/></Ifmgr></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Ifmgr.Statistics.Interfaces, nil
}

func (targetDevice *TargetDevice) GetEthPortStatistics() ([]InterfaceEthPortStatistics, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Ifmgr><EthPortStatistics/></Ifmgr></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Ifmgr.EthPortStatistics.Interfaces, nil
}

//ClearAllIfStatistics clear statistics on all interfaces
func (targetDevice *TargetDevice) ClearAllIfStatistics() error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
<action>
  <top xmlns="http://www.hp.com/netconf/action:1.0">
    <Ifmgr>
      <ClearAllIfStatistics>
        <Clear/>
      </ClearAllIfStatistics>
    </Ifmgr>
  </top>
</action>`),
		Xmlns: []string{netconf.BaseURI},
	}
	err := targetDevice.PerformAction(request)

	return err
}

func (targetDevice *TargetDevice) ClearIfStatistics(ifIndex int) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
<action>
  <top xmlns="http://www.hp.com/netconf/action:1.0">
    <Ifmgr>
      <Interfaces>
        <Interface>
          <IfIndex/>
          <Clear/>
        </Interface>
      </Interfaces>
    </Ifmgr>
  </top>
</action>`),
		Xmlns: []string{netconf.BaseURI},
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("<IfIndex/>"), append([]byte("<IfIndex>"), append([]byte(strconv.Itoa(ifIndex)), []byte("</IfIndex>")...)...), 1)
	err := targetDevice.PerformAction(request)

	return err
}

func (targetDevice *TargetDevice) GetIfIndexes(filters []XMLFilter) (ifIndexes []int, err error) {
	filter := convertFiltersToXML(filters)

	request := netconf.RPCMessage{
		InnerXML: []byte(`
        <get>
          <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
              <Ifmgr>
                <Interfaces>
                  <Interface>
                    <filters/>
                    <IfIndex/>
                  </Interface>
                </Interfaces>
              </Ifmgr>
            </top>
          </filter>
        </get>`),
		Xmlns: []string{netconf.BaseURI},
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("<filters/>"), filter, 1)

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return ifIndexes, err
	}

	if data.Top != nil {
		for _, v := range data.Top.Ifmgr.Interfaces.Interfaces {
			ifIndexes = append(ifIndexes, v.IfIndex)
		}
	}

	return ifIndexes, nil
}

func (targetDevice *TargetDevice) GetIfIndexesByName(ifName string, isRegExp bool) ([]int, error) {
	var filters []XMLFilter
	if isRegExp {
		filters = []XMLFilter{
			{
				Key:      "Name",
				Value:    ifName,
				IsRegExp: true,
			},
		}
	} else {
		filters = []XMLFilter{
			{
				Key:   "Name",
				Value: ifName,
			},
		}
	}

	return targetDevice.GetIfIndexes(filters)
}

func (targetDevice *TargetDevice) GetIfIndexesByDecription(description string, isRegExp bool) (ifIndexes []int, err error) {
	var filters []XMLFilter
	if isRegExp {
		filters = []XMLFilter{
			{
				Key:      "Description",
				Value:    description,
				IsRegExp: true,
			},
		}
	} else {
		filters = []XMLFilter{
			{
				Key:   "Description",
				Value: description,
			},
		}
	}

	return targetDevice.GetIfIndexes(filters)
}
