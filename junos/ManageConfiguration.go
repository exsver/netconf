package junos

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"

	"github.com/exsver/netconf/netconf"
)

// Operations:
//  merge:
//    The device merges new configuration data into the existing configuration data. This is the default.
//  replace:
//    The device replaces existing configuration data with the new configuration data.
//  none:
//    The device does not change the existing configuration unless the new configuration element includes an operation attribute.
func (targetDevice *TargetDevice) EditConfig(config *Config, operation string) error {
	if config == nil {
		return fmt.Errorf("nothing to configure")
	}

	configXML, err := xml.Marshal(config)
	if err != nil {
		return err
	}

	configXML = netconf.ConvertToSelfClosingTag(configXML)

	request := netconf.RPCMessage{
		InnerXML: []byte(`<edit-config><target><candidate/></target><default-operation>merge</default-operation><configuration/></edit-config>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	switch operation {
	case "merge", "":
	case "replace":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<default-operation>merge</default-operation>"), []byte(`<default-operation>replace</default-operation>`), 1)
	case "none":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("<default-operation>merge</default-operation>"), []byte(`<default-operation>none</default-operation>`), 1)
	default:
		return fmt.Errorf(`invalid operation string: "%s". Valid values are: "merge", "replace", "none"`, operation)
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte("<configuration/>"), configXML, 1)
	request.InnerXML = netconf.Normalize(request.InnerXML)

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return err
	}

	if rpcReply.Error() != nil {
		return rpcReply.Error()
	}

	if !rpcReply.OK {
		return fmt.Errorf("unknown Error")
	}

	return nil
}

// CLI equivalent: show configuration *subtree*
// Source:
//  - "running"
//  - "candidate"
// Subtree examples:
//  - ""
//  - "system"
//  - "interfaces"
//  - "snmp"
//  - "forwarding-options"
//  - "routing-options"
//  - "routing-options/static"
//  - "policy-options"
//  - "protocols"
//  - "protocols/bgp"
//  - "vlans"
//  - "firewall"
func (targetDevice *TargetDevice) GetConfig(source string, subtree string) (*Configuration, error) {
	request := netconf.RPCMessage{Xmlns: []string{netconf.BaseURI}}
	if subtree == "" {
		request = netconf.RPCMessage{
			InnerXML: []byte(`
             <get-config>
               <source>
                 <candidate/>
               </source>
             </get-config>`)}
	} else {
		request = netconf.RPCMessage{
			InnerXML: []byte(`
             <get-config>
               <source>
                 <candidate/>
               </source>
               <filter type="subtree">
                 <configuration>
                   sub_tree
                 </configuration>
               </filter>
             </get-config>`)}
		request.InnerXML = bytes.Replace(request.InnerXML, []byte(`sub_tree`), netconf.ConvertToXML([]byte(subtree)), 1)
	}

	if source != "" {
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("candidate"), []byte(source), 1)
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return nil, err
	}

	if rpcReply.Error() != nil {
		return nil, rpcReply.Error()
	}

	var data Data

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)

	err = xml.Unmarshal(rpcReply.Content, &data)
	if err != nil {
		return nil, err
	}

	return data.Configuration, nil
}

type ConfigurationBytes struct {
	Data []byte `xml:",innerxml"`
}

// format: text | set | xml | json
// database: candidate | committed
// inherit: defaults | inherit
// compare rollback [rollback="[0-49]"
func (targetDevice *TargetDevice) GetConfiguration(format string) ([]byte, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(` <get-configuration format="text"/>`),
	}

	switch format {
	case "text":
	case "set":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("text"), []byte("set"), 1)
	case "xml":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("text"), []byte("xml"), 1)
	case "json":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("text"), []byte("json"), 1)
	default:
		return nil, errors.New("wrong format string. Allowed formats are: text | set | xml | json ")
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return nil, err
	}

	if rpcReply.Error() != nil {
		return nil, rpcReply.Error()
	}

	if format == "xml" || format == "json" {
		return rpcReply.Content, nil
	}

	var configurationText ConfigurationBytes

	err = xml.Unmarshal(rpcReply.Content, &configurationText)
	if err != nil {
		return configurationText.Data, err
	}

	return configurationText.Data, nil
}

type LoadConfigurationResults struct {
	XMLName xml.Name         `xml:"load-configuration-results"`
	OK      bool             `xml:"ok"`
	Error   netconf.RPCError `xml:"rpc-error"`
}

// CLI equivalent:  rollback [0..49]
func (targetDevice *TargetDevice) LoadConfigurationRolback(rollback int) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<load-configuration rollback="0"/>`),
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("0"), []byte(strconv.Itoa(rollback)), 1)

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return err
	}

	if rpcReply.Error() != nil {
		return rpcReply.Error()
	}

	var loadConfigurationResults LoadConfigurationResults

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)

	err = xml.Unmarshal(rpcReply.Content, &loadConfigurationResults)
	if err != nil {
		return err
	}

	if loadConfigurationResults.Error.Error() != nil {
		return loadConfigurationResults.Error.Error()
	}

	if !loadConfigurationResults.OK {
		return errors.New("LoadConfigurationRolback: Unknown status")
	}

	return nil
}

// format attribute value json added in Junos OS Release 16.1.
// load-configuration capability URI: http://xml.juniper.net/netconf/junos/1.0
func (targetDevice *TargetDevice) LoadConfigurationText(configuration string, action string) error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<load-configuration action="merge" format="text"><configuration-text>conf_text</configuration-text></load-configuration>`),
	}

	switch action {
	case "merge":
	case "update":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("merge"), []byte("update"), 1)
	case "override":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("merge"), []byte("override"), 1)
	case "replace":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("merge"), []byte("replace"), 1)
	default:
		return errors.New("wrong action string. Allowed actions are: merge | override | replace | update")
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte("conf_text"), []byte(configuration), 1)

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return err
	}

	if rpcReply.Error() != nil {
		return rpcReply.Error()
	}

	var loadConfigurationResults LoadConfigurationResults

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)

	err = xml.Unmarshal(rpcReply.Content, &loadConfigurationResults)
	if err != nil {
		return err
	}

	if loadConfigurationResults.Error.Error() != nil {
		return loadConfigurationResults.Error.Error()
	}

	if !loadConfigurationResults.OK {
		return errors.New("LoadConfiguration: Unknown status")
	}

	return nil
}
