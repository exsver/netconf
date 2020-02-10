package comware

import (
	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) GetDataARP() (*ARP, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><ARP/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.ARP, nil
}

//CLI equivalent
// true - "arp rate-limit log enable"
// false - "undo arp rate-limit log enable"
func (targetDevice *TargetDevice) ARPRateLimitLogEnable(enable bool) error {
	arpRateLimitLog := ArpRateLimitLog{
		LogEnable: enable,
	}

	return targetDevice.Configure(*arpRateLimitLog.ConvertToTop(), "merge")
}
