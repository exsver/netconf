package junos

import (
	"encoding/xml"
	"github.com/exsver/netconf"
)

/*
Struct for cli Commands:
- show chassis hardware						junos:style="inventory"
- show chassis hardware clei-models			junos:style="clei-model-inventory"
*/
type ChassisInventory struct {
	XMLName xml.Name             `xml:"chassis-inventory"`
	Chassis ChassisInventoryItem `xml:"chassis"`
}

type ChassisInventoryItem struct {
	XMLName        xml.Name        `xml:"chassis"`
	Name           string          `xml:"name"`
	SerialNumber   string          `xml:"serial-number"`
	Description    string          `xml:"description"`
	ChassisModules []ChassisModule `xml:"chassis-module"`
}

type ChassisModule struct { //For example RE, CB, FPC, PEM and so on.
	Name             string             `xml:"name"`
	Version          string             `xml:"version"`
	PartNumber       string             `xml:"part-number"`
	SerialNumber     string             `xml:"serial-number"`
	Description      string             `xml:"description"`
	CleiCode         string             `xml:"clei-code"`
	ModelNumber      string             `xml:"model-number"`
	ChassisSubModule []ChassisSubModule `xml:"chassis-sub-module"`
}

type ChassisSubModule struct { //For example PIC
	Name                 string                `xml:"name"`
	Version              string                `xml:"version"`
	PartNumber           string                `xml:"part-number"`
	SerialNumber         string                `xml:"serial-number"`
	Description          string                `xml:"description"`
	CleiCode             string                `xml:"clei-code"`
	ModelNumber          string                `xml:"model-number"`
	ChassisSubSubModules []ChassisSubSubModule `xml:"chassis-sub-sub-module"`
}

type ChassisSubSubModule struct { //For example Transceivers
	Name         string `xml:"name"`
	Version      string `xml:"version"`
	PartNumber   string `xml:"part-number"`
	SerialNumber string `xml:"serial-number"`
	Description  string `xml:"description"`
}

func (targetDevice *TargetDevice) GetChassisHardware() (ChassisInventory, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get-chassis-inventory/>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return ChassisInventory{}, err
	}

	if rpcReply.Error() != nil {
		return ChassisInventory{}, rpcReply.Error()

	}
	var inventory ChassisInventory
	err = xml.Unmarshal(rpcReply.Content, &inventory)
	if err != nil {
		return inventory, err
	}
	return inventory, nil
}

/*
<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos">
<system-information>
  <hardware-model>qfx5200-32c-32q</hardware-model>
  <os-name>junos-qfx</os-name>
  <os-version>15.1X53-D30.5</os-version>
  <serial-number>WH0217100005</serial-number>
  <host-name>test-qfx5200</host-name>
</system-information>
</rpc-reply>
*/

type SystemInformation struct {
	XMLName       xml.Name `xml:"system-information"`
	HardwareModel string   `xml:"hardware-model"`
	OSName        string   `xml:"os-name"`
	/*
		os-name:
		junos - ACX Series, EX Series (certain platforms), MX Series, PTX Series
		junos-es - J Series, LN Series, SRX Series
		junos-ex - EX Series (certain platforms)
		junos-qfx - QFX Series
	*/
	OSVersion    string `xml:"os-version"`
	SerialNumber string `xml:"serial-number"`
	HostName     string `xml:"host-name"`
}

func (targetDevice *TargetDevice) GetSystemInformation() (SystemInformation, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get-system-information/>`),
	}
	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return SystemInformation{}, err
	}

	if rpcReply.Error() != nil {
		return SystemInformation{}, rpcReply.Error()

	}
	var systemInformation SystemInformation
	err = xml.Unmarshal(rpcReply.Content, &systemInformation)
	return systemInformation, err
}
