package comware

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) GetDataVLAN() (VLAN, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
        <get>
          <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0"><VLAN/></top>
          </filter>
        </get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return VLAN{}, err
	}

	var data Data

	err = xml.Unmarshal(rpcReply.Content, &data)
	if err != nil {
		return VLAN{}, err
	}

	return *data.Top.VLAN, nil
}

func (targetDevice *TargetDevice) SetAccessInterfaceVlan(ifIndex int, vlanID int) error {
	if vlanID < 1 || vlanID > 4094 {
		return fmt.Errorf("invalid vlan id. Valid values are between 1-4094 ")
	}

	accessInterface := AccessInterface{
		IfIndex: ifIndex,
		PVID:    vlanID,
	}

	return targetDevice.Configure(*accessInterface.ConvertToTop(), "replace")
}

// pvid = 0 -- mean no pvid
func (targetDevice *TargetDevice) SetTrunkInterfaceVlans(ifIndex int, permitVlanList []int, pvid int) error {
	permitVlanListString := ""
	for _, v := range permitVlanList {
		permitVlanListString = fmt.Sprintf("%s%s,", permitVlanListString, strconv.Itoa(v))
	}

	permitVlanListString = permitVlanListString[:len(permitVlanListString)-1]
	trunkInterface := TrunkInterface{
		IfIndex:        ifIndex,
		PermitVlanList: permitVlanListString,
	}

	if pvid != 0 {
		trunkInterface.PVID = pvid
	}

	return targetDevice.Configure(*trunkInterface.ConvertToTop(), "replace")
}

func (targetDevice *TargetDevice) IsVlanExist(vlanID int) (bool, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
        <get>
          <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
              <VLAN>
                <VLANs>
                  <VLANID>
                    <ID/>
                  </VLANID>
                </VLANs>
              </VLAN>
            </top>
          </filter>
        </get>`),
		Xmlns: []string{netconf.BaseURI},
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte("<ID/>"), append([]byte("<ID>"), append([]byte(strconv.Itoa(vlanID)), []byte("</ID>")...)...), 1)

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return false, err
	}

	if data.Top != nil {
		return true, nil
	}

	return false, nil
}

func (targetDevice *TargetDevice) IsVlanInTrunk(ifName string, vlanID int) (bool, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`
        <get>
          <filter type="subtree">
            <top xmlns="http://www.hp.com/netconf/data:1.0">
              <VLAN>
                <Interfaces>
                  <Interface>
                    <Name/>
                    <LinkType/>
                    <PermitVlanList/>
                  </Interface>
                </Interfaces>
              </VLAN>
            </top>
          </filter>
        </get>`),
		Xmlns: []string{netconf.BaseURI},
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("<Name/>"), append([]byte("<Name>"), append([]byte(ifName), []byte("</Name>")...)...), 1)

	data, err := targetDevice.RetrieveData(request)
	if err != nil {
		return false, err
	}

	if data.Top == nil {
		return false, fmt.Errorf("interface with name %s doesn't exist on the system", ifName)
	}

	iface := data.Top.VLAN.Interfaces.Interfaces[0]
	if iface.LinkType != 2 {
		return false, fmt.Errorf("interface with name %s isn't trunk", ifName)
	}

	vlans, err := (VlanListToIntSlice(iface.PermitVlanList))
	if err != nil {
		return false, err
	}

	for _, v := range vlans {
		if v == vlanID {
			return true, nil
		}
	}

	return false, nil
}

func VlanListToIntSlice(vlanList string) (list []int, err error) {
	if vlanList == "" {
		return
	}

	vlanStrings := strings.Split(vlanList, ",")

	for _, v := range vlanStrings {
		vs := strings.Split(v, "-")
		switch len(vs) {
		case 1:
			i, err := strconv.Atoi(vs[0])
			if err != nil {
				return list, fmt.Errorf("error while trying to convert a string into a int slice: %s", err.Error())
			}

			list = append(list, i)
		case 2:
			min, err := strconv.Atoi(vs[0])
			if err != nil {
				return list, fmt.Errorf("error while trying to convert a string into a int slice: %s", err.Error())
			}

			max, err := strconv.Atoi(vs[1])
			if err != nil {
				return list, fmt.Errorf("error while trying to convert a string into a int slice: %s", err.Error())
			}

			if min > max {
				return nil, fmt.Errorf("error while trying to convert a string into a int slice. min>max: %s", v)
			}

			for i := min; i <= max; i++ {
				list = append(list, i)
			}
		default:
			return nil, fmt.Errorf("error while trying to convert a string into a int slice. invalid value %s", v)
		}
	}

	return
}
