package netconf

import (
	"bytes"
	"encoding/xml"
	"errors"
	"reflect"
	"testing"
)

func TestMarshalRPCMessage(t *testing.T) {
	cases := []struct {
		rpcStruct RPCMessage
		rpc       []byte //result
		err       error  //result
	}{
		{rpcStruct: RPCMessage{InnerXML: []byte("<get-sessions/>")}, rpc: []byte(`<rpc><get-sessions/></rpc>`), err: nil},
		{rpcStruct: RPCMessage{InnerXML: []byte("<get-sessions/>"), MessageID: "10"}, rpc: []byte(`<rpc message-id="10"><get-sessions/></rpc>`), err: nil},
		{rpcStruct: RPCMessage{InnerXML: []byte("<get-sessions/>"), MessageID: "10", CustomAttrs: []string{`xmlns:hp="http://www.hp.com/netconf/base:1.0"`}}, rpc: []byte(`<rpc xmlns:hp="http://www.hp.com/netconf/base:1.0" message-id="10"><get-sessions/></rpc>`), err: nil},
		{rpcStruct: RPCMessage{InnerXML: []byte("<get-sessions/>"), MessageID: "10", CustomAttrs: []string{`xmlns:data="http://www.hp.com/netconf/data:1.0"`, `xmlns:config="http://www.hp.com/netconf/config:1.0"`}}, rpc: []byte(`<rpc xmlns:data="http://www.hp.com/netconf/data:1.0" xmlns:config="http://www.hp.com/netconf/config:1.0" message-id="10"><get-sessions/></rpc>`), err: nil},
		{rpcStruct: RPCMessage{InnerXML: []byte("<get-sessions/>"), MessageID: "10", CustomAttrs: []string{`xmlns:data="http://www.hp.com/netconf/data:1.0"`, `xmlns:config="http://www.hp.com/netconf/config:1.0"`}, Xmlns: []string{BaseURI}}, rpc: []byte(`<rpc xmlns:data="http://www.hp.com/netconf/data:1.0" xmlns:config="http://www.hp.com/netconf/config:1.0" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="10"><get-sessions/></rpc>`), err: nil},
		{rpcStruct: RPCMessage{InnerXML: []byte("<get-sessions/>"), MessageID: "10", AppendXMLHeader: true}, rpc: []byte(`<?xml version="1.0" encoding="UTF-8"?><rpc message-id="10"><get-sessions/></rpc>`), err: nil},
	}

	for _, testCase := range cases {
		rpc, err := testCase.rpcStruct.MarshalRPCMessage()
		if !reflect.DeepEqual(rpc, testCase.rpc) || err != testCase.err {
			t.Errorf("Got: rpc: %s err: %v", rpc, err)
		}
	}
}

