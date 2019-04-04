package main

import (
	"github.com/exsver/netconf"
	"github.com/exsver/netconf/rawxml"
	"golang.org/x/crypto/ssh"
	"log"
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

	err := targetDevice.Connect(300 * time.Second)
	if err != nil {
		log.Printf("%s\n", err.Error())
	}

	defer targetDevice.NetconfSession.Close()

	_, _ = targetDevice.NetconfSession.SendAndReceive([]byte(netconf.XmlHello))
	_, _ = targetDevice.NetconfSession.SendAndReceive([]byte(rawxml.XMLMessagesHPE["GetDeviceBaseHostname"]))
	_, _ = targetDevice.NetconfSession.SendAndReceive([]byte(netconf.XmlClose))

}
