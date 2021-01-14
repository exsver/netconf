package comware

import "encoding/xml"

type PortSecurity struct {
	/* top level
	PortSecurity
	  Common
	  Interfaces
	    []Interface
	*/

	Common     *PortSecurityCommon     `xml:"Common"`
	Interfaces *PortSecurityInterfaces `xml:"Interfaces"`
}

type PortSecurityCommon struct {
	XMLName                   xml.Name `xml:"Common"`
	Enable                    bool     `xml:"Enable"`
	MACMove                   bool     `xml:"MACMove,omitempty"`
	AuthorStrict              bool     `xml:"AuthorStrict,omitempty"`
	SecMACAgingInterval       int      `xml:"SecMACAgingInterval,omitempty"`
	IntrusionShutdownInterval int      `xml:"IntrusionShutdownInterval,omitempty"`
	SecMACCnt                 int      `xml:"SecMACCnt,omitempty"`
}

type PortSecurityInterfaces struct {
	Interfaces []PortSecurityInterface `xml:"Interface"`
}

type PortSecurityInterface struct {
	XMLName          xml.Name `xml:"Interface"`
	IfIndex          int      `xml:"IfIndex"`
	AuthenMode       int      `xml:"AuthenMode,omitempty"`
	ProtectionMode   int      `xml:"ProtectionMode,omitempty"`
	SecMACAgeMode    int      `xml:"SecMACAgeMode,omitempty"`
	SecMACCurrentCnt int      `xml:"SecMACCurrentCnt,omitempty"`
	NTKMode          int      `xml:"NTKMode,omitempty"`
	AuthorIgnore     bool     `xml:"AuthorIgnore,omitempty"`
	SecMACSticky     bool     `xml:"SecMACSticky,omitempty"`
	SecMACDynamic    bool     `xml:"SecMACDynamic,omitempty"`
}
