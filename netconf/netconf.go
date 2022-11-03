package netconf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

// NetconfSession combine IO Reader and WriteCloser.
type NetconfSession struct {
	io.Reader
	io.WriteCloser
	Capabilities     []string
	SessionID        string
	SessionTransport *Transport
	lastMessageID    int
}

// TargetDevice represents all configuration options for ssh transport.
type TargetDevice struct {
	SSHConfig      ssh.ClientConfig
	IP             string
	Port           int
	NetconfSession *NetconfSession
}

func (netconfSession *NetconfSession) Send(data []byte) error {
	w := bufio.NewWriter(netconfSession.WriteCloser)

	n, err := w.Write(data)
	if err != nil {
		return err
	}

	// Pad to make sure the msgSeparator isn't sent across a 4096-byte boundary
	if (n+len(messageSeparator))%4096 < len(messageSeparator) {
		padding := make([]byte, len(messageSeparator))
		_, err = w.Write(padding)

		if err != nil {
			return err
		}
	}

	_, err = w.Write([]byte(messageSeparator))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	w.Reset(w)

	return nil
}

func (netconfSession *NetconfSession) Receive() ([]byte, error) {
	r := bufio.NewReader(netconfSession.Reader)

	var out = make([]byte, 0)

	//	time.Sleep(time.Millisecond * 100)
	for {
		b, err := r.ReadByte()
		out = append(out, b)

		if bytes.HasSuffix(out, []byte(messageSeparator)) {
			return bytes.TrimSpace(bytes.TrimSuffix(out, []byte(messageSeparator))), nil
		}

		if err != nil {
			return bytes.TrimSpace(bytes.TrimSuffix(out, []byte(messageSeparator))), err
		}
	}
}

func (netconfSession *NetconfSession) SendAndReceive(data []byte) ([]byte, error) {
	LogLevel.Info.Printf("Sending payload...\n")
	LogLevel.Message.Printf("\nSend:\n%s\n", string(bytes.Trim(data, "\n")))

	err := netconfSession.Send(data)
	if err != nil {
		LogLevel.Fail.Printf("Sending payload failed: %s\n", err.Error())

		return nil, err
	}

	LogLevel.Info.Printf("Payload send.\n")
	LogLevel.Info.Printf("Receiving payload...\n")

	var rcv []byte

	for {
		rcv, err = netconfSession.Receive()
		if err != nil {
			LogLevel.Fail.Printf("Receiving payload failed: %s\n", err.Error())
			return nil, err
		}

		if bytes.Equal(Normalize(rcv), Normalize(data)) {
			LogLevel.Info.Printf("Message come back")
			continue
		}

		break
	}

	LogLevel.Info.Printf("Payload received.\n")
	LogLevel.Message.Printf("\nReceived:\n%s\n", string(bytes.Trim(rcv, "\n")))

	return rcv, nil
}