func TestUnmarshalRpcReply(t *testing.T) {
	cases := []struct {
		rawRpcReply []byte
		rpcReply    RPCReply //result
		err         error    //result
	}{
		{ //GetHostame
			rawRpcReply: []byte(`<?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="101"><data><top xmlns="http://www.hp.com/netconf/config:1.0"><Device><Base><HostName>hp-5940-1</HostName></Base></Device></top></data></rpc-reply>`),
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "101",
				Content:   []byte(`<data><top xmlns="http://www.hp.com/netconf/config:1.0"><Device><Base><HostName>hp-5940-1</HostName></Base></Device></top></data>`)},
			err: nil,
		},
		{ //GetSessions
			rawRpcReply: []byte(`<?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="102"><get-sessions><Session><SessionID>1</SessionID><Line>vty2</Line><UserName>s3rj1k</UserName><Since>2018-05-05T14:19:45</Since><LockHeld>false</LockHeld></Session></get-sessions></rpc-reply>`),
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "102",
				Content:   []byte(`<get-sessions><Session><SessionID>1</SessionID><Line>vty2</Line><UserName>s3rj1k</UserName><Since>2018-05-05T14:19:45</Since><LockHeld>false</LockHeld></Session></get-sessions>`)},
			err: nil,
		},
		{ //OK-message test1
			rawRpcReply: []byte(`<?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="2"><ok/></rpc-reply>`),
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "2",
				Content:   []byte(`<ok/>`),
				OK:        true},
			err: nil,
		},
		{ //OK-message test2
			rawRpcReply: []byte(`<?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="5"><some-tag/></rpc-reply>`),
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "5",
				Content:   []byte(`<some-tag/>`),
				OK:        false},
			err: nil,
		},
		{ //OK-message test3 (ok with vertical tabs - juniper)
			rawRpcReply: []byte{
				0x3c, 0x72, 0x70, 0x63, 0x2d, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, //|<rpc-reply xmlns|
				0x3a, 0x6a, 0x75, 0x6e, 0x6f, 0x73, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x78, //|:junos="http://x|
				0x6d, 0x6c, 0x2e, 0x6a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x6a, //|ml.juniper.net/j|
				0x75, 0x6e, 0x6f, 0x73, 0x2f, 0x31, 0x35, 0x2e, 0x31, 0x58, 0x35, 0x33, 0x2f, 0x6a, 0x75, 0x6e, //|unos/15.1X53/jun|
				0x6f, 0x73, 0x22, 0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3d, 0x22, 0x75, 0x72, 0x6e, 0x3a, 0x69, //|os" xmlns="urn:i|
				0x65, 0x74, 0x66, 0x3a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x78, 0x6d, 0x6c, 0x3a, 0x6e, //|etf:params:xml:n|
				0x73, 0x3a, 0x6e, 0x65, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x3a, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x31, //|s:netconf:base:1|
				0x2e, 0x30, 0x22, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x69, 0x64, 0x3d, 0x22, //|.0" message-id="|
				0x32, 0x22, 0x3e, 0x0a, 0x3c, 0x6f, 0x6b, 0x2f, 0x3e, 0x0a, 0x3c, 0x2f, 0x72, 0x70, 0x63, 0x2d, //|2">.<ok/>.</rpc-|
				0x72, 0x65, 0x70, 0x6c, 0x79, 0x3e}, //|reply>|
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "2",
				Content:   []byte{0x0a, 0x3c, 0x6f, 0x6b, 0x2f, 0x3e, 0x0a},
				OK:        true},
			err: nil,
		},
		{ //CDATA CLI
			//<?xml version="1.0" encoding="UTF-8"?>
			//  <rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="109">
			//    <CLI>
			//      <Execution><![CDATA[
			//        <hp-5940-1>display vlan....
			//        Total VLANs: 8....
			//        The VLANs include:....
			//        1(default), 99, 111, 208, 220, 855, 960, 3333, .. ....]]></Execution>
			//    </CLI>
			//  </rpc-reply>
			rawRpcReply: []byte{
				0x3c, 0x3f, 0x78, 0x6d, 0x6c, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3d, 0x22, 0x31, //|<?xml version="1|
				0x2e, 0x30, 0x22, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3d, 0x22, 0x55, 0x54, //|.0" encoding="UT|
				0x46, 0x2d, 0x38, 0x22, 0x3f, 0x3e, 0x3c, 0x72, 0x70, 0x63, 0x2d, 0x72, 0x65, 0x70, 0x6c, 0x79, //|F-8"?><rpc-reply|
				0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3d, 0x22, 0x75, 0x72, 0x6e, 0x3a, 0x69, 0x65, 0x74, 0x66, //| xmlns="urn:ietf|
				0x3a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x78, 0x6d, 0x6c, 0x3a, 0x6e, 0x73, 0x3a, 0x6e, //|:params:xml:ns:n|
				0x65, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x3a, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x31, 0x2e, 0x30, 0x22, //|etconf:base:1.0"|
				0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x69, 0x64, 0x3d, 0x22, 0x31, 0x30, 0x39, //| message-id="109|
				0x22, 0x3e, 0x3c, 0x43, 0x4c, 0x49, 0x3e, 0x3c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, //|"><CLI><Executio|
				0x6e, 0x3e, 0x3c, 0x21, 0x5b, 0x43, 0x44, 0x41, 0x54, 0x41, 0x5b, 0x3c, 0x68, 0x70, 0x2d, 0x35, //|n><![CDATA[<hp-5|
				0x39, 0x34, 0x30, 0x2d, 0x31, 0x3e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x20, 0x76, 0x6c, //|940-1>display vl|
				0x61, 0x6e, 0x0d, 0x0d, 0x0d, 0x0a, 0x20, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x20, 0x56, 0x4c, 0x41, //|an.... Total VLA|
				0x4e, 0x73, 0x3a, 0x20, 0x38, 0x0d, 0x0d, 0x0d, 0x0a, 0x20, 0x54, 0x68, 0x65, 0x20, 0x56, 0x4c, //|Ns: 8.... The VL|
				0x41, 0x4e, 0x73, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x3a, 0x0d, 0x0d, 0x0d, 0x0a, //|ANs include:....|
				0x20, 0x31, 0x28, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x29, 0x2c, 0x20, 0x39, 0x39, 0x2c, //| 1(default), 99,|
				0x20, 0x31, 0x31, 0x31, 0x2c, 0x20, 0x32, 0x30, 0x38, 0x2c, 0x20, 0x32, 0x32, 0x30, 0x2c, 0x20, //| 111, 208, 220, |
				0x38, 0x35, 0x35, 0x2c, 0x20, 0x39, 0x36, 0x30, 0x2c, 0x20, 0x33, 0x33, 0x33, 0x33, 0x2c, 0x20, //|855, 960, 3333, |
				0x08, 0x08, 0x20, 0x0d, 0x0d, 0x0d, 0x0a, 0x5d, 0x5d, 0x3e, 0x3c, 0x2f, 0x45, 0x78, 0x65, 0x63, //|.. ....]]></Exec|
				0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3e, 0x3c, 0x2f, 0x43, 0x4c, 0x49, 0x3e, 0x3c, 0x2f, 0x72, 0x70, //|ution></CLI></rp|
				0x63, 0x2d, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x3e, //|c-reply>|
			},
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "109",
				Content: []byte{
					0x3c, 0x43, 0x4c, 0x49, 0x3e, 0x3c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
					0x6e, 0x3e, 0x3c, 0x21, 0x5b, 0x43, 0x44, 0x41, 0x54, 0x41, 0x5b, 0x3c, 0x68, 0x70, 0x2d, 0x35,
					0x39, 0x34, 0x30, 0x2d, 0x31, 0x3e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x20, 0x76, 0x6c,
					0x61, 0x6e, 0x0d, 0x0d, 0x0d, 0x0a, 0x20, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x20, 0x56, 0x4c, 0x41,
					0x4e, 0x73, 0x3a, 0x20, 0x38, 0x0d, 0x0d, 0x0d, 0x0a, 0x20, 0x54, 0x68, 0x65, 0x20, 0x56, 0x4c,
					0x41, 0x4e, 0x73, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x3a, 0x0d, 0x0d, 0x0d, 0x0a,
					0x20, 0x31, 0x28, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x29, 0x2c, 0x20, 0x39, 0x39, 0x2c,
					0x20, 0x31, 0x31, 0x31, 0x2c, 0x20, 0x32, 0x30, 0x38, 0x2c, 0x20, 0x32, 0x32, 0x30, 0x2c, 0x20,
					0x38, 0x35, 0x35, 0x2c, 0x20, 0x39, 0x36, 0x30, 0x2c, 0x20, 0x33, 0x33, 0x33, 0x33,
					0x20, 0x0d, 0x0d, 0x0d, 0x0a, 0x5d, 0x5d, 0x3e, 0x3c, 0x2f, 0x45, 0x78, 0x65, 0x63,
					0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3e, 0x3c, 0x2f, 0x43, 0x4c, 0x49, 0x3e,
				}},
			err: nil,
		},
		{ //HPE RPC Error:
			//<?xml version="1.0" encoding="UTF-8"?>
			//<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="102">
			//  <rpc-error xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
			//    <error-type>rpc</error-type>
			//    <error-tag>unknown-element</error-tag>
			//    <error-severity>error</error-severity>
			//    <error-message xml:lang="en">Unexpected element 'urn:ietf:params:xml:ns:netconf:base:1.0':'get-sessionss' under element '/rpc</error-message>
			//    <error-info>
			//      <bad-element>get-sessionss</bad-element>
			//    </error-info>
			//  </rpc-error>
			//</rpc-reply>
			rawRpcReply: []byte(`<?xml version="1.0" encoding="UTF-8"?><rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="102"><rpc-error xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><error-type>rpc</error-type><error-tag>unknown-element</error-tag><error-severity>error</error-severity><error-message xml:lang="en">Unexpected element 'urn:ietf:params:xml:ns:netconf:base:1.0':'get-sessionss' under element '/rpc</error-message><error-info><bad-element>get-sessionss</bad-element></error-info></rpc-error></rpc-reply>`),
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "102",
				Errors: []RPCError{
					{XMLName: xml.Name{Space: "urn:ietf:params:xml:ns:netconf:base:1.0", Local: "rpc-error"},
						ErrorType:     "rpc",
						ErrorTag:      "unknown-element",
						ErrorSeverity: "error",
						ErrorMessage:  "Unexpected element 'urn:ietf:params:xml:ns:netconf:base:1.0':'get-sessionss' under element '/rpc",
						ErrorInfo:     RPCErrorInfo{[]byte(`<bad-element>get-sessionss</bad-element>`)},
					},
				},
				Content: []byte(`<rpc-error xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><error-type>rpc</error-type><error-tag>unknown-element</error-tag><error-severity>error</error-severity><error-message xml:lang="en">Unexpected element 'urn:ietf:params:xml:ns:netconf:base:1.0':'get-sessionss' under element '/rpc</error-message><error-info><bad-element>get-sessionss</bad-element></error-info></rpc-error>`)},
			err: nil,
		},
		{ /*The following <rpc-reply> illustrates the case of returning multiple <rpc-error> elements (Example from RFC6241).:
			<rpc-reply message-id="101"
			  xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"
			  xmlns:xc="urn:ietf:params:xml:ns:netconf:base:1.0">
			  <rpc-error>
			    <error-type>application</error-type>
			    <error-tag>invalid-value</error-tag>
			    <error-severity>error</error-severity>
			    <error-path xmlns:t="http://example.com/schema/1.2/config">/t:top/t:interface[t:name="Ethernet0/0"]/t:mtu</error-path>
			    <error-message xml:lang="en">MTU value 25000 is not within range 256..9192</error-message>
			  </rpc-error>
			  <rpc-error>
			    <error-type>application</error-type>
			    <error-tag>invalid-value</error-tag>
			    <error-severity>error</error-severity>
			    <error-path xmlns:t="http://example.com/schema/1.2/config">/t:top/t:interface[t:name="Ethernet1/0"]/t:address/t:name</error-path>
			    <error-message xml:lang="en">Invalid IP address for interface Ethernet1/0</error-message>
			  </rpc-error>
			</rpc-reply>*/
			rawRpcReply: []byte(`<rpc-reply message-id="101" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:xc="urn:ietf:params:xml:ns:netconf:base:1.0"><rpc-error><error-type>application</error-type><error-tag>invalid-value</error-tag><error-severity>error</error-severity><error-path xmlns:t="http://example.com/schema/1.2/config">/t:top/t:interface[t:name="Ethernet0/0"]/t:mtu</error-path><error-message xml:lang="en">MTU value 25000 is not within range 256..9192</error-message></rpc-error><rpc-error><error-type>application</error-type><error-tag>invalid-value</error-tag><error-severity>error</error-severity><error-path xmlns:t="http://example.com/schema/1.2/config">/t:top/t:interface[t:name="Ethernet1/0"]/t:address/t:name</error-path><error-message xml:lang="en">Invalid IP address for interface Ethernet1/0</error-message></rpc-error></rpc-reply>`),
			rpcReply: RPCReply{
				XMLName:   xml.Name{Local: "rpc-reply", Space: "urn:ietf:params:xml:ns:netconf:base:1.0"},
				MessageID: "101",
				Errors: []RPCError{
					{XMLName: xml.Name{Space: "urn:ietf:params:xml:ns:netconf:base:1.0", Local: "rpc-error"},
						ErrorType:     "application",
						ErrorTag:      "invalid-value",
						ErrorSeverity: "error",
						ErrorPath:     `/t:top/t:interface[t:name="Ethernet0/0"]/t:mtu`,
						ErrorMessage:  "MTU value 25000 is not within range 256..9192",
					},
					{XMLName: xml.Name{Space: "urn:ietf:params:xml:ns:netconf:base:1.0", Local: "rpc-error"},
						ErrorType:     "application",
						ErrorTag:      "invalid-value",
						ErrorSeverity: "error",
						ErrorPath:     `/t:top/t:interface[t:name="Ethernet1/0"]/t:address/t:name`,
						ErrorMessage:  "Invalid IP address for interface Ethernet1/0",
					},
				},
				Content: []byte(`<rpc-error><error-type>application</error-type><error-tag>invalid-value</error-tag><error-severity>error</error-severity><error-path xmlns:t="http://example.com/schema/1.2/config">/t:top/t:interface[t:name="Ethernet0/0"]/t:mtu</error-path><error-message xml:lang="en">MTU value 25000 is not within range 256..9192</error-message></rpc-error><rpc-error><error-type>application</error-type><error-tag>invalid-value</error-tag><error-severity>error</error-severity><error-path xmlns:t="http://example.com/schema/1.2/config">/t:top/t:interface[t:name="Ethernet1/0"]/t:address/t:name</error-path><error-message xml:lang="en">Invalid IP address for interface Ethernet1/0</error-message></rpc-error>`)},
			err: nil,
		},
	}
	for _, testCase := range cases {
		rpc, err := UnmarshalRpcReply(testCase.rawRpcReply)
		if !reflect.DeepEqual(rpc, &testCase.rpcReply) || err != testCase.err {
			t.Errorf("\n Got:  rpcReply: %s err: %v\n Want: rpcReply: %s err: %v", rpc, err, testCase.rpcReply, testCase.err)
		}
	}
}

