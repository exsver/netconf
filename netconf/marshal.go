package netconf

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type RPCMessage struct {
	XMLName   xml.Name `xml:"rpc"`
	Xmlns     []string `xml:"xmlns,attr,omitempty"`
	MessageID string   `xml:"message-id,attr,omitempty"`
	// RPC-message body without rpc tag.
	InnerXML []byte `xml:",innerxml"`
	// Use CustomAttrs to set a custom attributes in rpc tag. For example custom namespaces.
	CustomAttrs []string `xml:"-"`
	// Set True for add optional XML header: `<?xml version="1.0" encoding="UTF-8"?>`.
	AppendXMLHeader bool `xml:"-"`
}

// MarshalRPCMessage generates and return a new RPC Message from RPCMessage struct.
func (message *RPCMessage) MarshalRPCMessage() (rpc []byte, err error) {
	rpc, err = xml.Marshal(message)
	if err != nil {
		return
	}

	if len(message.CustomAttrs) != 0 {
		nss := []byte("<rpc ")
		for _, v := range message.CustomAttrs {
			nss = append(nss, fmt.Sprintf("%s ", v)...)
		}

		rpc = bytes.Replace(rpc, []byte("<rpc "), nss, 1)
		rpc = bytes.Replace(rpc, []byte("<rpc>"), append(nss[:len(nss)-1], ">"...), 1)
	}

	if message.AppendXMLHeader {
		rpc = append([]byte(XMLHeader), rpc...)
	}

	return
}
func generateLock(config string) []byte {
	lockXML := []byte(`<lock><target><running/></target></lock>`)
	if config != "" {
		lockXML = bytes.Replace(lockXML, []byte("running"), []byte(config), 1)
	}

	return lockXML
}

func generateUnLock(config string) []byte {
	lockXML := []byte(`<unlock><target><running/></target></unlock>`)
	if config != "" {
		lockXML = bytes.Replace(lockXML, []byte("running"), []byte(config), 1)
	}

	return lockXML
}
