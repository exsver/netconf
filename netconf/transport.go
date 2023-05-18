package netconf

import "golang.org/x/crypto/ssh"

// Transport stores runtime variables provided by ssh.
type Transport struct {
	Session    *ssh.Session
	Connection *ssh.Client
}

func (t *Transport) Close() error {
	t.Session.Close()

	err := t.Connection.Close()
	if err != nil {
		if err.Error() == "EOF" {
			return nil
		}

		LogLevel.Info.Printf("Failed to close session: %s\n", err.Error())

		return err
	}

	return nil
}
