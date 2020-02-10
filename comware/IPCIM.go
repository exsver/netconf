package comware

import (
	"github.com/exsver/netconf/netconf"
)

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

func (targetDevice *TargetDevice) AddIPSourceBinding(IfIndex string, Ipv4Address string, MacAddress string, VLANID string) error {
	binding := SourceBinding{
		IfIndex:     IfIndex,
		Ipv4Address: Ipv4Address,
		MacAddress:  MacAddress,
		VLANID:      VLANID,
	}
	return targetDevice.Configure(*binding.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) DeleteIPSourceBinding(IfIndex string, Ipv4Address string, MacAddress string, VLANID string) error {
	binding := SourceBinding{
		IfIndex:     IfIndex,
		Ipv4Address: Ipv4Address,
		MacAddress:  MacAddress,
		VLANID:      VLANID,
	}
	return targetDevice.Configure(*binding.ConvertToTop(), "remove")
}
