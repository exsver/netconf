package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/exsver/netconf"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func main() {
	// netconf.LogLevel.Verbose()
	targetDevice := &netconf.TargetDevice{
		IP:   "10.10.10.10",
		Port: 830,
		SSHConfig: ssh.ClientConfig{
			Config: ssh.Config{
				Ciphers: []string{"aes128-cbc", "hmac-sha1"},
			},
			User:            "netconf-user",
			Auth:            []ssh.AuthMethod{ssh.Password("netconf-password")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         30 * time.Second},
	}

	message := netconf.RPCMessage{
		InnerXML: []byte(`<get><filter type="subtree"><top xmlns="http://www.hp.com/netconf/data:1.0"><Device><Base><HostName/></Base></Device></top></filter></get>`),
		Xmlns:    []string{netconf.BaseURI},
	}

	rpcReply, err := targetDevice.Action(message, "")
	if err != nil {
		log.Fatal(rpcReply)
	}
	spew.Dump(rpcReply)
}
