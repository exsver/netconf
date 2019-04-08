package comware

import (
	"bytes"
	"encoding/xml"
	"github.com/exsver/netconf"
	"strings"
)

type CLIResponse struct {
	XMLName       xml.Name `xml:"CLI"`
	Execution     CDATA    `xml:"Execution,omitempty"`
	Configuration CDATA    `xml:"Configuration,omitempty"`
}

type CDATA struct {
	Data []byte `xml:",cdata"`
}

// RunCLICommand sends the specified cli commands via netconf.
// Use configurationMode:
// 		false - for execute commands in unprivileged mode
//		true  - for execute commands in privileged mode (system-view)
func (targetDevice *TargetDevice) RunCLICommand(command string, configurationMode bool) ([]byte, error) {

	request := netconf.RPCMessage{Xmlns: []string{netconf.BaseURI}}
	if configurationMode == true {
		request.InnerXML = []byte(`<CLI><Configuration>cmd_line</Configuration></CLI>`)
	} else {
		request.InnerXML = []byte(`<CLI><Execution>cmd_line</Execution></CLI>`)
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("cmd_line"), []byte(command), 1)

	rpcReply, err := targetDevice.Action(request, "")
	if err != nil {
		return nil, err
	}

	if rpcReply.Error() != nil {
		return nil, rpcReply.Error()
	}
	var cliResponse CLIResponse
	err = xml.Unmarshal([]byte(rpcReply.Content), &cliResponse)
	if err != nil {
		return nil, err
	}
	if configurationMode == true {
		return cliResponse.Configuration.Data, nil
	}
	return cliResponse.Execution.Data, nil
}

// required Comvare version >= 7.1.070
func (targetDevice *TargetDevice) IsConfigurationSaved() (saved bool, diff []byte, err error) {
	diff, err = targetDevice.RunCLICommand(`display current-configuration diff`, false)
	diff = bytes.Replace(diff, []byte("\n\n\n"), []byte("\n"), -1)
	if err != nil {
		return
	}
	diffLines := bytes.Split(diff, []byte("\n"))
	if len(diffLines) == 2 && bytes.Equal(diffLines[1], []byte{}) {
		return true, diff, err
	}
	return false, diff, err
}

func vlanListCLIToIntSlice(vlanList string) ([]int, error) {
	vlanList = strings.TrimPrefix(vlanList, "port trunk permit vlan ") // for trunk port config
	vlanList = strings.TrimPrefix(vlanList, "port hybrid vlan ")       // for hybrid port config
	vlanList = strings.TrimSuffix(vlanList, " tagged")                 // for hybrid port config
	vlanList = strings.TrimSuffix(vlanList, " untagged")               // for hybrid port config
	vlanList = strings.TrimPrefix(vlanList, "port access vlan ")       // for access port config
	vlanList = strings.Replace(vlanList, " to ", "-", -1)
	vlanList = strings.Replace(vlanList, " ", ",", -1)
	return VlanListToIntSlice(vlanList)
}

func ParseVlansFromConfigString(configString string) (vlans []int) {
	lines := strings.Split(configString, "\n")
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "port trunk permit vlan ") || strings.HasPrefix(line, "port hybrid vlan ") || strings.HasPrefix(line, "port access vlan ") {
			v, err := vlanListCLIToIntSlice(line)
			if err == nil {
				vlans = append(vlans, v...)
			}
		}
	}
	return vlans
}

// CorrectNewLines replaces "\n\n\n" with "\n"
func CorrectNewLines(in []byte) []byte {
	return bytes.Replace(in, []byte("\n\n\n"), []byte("\n"), -1)
}