func TestCorrectBackspaces(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{ //2 backspaces
			caseDescription: "1",
			in:              []byte{0x53, 0x53, 0x8, 0x53, 0x53, 0x53, 0x53, 0x53, 0x8, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53},
			out:             []byte{0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53, 0x53},
		},
		{
			caseDescription: "2",
			in:              []byte{0x50, 0x51, 0x52, 0x8, 0x8, 0x53, 0x54, 0x55},
			out:             []byte{0x50, 0x53, 0x54, 0x55},
		},
		{
			caseDescription: "3",
			in:              []byte{0x50, 0x51, 0x52, 0x53, 0x54, 0x8, 0x8, 0x8, 0x8, 0x8, 0x55, 0x56},
			out:             []byte{0x55, 0x56},
		},
		{
			caseDescription: "4",
			in:              []byte{0x50, 0x51, 0x52, 0x8},
			out:             []byte{0x50, 0x51},
		},
		{
			caseDescription: "5",
			in:              []byte{0x50, 0x51, 0x52, 0x8, 0x8},
			out:             []byte{0x50},
		},
		{
			caseDescription: "6",
			in:              []byte{0x50, 0x51, 0x52, 0x8, 0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "7",
			in:              []byte{0x8, 0x50, 0x51},
			out:             []byte{0x50, 0x51},
		},
		{
			caseDescription: "8",
			in:              []byte{0x8, 0x8, 0x50, 0x51},
			out:             []byte{0x50, 0x51},
		},
		{
			caseDescription: "9",
			in:              []byte{0x8, 0x8, 0x50, 0x51, 0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "10",
			in:              []byte{0x8, 0x8, 0x50, 0x51, 0x52, 0x8, 0x8},
			out:             []byte{0x50},
		},
		{
			caseDescription: "11",
			in:              []byte{0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "12",
			in:              []byte{0x50, 0x51, 0x8, 0x8, 0x8, 0x8},
			out:             []byte{},
		},
		{
			caseDescription: "13",
			in:              []byte{0x8, 0x8, 0x50},
			out:             []byte{0x50},
		},
		{
			caseDescription: "14",
			in:              []byte{0x41, 0x42, 0x8, 0x43, 0x44, 0x45, 0x8, 0x46, 0x47, 0x48, 0x49, 0x50, 0x51, 0x52, 0x53, 0x54, 0x8, 0x8, 0x8, 0x55, 0x56, 0x57, 0x8, 0x58, 0x59, 0x8},
			out:             []byte{0x41, 0x43, 0x44, 0x46, 0x47, 0x48, 0x49, 0x50, 0x51, 0x55, 0x56, 0x58},
		},
		{
			caseDescription: "15",
			in:              []byte{0x50, 0x51, 0x52},
			out:             []byte{0x50, 0x51, 0x52},
		},
	}
	for _, testCase := range cases {
		out := CorrectBackspaces(testCase.in)
		if !reflect.DeepEqual(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{
			caseDescription: "1",
			in: []byte(`<rpc>
                          <validate>
                            <source>
                              <candidate/>
                            </source>
                          </validate>
                        </rpc>`),
			out: []byte(`<rpc><validate><source><candidate/></source></validate></rpc>`),
		},
		{
			caseDescription: "2",
			in: []byte(`
                       <rpc message-id="102" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
                          <get-sessions/>
                       </rpc>`),
			out: []byte(`<rpc message-id="102" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0"><get-sessions/></rpc>`),
		},
		{
			caseDescription: "3",
			in: []byte(`
<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos">
<ping-results xmlns="http://xml.juniper.net/junos/15.1X53/junos-probe-tests">
<target-host>
8.8.8.8
</target-host>
<target-ip>
8.8.8.8
</target-ip>
<packet-size>
56
</packet-size>
<probe-result date-determined="1527113979">
<probe-index>
1
</probe-index>
<probe-success/>
<sequence-number>
0
</sequence-number>
<ip-address>
8.8.8.8
</ip-address>
<time-to-live>
48
</time-to-live>
<response-size>
64
</response-size>
<rtt>
33099
</rtt>
</probe-result>
<probe-results-summary>
<probes-sent>
1
</probes-sent>
<responses-received>
1
</responses-received>
<packet-loss>
0
</packet-loss>
<rtt-minimum>
33099
</rtt-minimum>
<rtt-maximum>
33099
</rtt-maximum>
<rtt-average>
33099
</rtt-average>
<rtt-stddev>
0
</rtt-stddev>
</probe-results-summary>
<ping-success/>
</ping-results>
</rpc-reply>`),
			out: []byte(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos"><ping-results xmlns="http://xml.juniper.net/junos/15.1X53/junos-probe-tests"><target-host>8.8.8.8</target-host><target-ip>8.8.8.8</target-ip><packet-size>56</packet-size><probe-result date-determined="1527113979"><probe-index>1</probe-index><probe-success/><sequence-number>0</sequence-number><ip-address>8.8.8.8</ip-address><time-to-live>48</time-to-live><response-size>64</response-size><rtt>33099</rtt></probe-result><probe-results-summary><probes-sent>1</probes-sent><responses-received>1</responses-received><packet-loss>0</packet-loss><rtt-minimum>33099</rtt-minimum><rtt-maximum>33099</rtt-maximum><rtt-average>33099</rtt-average><rtt-stddev>0</rtt-stddev></probe-results-summary><ping-success/></ping-results></rpc-reply>`),
		},
	}
	for _, testCase := range cases {
		out := Normalize(testCase.in)
		if !bytes.Equal(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestConvertToPairedTags(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{
			caseDescription: "1", //simple changes
			in:              []byte("<if-device-flags><ifdf-present/><ifdf-running/></if-device-flags><if-config-flags><iff-snmp-traps/></if-config-flags>"),
			out:             []byte("<if-device-flags><ifdf-present>true</ifdf-present><ifdf-running>true</ifdf-running></if-device-flags><if-config-flags><iff-snmp-traps>true</iff-snmp-traps></if-config-flags>"),
		},
		{
			caseDescription: "2", //simple changes #2
			in: []byte(`
									<arp-table-entry>
    									<mac-address>d8:50:e6:ae:67:28</mac-address>
    									<ip-address>172.21.1.128</ip-address>
										<hostname>172.21.1.128</hostname>
										<interface-name>em0.0</interface-name>
    									<arp-table-entry-flags>
        									<none/>
    									</arp-table-entry-flags>
									</arp-table-entry>`),
			out: []byte(`
									<arp-table-entry>
    									<mac-address>d8:50:e6:ae:67:28</mac-address>
    									<ip-address>172.21.1.128</ip-address>
										<hostname>172.21.1.128</hostname>
										<interface-name>em0.0</interface-name>
    									<arp-table-entry-flags>
        									<none>true</none>
    									</arp-table-entry-flags>
									</arp-table-entry>`),
		},
		{
			caseDescription: "3", //No change
			in: []byte(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos">
                                       <system-information>
                                         <hardware-model>qfx5200-32c-32q</hardware-model>
                                         <os-name>junos-qfx</os-name>
                                         <os-version>15.1X53-D30.5</os-version>
                                         <serial-number>WH0000000000</serial-number>
                                         <host-name>test-qfx5200</host-name>
                                       </system-information>
                                     </rpc-reply>`),
			out: []byte(`<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" xmlns:junos="http://xml.juniper.net/junos/15.1X53/junos">
                                       <system-information>
                                         <hardware-model>qfx5200-32c-32q</hardware-model>
                                         <os-name>junos-qfx</os-name>
                                         <os-version>15.1X53-D30.5</os-version>
                                         <serial-number>WH0000000000</serial-number>
                                         <host-name>test-qfx5200</host-name>
                                       </system-information>
                                     </rpc-reply>`),
		},
		{
			caseDescription: "4", //Multilevel
			in: []byte(`<level1>
                                       <tagA/>
                                       <tagB/>
                                       <tagC/>
                                       <level2>
                                         <tagAA/>
                                         <tagBB/>
                                         <tagCC/>
                                         <level3>
                                           <tagAAA/>
                                           <tagBBB/>
                                           <tagCCC/>
                                         </level3>
                                       </level2>
                                     </level1>`),
			out: []byte(`<level1>
                                       <tagA>true</tagA>
                                       <tagB>true</tagB>
                                       <tagC>true</tagC>
                                       <level2>
                                         <tagAA>true</tagAA>
                                         <tagBB>true</tagBB>
                                         <tagCC>true</tagCC>
                                         <level3>
                                           <tagAAA>true</tagAAA>
                                           <tagBBB>true</tagBBB>
                                           <tagCCC>true</tagCCC>
                                         </level3>
                                       </level2>
                                     </level1>`),
		},
	}
	for _, testCase := range cases {
		out := ConvertToPairedTags(testCase.in)
		if !bytes.Equal(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestConvertToXML(t *testing.T) {
	cases := []struct {
		caseDescription string
		in              []byte
		out             []byte
	}{
		{
			caseDescription: "1",
			in:              []byte(`level1`),
			out:             []byte(`<level1/>`),
		},
		{
			caseDescription: "2",
			in:              []byte(`level1/level2`),
			out:             []byte(`<level1><level2/></level1>`),
		},
		{
			caseDescription: "3",
			in:              []byte(`level1/level2/level3`),
			out:             []byte(`<level1><level2><level3/></level2></level1>`),
		},
	}
	for _, testCase := range cases {
		out := ConvertToXML(testCase.in)
		if !bytes.Equal(out, testCase.out) {
			t.Errorf("CaseDescription: %s\n Got :%s\n Want:%s", testCase.caseDescription, out, testCase.out)
		}
	}
}

func TestRPCReply_Error(t *testing.T) {
	cases := []struct {
		caseDescription string
		rpcReply        RPCReply
		err             error
	}{
		/*	{
			caseDescription: "1",
			rpcReply:        RPCReply{},
			err:             nil,
		}, */
		{
			caseDescription: "2",
			rpcReply: RPCReply{
				Errors: []RPCError{
					{
						ErrorMessage: "Some error message",
					},
				},
			},
			err: errors.New("ERROR-MESSAGE: Some error message \n"),
		},
		{
			caseDescription: "3",
			rpcReply: RPCReply{
				Errors: []RPCError{
					{
						ErrorType:     "protocol",
						ErrorTag:      "operation-failed",
						ErrorSeverity: "error",
						ErrorPath:     "[edit interfaces xe-0/0/1]",
						ErrorMessage:  "VLAN-ID can only be specified on tagged ethernet interfaces",
						ErrorInfo: RPCErrorInfo{
							Info: []byte("<bad-element>unit 0</bad-element>")},
					},
				},
			},
			err: errors.New("ERROR-TYPE: protocol ERROR-TAG: operation-failed ERROR-SEVERITY: error ERROR-PATH: [edit interfaces xe-0/0/1] ERROR-MESSAGE: VLAN-ID can only be specified on tagged ethernet interfaces ERROR-INFO: <bad-element>unit 0</bad-element> \n"),
		},
		{
			caseDescription: "4",
			rpcReply: RPCReply{
				Errors: []RPCError{
					{
						ErrorType:     "protocol",
						ErrorTag:      "operation-failed",
						ErrorSeverity: "error",
						ErrorPath:     "[edit interfaces xe-0/0/1]",
						ErrorMessage:  "VLAN-ID can only be specified on tagged ethernet interfaces",
						ErrorInfo: RPCErrorInfo{
							Info: []byte("<bad-element>unit 0</bad-element>")},
					},
					{
						ErrorType:     "protocol",
						ErrorTag:      "operation-failed",
						ErrorSeverity: "error",
						ErrorMessage:  "configuration check-out failed",
					},
				},
			},
			err: errors.New("ERROR-TYPE: protocol ERROR-TAG: operation-failed ERROR-SEVERITY: error ERROR-PATH: [edit interfaces xe-0/0/1] ERROR-MESSAGE: VLAN-ID can only be specified on tagged ethernet interfaces ERROR-INFO: <bad-element>unit 0</bad-element> \n" +
				"ERROR-TYPE: protocol ERROR-TAG: operation-failed ERROR-SEVERITY: error ERROR-MESSAGE: configuration check-out failed \n"),
		},
	}
	for _, testCase := range cases {
		out := testCase.rpcReply.Error()
		if out.Error() != testCase.err.Error() {
			t.Errorf("CaseDescription: %s\n Got :\n %s\n Want:\n %s", testCase.caseDescription, []byte(out.Error()), []byte(testCase.err.Error()))
		}
	}
}
