package comware

import (
	"bytes"
	"fmt"

	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) GetDataMAC() (*MAC, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><MAC/></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	return data.Top.MAC, nil
}

// GetMacTable returns mac-address table of the target device ([]MacTableEntry) or error.
// Available filters are:
//  - VLANID
//  - MacAddress
//  - PortIndex
//  - Status
//  - Aging
// Filter examples:
//  - all mac-addresses									   nil
//  - all mac-addresses with VLANID 99  --                 []XMLFilter{{Key: "VLANID", Value:"99", IsRegExp:false,},}
//  - all mac-addresses with VLANID 99 and PortIndex 1 --  []XMLFilter{{Key: "VLANID", Value:"99", IsRegExp:false,},{Key: "PortIndex", Value:"1", IsRegExp:false,},}
//  - all mac-addresses starts with "40-B0-34" --          []XMLFilter{{Key: "MacAddress", Value:"^40-B0-34", IsRegExp:true,},}
//  													   []XMLFilter{{Key: "PortIndex", Value:"^(719)$" IsRegExp:true,},}
func (targetDevice *TargetDevice) GetMacTable(filters []XMLFilter) ([]MacTableEntry, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
        <get>
          <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
              <MAC>
                <MacUnicastTable>
                  <Unicast>
                    <VLANID/>
                    <MacAddress/>
                    <PortIndex/>
                    <Status/>
                    <Aging/>
                  </Unicast>
                </MacUnicastTable>
              </MAC>
            </top>
          </filter>
        </get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	for _, filter := range filters {
		if !((filter.Key == "VLANID") || (filter.Key == "MacAddress") || (filter.Key == "PortIndex") || (filter.Key == "Status") || (filter.Key == "Aging")) {
			return nil, fmt.Errorf("invalid filter: %s. Valid filters are: VLANID,  MacAddress, PortIndex, Status, Aging", filter.Key)
		}

		request.InnerXML = bytes.Replace(request.InnerXML, []byte(fmt.Sprintf("<%s/>", filter.Key)), filter.ConvertToXML(), 1)
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return nil, err
	}

	// nothing found
	if data.Top == nil {
		return nil, nil
	}

	return data.Top.MAC.MacUnicastTable.Unicast, nil
}

func (targetDevice *TargetDevice) GetMacAgingTime() (*MacAging, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><MAC><MacAging/></MAC></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return &MacAging{}, err
	}

	return data.Top.MAC.MacAging, nil
}

func (targetDevice *TargetDevice) GetMacSpecification() (*MacSpecification, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><MAC><MacSpecification/></MAC></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return &MacSpecification{}, err
	}

	return data.Top.MAC.MacSpecification, nil
}
