package junos

import (
	"encoding/xml"

	"github.com/exsver/netconf/netconf"
	)

type RebootResults struct {
	RequestRebootResults xml.Name `xml:"request-reboot-results"`
	RequestRebootStatus  string   `xml:"request-reboot-status"`
}

func (targetDevice *TargetDevice) Reboot() error {
	request := netconf.RPCMessage{
		InnerXML: []byte(`<request-reboot/>`),
		Xmlns:    []string{netconf.BaseURI},
	}
	_, err := targetDevice.Action(request, "")
	return err

	/*
	   <rpc-reply xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="1">
	   <request-reboot-results>
	   <request-reboot-status reboot-time>
	   Shutdown at Tue May 22 18:49:01 2018.
	   [pid 14327]
	   </request-reboot-status>
	   </request-reboot-results>
	   </rpc-reply>

	   <rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos"  message-id="1">
	   <rpc-error>
	   <error-severity>warning</error-severity>
	   <error-message>
	   The configuration has been changed but not committed
	   <error-message>
	   <rpc-error>
	   <request-reboot-results>
	   <request-reboot-status reboot-time>
	   Shutdown at Tue May 22 19:14:07 2018.
	   [pid 3523]
	   </request-reboot-status>
	   </request-reboot-results>
	   <rpc-reply>
	*/
}
