package comware

import "github.com/exsver/netconf/netconf"

func (targetDevice *TargetDevice) GetDataSTP() (*STP, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><STP/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.STP, nil
}

func (targetDevice *TargetDevice) SetSTPBaseParams(base STPBase) error {
	return targetDevice.Configure(*base.ConvertToTop(), "merge")
}
