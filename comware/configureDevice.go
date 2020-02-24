package comware

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/exsver/netconf/netconf"
)

//Operations:
// create:
//   - If the table supports target creation and the specified target does not exist, the operation creates and then configures the target.
//   - If the specified target exists, a data-exist error message is returned.
// merge:
//   - If the specified target exists, the operation directly changes the configuration for the target.
//   - If the specified target does not exist, the operation creates and configures the target.
//   - If the specified target does not exist and it cannot be created, an error message is returned.
// replace:
//   - If the specified target exists, the operation replaces the configuration of the target with the configuration carried in the message.
//   - If the specified target does not exist but is allowed to be created, create the target and then apply the configuration of the target.
//   - If the specified target does not exist and is not allowed to be created, the operation is not conducted and an invalid-value error message is returned.
// remove:
//   - If the specified target has only the table index, the operation removes all configuration of the specified target, and the target itself.
//   - If the specified target has the table index and configuration data, the operation removes the specified configuration data of this target.
//   - If the specified target does not exist, or the XML message does not specify any target, a success message is returned.
// delete:
//   - If the specified target has only the table index, the operation removes all configuration of the specified target, and the target itself.
//   - If the specified target has the table index and configuration data, the operation removes the specified configuration data of this target.
//   - If the specified target does not exist, an error message is returned, showing that the target does not exist.
func (targetDevice *TargetDevice) Configure(config Top, operation string) error {
	configXML, err := xml.Marshal(config)
	if err != nil {
		return err
	}

	request := netconf.RPCMessage{
		InnerXML: []byte(`<edit-config><target><running/></target><config><top/></config></edit-config>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("<top/>"), configXML, 1)

	switch operation {
	case "merge":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<top>"), []byte(`<top xmlns="http://www.hp.com/netconf/config:1.0" xmlns:base="urn:ietf:params:xml:ns:netconf:base:1.0" base:operation="merge">`), 1)
	case "create":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<top>"), []byte(`<top xmlns="http://www.hp.com/netconf/config:1.0" xmlns:base="urn:ietf:params:xml:ns:netconf:base:1.0" base:operation="create">`), 1)
	case "replace":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<top>"), []byte(`<top xmlns="http://www.hp.com/netconf/config:1.0" xmlns:base="urn:ietf:params:xml:ns:netconf:base:1.0" base:operation="replace">`), 1)
	case "remove":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<top>"), []byte(`<top xmlns="http://www.hp.com/netconf/config:1.0" xmlns:base="urn:ietf:params:xml:ns:netconf:base:1.0" base:operation="remove">`), 1)
	case "delete":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<top>"), []byte(`<top xmlns="http://www.hp.com/netconf/config:1.0" xmlns:base="urn:ietf:params:xml:ns:netconf:base:1.0" base:operation="delete">`), 1)
	default:
		return fmt.Errorf(`invalid operation string: "%s". Valid values are: "create", "merge", "replace", "remove", "delete"`, operation)
	}

	request.InnerXML = netconf.Normalize(request.InnerXML)

	rpcReply, err := targetDevice.Action(request, "running")
	if err != nil {
		return err
	}

	if rpcReply.GetErrors() != nil {
		return rpcReply.GetErrors()
	}

	if !rpcReply.OK {
		return fmt.Errorf("unknown Error")
	}

	return nil
}

// RetrieveData может вернуть структуру data без элементов. Это может возникнуть в случае запроса данных о несуществующем элементе. Например, если у 24-портового коммутатора спросить информацию о порте №30.
// Что-бы избежать nil pointer exception, в функциях, которые вызывают RetrieveData должна производится проверка на неравенство nil структуры data.top
// По этому принцыпу работают все функции с именами Is****Exist(). Они вызывают RetrieveData() и в случае отсутсвия ошибки проверяют data.top == nil. Если data.top существует - возвращают true.
func (targetDevice *TargetDevice) RetrieveData(request netconf.RPCMessage) (data Data, err error) {
	request.InnerXML = netconf.Normalize(request.InnerXML)

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return data, err
	}

	if rpcReply.GetErrors() != nil {
		return data, rpcReply.GetErrors()
	}

	err = xml.Unmarshal(rpcReply.Content, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (targetDevice *TargetDevice) PerformAction(request netconf.RPCMessage) error {
	request.InnerXML = netconf.Normalize(request.InnerXML)

	rpcReply, err := targetDevice.Action(request, "running")
	if err != nil {
		return err
	}

	if rpcReply.GetErrors() != nil {
		return rpcReply.GetErrors()
	}

	return nil
}
