package comware

import "github.com/exsver/netconf/netconf"

func (targetDevice *TargetDevice) GetDataPortSecurity() (*PortSecurity, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><PortSecurity/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.PortSecurity, nil
}

// CLI equivalent: "port-security enable"
func (targetDevice *TargetDevice) PortSecurityEnable(common *PortSecurityCommon) error {
	return targetDevice.Configure(*common.ConvertToTop(), "merge")
}
