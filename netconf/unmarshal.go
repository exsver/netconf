package netconf

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type RPCReply struct {
	XMLName   xml.Name   `xml:"rpc-reply"`
	MessageID string     `xml:"message-id,attr"`
	Errors    []RPCError `xml:"rpc-error"`
	Content   []byte     `xml:",innerxml"`
	OK        bool       `xml:"-"`
}

type RPCError struct {
	XMLName       xml.Name     `xml:"rpc-error"`
	ErrorType     string       `xml:"error-type"`     // RFC6241: "transport" | "rpc" | "protocol" | "application"
	ErrorTag      string       `xml:"error-tag"`      // RFC6241: "in-use" | "invalid-value" | "bad-element" | ...   https://tools.ietf.org/html/rfc6241#appendix-A
	ErrorSeverity string       `xml:"error-severity"` // RFC6241: "error" | "warning"
	ErrorAppTag   string       `xml:"error-app-tag"`
	ErrorPath     string       `xml:"error-path"`
	ErrorMessage  string       `xml:"error-message"`
	ErrorInfo     RPCErrorInfo `xml:"error-info"`
}

type RPCErrorInfo struct {
	Info []byte `xml:",innerxml"`
}

func UnmarshalRPCReply(raw []byte) (rpcReply *RPCReply, err error) {
	raw = CorrectBackspaces(raw)

	err = xml.Unmarshal(raw, &rpcReply)

	if len(rpcReply.Errors) == 0 && (bytes.Equal(rpcReply.Content, []byte{0x3c, 0x6f, 0x6b, 0x2f, 0x3e}) || bytes.Equal(rpcReply.Content, []byte{0x0a, 0x3c, 0x6f, 0x6b, 0x2f, 0x3e, 0x0a})) {
		rpcReply.OK = true
	}

	return rpcReply, err
}

func (rpcReply *RPCReply) GetErrors() error {
	if rpcReply.Errors == nil {
		return nil
	}

	var errString string

	for _, rpcErr := range rpcReply.Errors {
		errString = fmt.Sprintf("%s%s\n", errString, rpcErr.Error())
	}

	return fmt.Errorf("%s", errString)
}

func (rpcError *RPCError) Error() string {
	errorString := ""
	if rpcError.ErrorType != "" {
		errorString = fmt.Sprintf("%serror-type: %s ", errorString, rpcError.ErrorType)
	}

	if rpcError.ErrorTag != "" {
		errorString = fmt.Sprintf("%serror-tag: %s ", errorString, rpcError.ErrorTag)
	}

	if rpcError.ErrorSeverity != "" {
		errorString = fmt.Sprintf("%serror-severity: %s ", errorString, rpcError.ErrorSeverity)
	}

	if rpcError.ErrorAppTag != "" {
		errorString = fmt.Sprintf("%serror-app-tag: %s ", errorString, rpcError.ErrorAppTag)
	}

	if rpcError.ErrorPath != "" {
		errorString = fmt.Sprintf("%serror-path: %s ", errorString, rpcError.ErrorPath)
	}

	if rpcError.ErrorMessage != "" {
		errorString = fmt.Sprintf("%serror-message: %s ", errorString, rpcError.ErrorMessage)
	}

	if rpcError.ErrorInfo.Info != nil {
		errorString = fmt.Sprintf("%serror-info: %s ", errorString, rpcError.ErrorInfo.Info)
	}

	return errorString
}

type hello struct {
	XMLName      xml.Name `xml:"hello"`
	Capabilities []string `xml:"capabilities>capability"`
	SessionID    string   `xml:"session-id"`
}

func unmarshalHello(raw []byte) (hello *hello, err error) {
	err = xml.Unmarshal(raw, &hello)
	return hello, err
}
