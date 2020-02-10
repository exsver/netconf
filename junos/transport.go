package junos

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"

	"github.com/exsver/netconf/netconf"
)

type TargetDevice struct {
	netconf.TargetDevice
}

func NewTargetDevice(ipString string, username string, password string) (*TargetDevice, error) {
	parsedIP := net.ParseIP(ipString)
	if parsedIP == nil {
		return nil, fmt.Errorf("create NewTargetDevice Error. Can't Parse IP from string: <%s>", ipString)
	}

	return &TargetDevice{netconf.TargetDevice{
		IP:   parsedIP.String(),
		Port: 830,
		SSHConfig: ssh.ClientConfig{
			Config: ssh.Config{
				Ciphers: []string{"aes128-ctr", "hmac-sha1"},
			},
			User:            username,
			Auth:            []ssh.AuthMethod{ssh.Password(password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         30 * time.Second},
	},
	}, nil
}
