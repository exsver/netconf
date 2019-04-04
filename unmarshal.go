package netconf

import (
	"bytes"
	"encoding/xml"
	"errors"
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
	ErrorType     string       `xml:"error-type"` //RFC6241: transport|rpc|protocol|application
	ErrorTag      string       `xml:"error-tag"`
	ErrorSeverity string       `xml:"error-severity"` //RFC6241: error|warning
	ErrorAppTag   string       `xml:"error-app-tag"`
	ErrorPath     string       `xml:"error-path"`
	ErrorMessage  string       `xml:"error-message"`
	ErrorInfo     RPCErrorInfo `xml:"error-info"`
}

type RPCErrorInfo struct {
	Info []byte `xml:",innerxml"`
}

func UnmarshalRpcReply(raw []byte) (rpcReply *RPCReply, err error) {
	raw = CorrectBackspaces(raw)
	err = xml.Unmarshal([]byte(raw), &rpcReply)
	if len(rpcReply.Errors) == 0 && (bytes.Equal(rpcReply.Content, []byte{0x3c, 0x6f, 0x6b, 0x2f, 0x3e}) || bytes.Equal(rpcReply.Content, []byte{0x0a, 0x3c, 0x6f, 0x6b, 0x2f, 0x3e, 0x0a})) {
		rpcReply.OK = true
	}
	return rpcReply, err
}

func (rpcReply *RPCReply) Error() error {
	if len(rpcReply.Errors) != 0 {
		errorString := ""
		for _, v := range rpcReply.Errors {
			if v.ErrorType != "" {
				errorString = fmt.Sprintf("%sERROR-TYPE: %s ", errorString, v.ErrorType)
			}
			if v.ErrorTag != "" {
				errorString = fmt.Sprintf("%sERROR-TAG: %s ", errorString, v.ErrorTag)
			}
			if v.ErrorSeverity != "" {
				errorString = fmt.Sprintf("%sERROR-SEVERITY: %s ", errorString, v.ErrorSeverity)
			}
			if v.ErrorAppTag != "" {
				errorString = fmt.Sprintf("%sERROR-APP-TAG: %s ", errorString, v.ErrorAppTag)
			}
			if v.ErrorPath != "" {
				errorString = fmt.Sprintf("%sERROR-PATH: %s ", errorString, v.ErrorPath)
			}
			if v.ErrorMessage != "" {
				errorString = fmt.Sprintf("%sERROR-MESSAGE: %s ", errorString, v.ErrorMessage)
			}
			if v.ErrorInfo.Info != nil {
				errorString = fmt.Sprintf("%sERROR-INFO: %s ", errorString, v.ErrorInfo.Info)
			}
			errorString = fmt.Sprintf("%s\n", errorString)
		}
		return errors.New(errorString)
	}
	return nil
}

func (rpcError *RPCError) Error() error {
	errorString := ""
	if rpcError.ErrorType != "" {
		errorString = fmt.Sprintf("%sERROR-TYPE: %s ", errorString, rpcError.ErrorType)
	}
	if rpcError.ErrorTag != "" {
		errorString = fmt.Sprintf("%sERROR-TAG: %s ", errorString, rpcError.ErrorTag)
	}
	if rpcError.ErrorSeverity != "" {
		errorString = fmt.Sprintf("%sERROR-SEVERITY: %s ", errorString, rpcError.ErrorSeverity)
	}
	if rpcError.ErrorAppTag != "" {
		errorString = fmt.Sprintf("%sERROR-APP-TAG: %s ", errorString, rpcError.ErrorAppTag)
	}
	if rpcError.ErrorPath != "" {
		errorString = fmt.Sprintf("%sERROR-PATH: %s ", errorString, rpcError.ErrorPath)
	}
	if rpcError.ErrorMessage != "" {
		errorString = fmt.Sprintf("%sERROR-MESSAGE: %s ", errorString, rpcError.ErrorMessage)
	}
	if rpcError.ErrorInfo.Info != nil {
		errorString = fmt.Sprintf("%sERROR-INFO: %s ", errorString, rpcError.ErrorInfo.Info)

	}

	if errorString != "" {
		return errors.New(errorString)
	}

	return nil
}

type hello struct {
	XMLName      xml.Name `xml:"hello"`
	Capabilities []string `xml:"capabilities>capability"`
	SessionID    string   `xml:"session-id"`
}

func unmarshalHello(raw []byte) (hello *hello, err error) {
	err = xml.Unmarshal([]byte(raw), &hello)
	return hello, err
}
