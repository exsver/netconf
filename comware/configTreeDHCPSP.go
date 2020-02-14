package comware

import "encoding/xml"

type DHCPSP struct {
	/* top level
	DHCPSP
	  DHCPSPBindingDatabase
	  DHCPSPConfig
	  DHCPSPInterface
	    []Interface
	  DHCPSPOpt82
	    []Option82
	  DHCPSPSpecification
	*/
	DHCPSPBindingDatabase *DHCPSPBindingDatabase `xml:"DHCPSPBindingDatabase"`
	DHCPSPConfig          *DHCPSPConfig          `xml:"DHCPSPConfig"`
	DHCPSPInterface       *DHCPSPInterface       `xml:"DHCPSPInterface"`
}

type DHCPSPBindingDatabase struct {
	XMLName        xml.Name `xml:"DHCPSPBindingDatabase"`
	UpdateInterval int      `xml:"UpdateInterval"`
}

type DHCPSPConfig struct {
	XMLName      xml.Name `xml:"DHCPSPConfig"`
	DHCPSPEnable bool     `xml:"DHCPSPEnable"`
}

type DHCPSPInterface struct {
	XMLName    xml.Name                `xml:"DHCPSPInterface"`
	Interfaces []DHCPSnoopingInterface `xml:"Interface"`
}

type DHCPSnoopingInterface struct {
	XMLName             xml.Name `xml:"Interface"`
	IfIndex             int      `xml:"IfIndex"`
	BindingRecord       bool     `xml:"BindingRecord"`
	CheckMacAddress     bool     `xml:"CheckMacAddress"`
	CheckRequestMessage bool     `xml:"CheckRequestMessage"`
	Trust               bool     `xml:"Trust"`
	LearnMaxNum         int      `xml:"LearnMaxNum"`
	RateLimitNum        int      `xml:"RateLimitNum"`
}
