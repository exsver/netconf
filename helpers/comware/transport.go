package comware

import (
	"fmt"
	"github.com/exsver/netconf"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
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
				Ciphers: []string{"aes128-cbc", "hmac-sha1"},
			},
			User:            username,
			Auth:            []ssh.AuthMethod{ssh.Password(password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         30 * time.Second},
	},
	}, nil
}
