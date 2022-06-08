package comware

import (
	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) GetDataIPCIM() (*IPCIM, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><IPCIM/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.IPCIM, nil
}

func (targetDevice *TargetDevice) GetIPSourceBindings() ([]SourceBinding, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
         <get-config>
            <source><running/></source>
            <filter type="subtree">
              <top xmlns="http://www.hp.com/netconf/config:1.0">
                <IPCIM><IpSourceBindingInterface/></IPCIM>
              </top>
            </filter>
          </get-config>`),
		Xmlns: []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.IPCIM.IPSourceBindingInterface.SourceBindings, nil
}

func (targetDevice *TargetDevice) AddIPArpFilterSource(ifIndex int, ipv4Addresses []string) error {
	var filterSources []FilterSource

	for _, address := range ipv4Addresses {
		filterSources = append(filterSources, FilterSource{
			IfIndex:     ifIndex,
			Ipv4Address: address,
		})
	}

	arpFilters := ArpFilterSource{
		FilterSources: filterSources,
	}

	return targetDevice.Configure(*arpFilters.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) AddIPSourceBinding(ifIndex string, ipv4Address string, macAddress string, vlanID string) error {
	if vlanID == "" {
		vlanID = "0"
	}

	binding := SourceBinding{
		IfIndex:     ifIndex,
		Ipv4Address: ipv4Address,
		MacAddress:  macAddress,
		VLANID:      vlanID,
	}

	return targetDevice.Configure(*binding.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) DeleteIPSourceBinding(ifIndex string, ipv4Address string, macAddress string, vlanID string) error {
	binding := SourceBinding{
		IfIndex:     ifIndex,
		Ipv4Address: ipv4Address,
		MacAddress:  macAddress,
		VLANID:      vlanID,
	}

	return targetDevice.Configure(*binding.ConvertToTop(), "remove")
}

func (targetDevice *TargetDevice) AddIpVerifySource(ifIndex int, verifyIP bool, verifyMac bool) error {
	ipSourceVerify := VerifySource{
		IfIndex:          ifIndex,
		VerifyIPAddress:  verifyIP,
		VerifyMacAddress: verifyMac,
	}

	return targetDevice.Configure(*ipSourceVerify.ConvertToTop(), "merge")
}
