package comware

import "github.com/exsver/netconf/netconf"

func (targetDevice *TargetDevice) GetDataDHCP() (*DHCP, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><DHCP/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.DHCP, nil
}

func (targetDevice *TargetDevice) GetDataDHCPServerIpPools() (*DHCPServerIPPool, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><DHCP><DHCPServerIpPool/></DHCP></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.DHCP.DHCPServerIPPool, nil
}

func (targetDevice *TargetDevice) GetDataDHCPServerIpInUse() (*DHCPServerIPInUse, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><DHCP><DHCPServerIpInUse/></DHCP></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.DHCP.DHCPServerIPInUse, nil
}
