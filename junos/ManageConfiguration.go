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

	rpcReply, err := targetDevice.Action(request, "")
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

	if rpcReply.GetErrors() != nil {
		return nil, rpcReply.GetErrors()
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
	if !(format == "text" || format == "xml" || format == "set" || format == "json") {
		return nil, fmt.Errorf("wrong format string: %s. Allowed formats are: text | set | xml | json", format)
	}

	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<get-configuration format="%s"/>`, format)),
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return nil, err
	}

	if rpcReply.GetErrors() != nil {
		return nil, rpcReply.GetErrors()
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
	XMLName xml.Name           `xml:"load-configuration-results"`
	OK      bool               `xml:"ok"`
	Errors  []netconf.RPCError `xml:"rpc-error"`
}

func (results *LoadConfigurationResults) GetErrors() error {
	if results.Errors == nil {
		if !results.OK { // OK not found in reply
			return fmt.Errorf("unknown status: neither OK neither Errors not found in the reply")
		}

		return nil
	}

	var errString string

	for _, rpcErr := range results.Errors {
		errString = fmt.Sprintf("%s%s\n", errString, rpcErr.Error())
	}

	return fmt.Errorf("%s", errString)
}

type ConfigurationInformation struct {
	ConfigurationOutput ConfigurationOutput `xml:"configuration-output"`
}

type ConfigurationOutput struct {
	Data []byte `xml:",innerxml"`
}

func (targetDevice *TargetDevice) CompareConfigurationRollback(rollback int, format string) ([]byte, error) {
	if rollback < 0 || rollback > 49 {
		return nil, fmt.Errorf("invalid rollback-number")
	}

	if !(format == "text" || format == "xml" || format == "set" || format == "json") {
		return nil, fmt.Errorf("wrong format string: %s. Allowed formats are: text | set | xml | json", format)
	}

	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<get-configuration compare="rollback" rollback="%v" format="%s"/>`, rollback, format)),
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return nil, err
	}

	if rpcReply.GetErrors() != nil {
		return nil, rpcReply.GetErrors()
	}

	if format == "xml" || format == "json" {
		return rpcReply.Content, nil
	}

	var configuration ConfigurationInformation

	err = xml.Unmarshal(rpcReply.Content, &configuration)
	if err != nil {
		return configuration.ConfigurationOutput.Data, err
	}

	return configuration.ConfigurationOutput.Data, nil
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

	if rpcReply.GetErrors() != nil {
		return rpcReply.GetErrors()
	}

	var loadConfigurationResults LoadConfigurationResults

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)

	err = xml.Unmarshal(rpcReply.Content, &loadConfigurationResults)
	if err != nil {
		return err
	}

	if loadConfigurationResults.Errors != nil {
		return loadConfigurationResults.GetErrors()
	}

	if !loadConfigurationResults.OK {
		return errors.New("LoadConfigurationRolback: Unknown status")
	}

	return nil
}

// format attribute value json added in Junos OS Release 16.1.
// load-configuration capability URI: http://xml.juniper.net/netconf/junos/1.0
func (targetDevice *TargetDevice) LoadConfiguration(format, configuration, action string) (*LoadConfigurationResults, error) {
	request := netconf.RPCMessage{}

	switch format {
	case "xml":
		request.InnerXML = []byte(`<load-configuration action="merge" format="xml"><configuration>conf_text</configuration></load-configuration>`)
	case "text":
		request.InnerXML = []byte(`<load-configuration action="merge" format="text"><configuration-text>conf_text</configuration-text></load-configuration>`)
	case "json":
		request.InnerXML = []byte(`<load-configuration action="merge" format="json"><configuration-json>conf_text</configuration-json></load-configuration>`)
	default:
		return nil, errors.New("wrong format string. Allowed formats are: xml | text | json")
	}

	switch action {
	case "merge":
	case "update":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("merge"), []byte("update"), 1)
	case "override":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("merge"), []byte("override"), 1)
	case "replace":
		request.InnerXML = bytes.Replace(request.InnerXML, []byte("merge"), []byte("replace"), 1)
	case "set":
		request.InnerXML = []byte(`<load-configuration action="set" format="text"><configuration-set>conf_text</configuration-set></load-configuration>`)
	default:
		return nil, errors.New("wrong action string. Allowed actions are: merge | override | replace | update | set")
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte("conf_text"), []byte(configuration), 1)

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return nil, err
	}

	if rpcReply.GetErrors() != nil {
		return nil, rpcReply.GetErrors()
	}

	var loadConfigurationResults LoadConfigurationResults

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)

	err = xml.Unmarshal(rpcReply.Content, &loadConfigurationResults)
	if err != nil {
		return nil, err
	}

	return &loadConfigurationResults, nil
}

func (targetDevice *TargetDevice) OpenConfigurationPrivate() (string, error) {
	return targetDevice.OpenConfiguration("private")
}

func (targetDevice *TargetDevice) OpenConfiguration(mode string) (string, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(fmt.Sprintf(`<open-configuration><%s/></open-configuration>`, mode)),
	}

	if targetDevice.NetconfSession == nil {
		return "", fmt.Errorf("open-configuration: cannot execute RPC, you need to open a netconf session first")
	}

	rpcReply, err := targetDevice.Exec(request, "")
	if err != nil {
		return "", err
	}

	var openConfigurationString string

	// loop for errors in rpcReply
	for _, rpcError := range rpcReply.Errors {
		if !rpcError.IsWarning() {
			return "", rpcReply.GetErrors()
		}

		openConfigurationString = fmt.Sprintf("%s %s", openConfigurationString, rpcError.Error())
	}

	return openConfigurationString, nil
}

type DatabaseStatusInformation struct {
	XMLName        xml.Name         `xml:"database-status-information"`
	DatabaseStatus []DatabaseStatus `xml:"database-status"`
}

type DatabaseStatus struct {
	XMLName   xml.Name `xml:"database-status"`
	User      string   `xml:"user"`
	Terminal  string   `xml:"terminal"`
	PID       string   `xml:"pid"`
	StartTime string   `xml:"start-time"`
	IdleTime  string   `xml:"idle-time"`
	EditPath  string   `xml:"edit-path"`
}

func (targetDevice *TargetDevice) GetDatabaseStatusInformation() (*DatabaseStatusInformation, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<get-database-status-information/>`),
	}

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return nil, err
	}

	if rpcReply.GetErrors() != nil {
		return nil, rpcReply.GetErrors()
	}

	var out DatabaseStatusInformation

	rpcReply.Content = netconf.ConvertToPairedTags(rpcReply.Content)

	err = xml.Unmarshal(rpcReply.Content, &out)

	return &out, err
}
