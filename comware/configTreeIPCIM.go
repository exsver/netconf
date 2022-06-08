package comware

import "encoding/xml"

type IPCIM struct {
	/* top level
	   IPCIM
	     IpSourceBindingInterface
	       []SourceBinding
	     IpVerifySource
	       []VerifySource
	*/
	IPSourceBindingInterface *IPSourceBindingInterface `xml:"IpSourceBindingInterface"`
	IPVerifySource           *IPVerifySource           `xml:"IpVerifySource"`
}

// IPSourceBindingInterface table contains Ip Source Binding table information.
type IPSourceBindingInterface struct {
	SourceBindings []SourceBinding `xml:"SourceBinding"`
}

type SourceBinding struct {
	XMLName     xml.Name `xml:"SourceBinding"`
	IfIndex     string   `xml:"IfIndex"`
	Ipv4Address string   `xml:"Ipv4Address"`
	MacAddress  string   `xml:"MacAddress"`
	VLANID      string   `xml:"VLANID,omitempty"`
}

// IPVerifySource table contains Ip Verify Source table information.
type IPVerifySource struct {
	VerifySourceInterfaces []VerifySource `xml:"VerifySource"`
}

type VerifySource struct {
	XMLName          xml.Name `xml:"VerifySource"`
	IfIndex          int      `xml:"IfIndex"`
	VerifyIPAddress  bool     `xml:"VerifyIpAddress"`
	VerifyMacAddress bool     `xml:"VerifyMacAddress"`
}
