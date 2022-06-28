package comware

import (
	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) GetDataSyslog() (*Syslog, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Syslog/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.Syslog, nil
}

func (targetDevice *TargetDevice) ClearLogBuffer() error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<action><top xmlns="http://www.hp.com/netconf/action:1.0"><Syslog><LogBuffer><Clear/></LogBuffer></Syslog></top></action>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	return targetDevice.PerformAction(request)
}
