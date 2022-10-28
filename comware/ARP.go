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

func (targetDevice *TargetDevice) GetARPRateLimitInterfaces() ([]ArpRateLimitInterface, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><ARP><ArpRateLimit><RateLimit/></ArpRateLimit></ARP></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.ARP.ArpRateLimit.RateLimitInterfaces, nil
}

// CLI equivalent
//  true - "arp rate-limit log enable"
//  false - "undo arp rate-limit log enable"
func (targetDevice *TargetDevice) ARPRateLimitLogEnable(enable bool) error {
	arpRateLimitLog := ArpRateLimitLog{
		LogEnable: enable,
	}

	return targetDevice.Configure(*arpRateLimitLog.ConvertToTop(), "merge")
}

func (targetDevice *TargetDevice) SetArpFilterBinding(ifIndex string, ipv4Address string, macAddress string) error {
	interfaceArpFilter := ArpInterfaceFilter{
		IfIndex:     ifIndex,
		Ipv4Address: ipv4Address,
		MacAddress:  macAddress,
	}

	return targetDevice.Configure(*interfaceArpFilter.ConvertToTop(), "merge")
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
