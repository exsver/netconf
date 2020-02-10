package comware

import "github.com/exsver/netconf/netconf"

func (targetDevice *TargetDevice) GetDataIRF() (*IRF, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><IRF/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}
	return data.Top.IRF, nil
}

//CLI equivalent "irf-port-configuration active"
func (targetDevice *TargetDevice) IRFPortConfigurationActive() error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
            <action>
              <top xmlns="http://www.hp.com/netconf/action:1.0">
                <IRF>
                  <PortConfiguration>
                    <Activate/>
                  </PortConfiguration>
                </IRF>
              </top>
            </action>`),
		Xmlns: []string{netconf.BaseURI},
	}
	err := targetDevice.PerformAction(request)
	return err
}
