package junos

import (
	"bytes"
	"fmt"

	"github.com/exsver/netconf/netconf"
)

func (targetDevice *TargetDevice) RunCLICommand(command string) ([]byte, error) {
	if len(command) == 0 {
		return nil, fmt.Errorf("empty command not alloved")
	}

	request := netconf.RPCMessage{
		InnerXML: []byte(`<command format="text">cmd_line</command>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	request.InnerXML = bytes.Replace(request.InnerXML, []byte("cmd_line"), []byte(command), 1)
	rpcReply, err := targetDevice.Action(request, "")

	return rpcReply.Content, err
}
