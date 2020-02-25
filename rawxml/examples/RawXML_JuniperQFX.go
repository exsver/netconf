package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/exsver/netconf/netconf"
	"github.com/exsver/netconf/rawxml"
)

func main() {
	netconf.LogLevel.Verbose()

	targetDevice := &netconf.TargetDevice{
		IP:   "10.10.10.10",
		Port: 830,
		SSHConfig: ssh.ClientConfig{
			Config: ssh.Config{
				Ciphers: []string{"aes128-ctr", "hmac-sha1"}, // aes128-cbc for HP5940  aes128-ctr for juniper QFX5100 or juniper MX
			},
			User:            "netconf-user",
			Auth:            []ssh.AuthMethod{ssh.Password("netconf-password")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         120 * time.Second,
		},
	}

	err := targetDevice.Connect(300 * time.Second)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	defer targetDevice.Disconnect()

	_, _ = targetDevice.NetconfSession.SendAndReceive([]byte(rawxml.XMLMessagesJunOS["GetChassisInventory"]))

}
