package junos

import (
	"bytes"
	"github.com/exsver/netconf"
)

func (targetDevice *TargetDevice) RunCLICommand(command string) ([]byte, error) {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<command format="text">cmd_line</command>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	request.InnerXML = bytes.Replace(request.InnerXML, []byte("cmd_line"), []byte(command), 1)
	rpcReply, err := targetDevice.Action(request, "")
	return rpcReply.Content, err
}