func (targetDevice *TargetDevice) PrepareTransport(deadline time.Duration) (err error) {
	if targetDevice.NetconfSession.SessionTransport == nil {
		targetDevice.NetconfSession.SessionTransport = &Transport{}
	}

	target := fmt.Sprintf("%s:%d", targetDevice.IP, targetDevice.Port)

	LogLevel.Info.Printf("Initiating SSH connection...\n")

	targetDevice.NetconfSession.SessionTransport.Connection, err = ssh.Dial("tcp", target, &targetDevice.SSHConfig)
	if err != nil {
		LogLevel.Fail.Printf("SSH connection failed: %s\n", err.Error())
		return err
	}

	LogLevel.Info.Printf("SSH connection initiated\n")

	LogLevel.Info.Printf("Creating new session...\n")

	targetDevice.NetconfSession.SessionTransport.Session, err = targetDevice.NetconfSession.SessionTransport.Connection.NewSession()
	if err != nil {
		LogLevel.Fail.Printf("Failed to create new session: %s\n", err.Error())
		return err
	}

	LogLevel.Info.Printf("New session created\n")

	LogLevel.Info.Printf("Requesting NETCONF subsystem...\n")

	err = targetDevice.NetconfSession.SessionTransport.Session.RequestSubsystem("netconf")
	if err != nil {
		LogLevel.Fail.Printf("Failed to set NETCONF subsystem: %s\n", err.Error())
		return err
	}

	LogLevel.Info.Printf("Enabled NETCONF subsystem\n")

	LogLevel.Info.Printf("Requesting Stdin redirection...\n")

	targetDevice.NetconfSession.WriteCloser, err = targetDevice.NetconfSession.SessionTransport.Session.StdinPipe()
	if err != nil {
		LogLevel.Fail.Printf("Failed to set StdinPipe: %s\n", err.Error())
		return err
	}

	LogLevel.Info.Printf("Stdin redirected\n")

	LogLevel.Info.Printf("Requesting Stdout redirection...\n")

	targetDevice.NetconfSession.Reader, err = targetDevice.NetconfSession.SessionTransport.Session.StdoutPipe()
	if err != nil {
		LogLevel.Fail.Printf("Failed to set StdoutPipe: %s\n", err.Error())
		return err
	}

	LogLevel.Info.Printf("Stdout redirected\n")

	if deadline != 0 {
		go func(t *Transport, deadline time.Duration) {
			LogLevel.Info.Printf("Starting deadline timer for current SSH session and connection...\n")

			_ = time.AfterFunc(deadline, func() {
				LogLevel.Info.Printf("Deadline timer triggered, closing SSH session\n")
				err := t.Close()
				if err != nil {
					LogLevel.Info.Printf("Deadline timer: %s\n", err)
				}
			})
		}(targetDevice.NetconfSession.SessionTransport, deadline)
	}

	return nil
}

func (netconfSession *NetconfSession) SendHello() (err error) {
	rawReply, err := netconfSession.SendAndReceive([]byte(XMLHello))
	if err != nil {
		return
	}

	hello, err := unmarshalHello(rawReply)
	if err != nil {
		return
	}

	netconfSession.SessionID = hello.SessionID
	netconfSession.Capabilities = hello.Capabilities

	return nil
}

// CloseSession request graceful termination of a NETCONF session.
func (netconfSession *NetconfSession) CloseSession(messageID string) error {
	closeMessage := []byte(XMLClose)
	if messageID != "" {
		closeMessage = bytes.Replace(closeMessage, []byte("close-uuid"), []byte(messageID), 1)
	}

	rawReply, err := netconfSession.SendAndReceive(closeMessage)
	if err != nil {
		return err
	}

	rpcReply, err := UnmarshalRPCReply(rawReply)
	if err != nil {
		return err
	}

	if rpcReply.MessageID != messageID {
		return fmt.Errorf("closeSession: message-ID mismatch for close message. Got : %s, Want: %s", rpcReply.MessageID, messageID)
	}

	if !rpcReply.OK {
		return fmt.Errorf("closeSession: can't find ok message in reply")
	}

	return nil
}

// After a user locks the configuration, other users cannot use NETCONF or other configuration methods such as CLI and SNMP to configure the device.
// When a NETCONF session is terminated, the related locked configuration is also unlocked.
func (netconfSession *NetconfSession) lockConfig(messageID string, config string) error {
	lockMessage := RPCMessage{
		MessageID: messageID,
		InnerXML:  generateLock(config),
		Xmlns:     []string{BaseURI},
	}

	lockRPCMessage, err := lockMessage.MarshalRPCMessage()
	if err != nil {
		return err
	}

	rawReply, err := netconfSession.SendAndReceive(lockRPCMessage)
	if err != nil {
		return err
	}

	rpcReply, err := UnmarshalRPCReply(rawReply)
	if err != nil {
		return err
	}

	if rpcReply.MessageID != lockMessage.MessageID {
		return fmt.Errorf("mismatch message ID for lock message. Got: %s, Want: %s", rpcReply.MessageID, lockMessage.MessageID)
	}

	if !rpcReply.OK || len(rpcReply.Errors) != 0 {
		return fmt.Errorf("can't lock config")
	}

	return nil
}

