package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf"
	"golang.org/x/crypto/ssh"
	"time"
)

func main() {
	netconf.LogLevel.Verbose()
	targetDevice := &netconf.TargetDevice{
		IP:   "10.10.10.10",
		Port: 830,
		SSHConfig: ssh.ClientConfig{
			Config: ssh.Config{
				Ciphers: []string{"aes128-cbc", "hmac-sha1"},
			},
			User:            "netconf",
			Auth:            []ssh.AuthMethod{ssh.Password("netconf")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         30 * time.Second},
	}
	targetDevice.Connect(30 * time.Second)
	spew.Dump(targetDevice.NetconfSession.SessionID)
	spew.Dump(targetDevice.NetconfSession.Capabilities)
	message := netconf.RPCMessage{InnerXML: []byte(`<get-sessions/>`), Xmlns: []string{netconf.BaseURI}}
	targetDevice.Exec(message, "")
	targetDevice.Exec(message, "")
	targetDevice.Exec(message, "")
	targetDevice.Exec(message, "")
	targetDevice.Disconnect()
}