func (netconfSession *NetconfSession) unLockConfig(messageID string, config string) error {
	lockMessage := RPCMessage{
		MessageID: messageID,
		InnerXML:  generateUnLock(config),
		Xmlns:     []string{BaseURI},
	}

	lockRPCMessage, err := lockMessage.MarshalRPCMessage()
	if err != nil {
		return err
	}

	rawReply, err := netconfSession.SendAndReceive(lockRPCMessage)
	if err != nil {
		return err
	}

	rpcReply, err := UnmarshalRPCReply(rawReply)
	if err != nil {
		return err
	}

	if !rpcReply.OK || len(rpcReply.Errors) != 0 {
		return fmt.Errorf("can't unlock config")
	}

	return nil
}

// Action simplifies netconf protocol operations. It performs initialization/closing sessions, message-id handling.
func (targetDevice *TargetDevice) Action(rpcMessage RPCMessage, lockConfig string) (rpcReply *RPCReply, err error) {
	if !rpcMessage.NotNormalize {
		rpcMessage.InnerXML = Normalize(rpcMessage.InnerXML)
	}

	isSessionExist := false
	if targetDevice.NetconfSession != nil {
		isSessionExist = true
	}

	if !isSessionExist {
		err = targetDevice.Connect(targetDevice.SSHConfig.Timeout)
		if err != nil {
			return nil, err
		}
	}

	rpcReply, err = targetDevice.Exec(rpcMessage, lockConfig)

	if !isSessionExist {
		targetDevice.Disconnect()
	}

	return rpcReply, err
}

func (targetDevice *TargetDevice) Exec(rpcMessage RPCMessage, lockConfig string) (rpcReply *RPCReply, err error) {
	if targetDevice.NetconfSession == nil {
		return nil, errors.New("netconf session is not exist")
	}
	// Lock config if lockConfig set
	if lockConfig != "" {
		targetDevice.NetconfSession.lastMessageID++

		err = targetDevice.NetconfSession.lockConfig(strconv.Itoa(targetDevice.NetconfSession.lastMessageID), lockConfig)
		if err != nil {
			return rpcReply, err
		}
	}

	targetDevice.NetconfSession.lastMessageID++
	rpcMessage.MessageID = strconv.Itoa(targetDevice.NetconfSession.lastMessageID)

	rpcMessageXML, err := rpcMessage.MarshalRPCMessage()
	if err != nil {
		return rpcReply, err
	}

	rawReply, err := targetDevice.NetconfSession.SendAndReceive(rpcMessageXML)
	if err != nil {
		return rpcReply, err
	}

	rpcReply, err = UnmarshalRPCReply(rawReply)
	if err != nil {
		return rpcReply, err
	}
	// Check message ID
	if rpcReply.MessageID != strconv.Itoa(targetDevice.NetconfSession.lastMessageID) {
		return rpcReply, fmt.Errorf("mismatch message ID for request . Got: %s, Want: %s", rpcReply.MessageID, strconv.Itoa(targetDevice.NetconfSession.lastMessageID))
	}
	// UnLock config if lockConfig set
	if lockConfig != "" {
		targetDevice.NetconfSession.lastMessageID++
		targetDevice.NetconfSession.unLockConfig(strconv.Itoa(targetDevice.NetconfSession.lastMessageID), lockConfig)
	}

	return
}

func (targetDevice *TargetDevice) Connect(timeout time.Duration) (err error) {
	if targetDevice.NetconfSession != nil {
		return errors.New("netconf session is exist")
	}

	targetDevice.NetconfSession = &NetconfSession{}

	err = targetDevice.PrepareTransport(timeout)
	if err != nil {
		return err
	}

	err = targetDevice.NetconfSession.SendHello()

	return err
}

// Disconnect request graceful termination of a NETCONF session and ssh connection.
func (targetDevice *TargetDevice) Disconnect() error {
	if targetDevice.NetconfSession == nil {
		return errors.New("netconf session is not exist")
	}

	targetDevice.NetconfSession.lastMessageID++

	// Closing netconf session
	err := targetDevice.NetconfSession.CloseSession(strconv.Itoa(targetDevice.NetconfSession.lastMessageID))
	if err != nil {
		LogLevel.Info.Printf("error while trying to close netconf session: %s /n", err.Error())
	}

	// Terminating ssh connection
	err = targetDevice.NetconfSession.SessionTransport.Close()
	if err != nil {
		LogLevel.Info.Printf("error while trying to terminate ssh connection: %s /n", err.Error())
	}

	targetDevice.NetconfSession = nil

	return err
}
