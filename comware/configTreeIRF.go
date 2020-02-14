package comware

import "encoding/xml"

type IRF struct {
	/* top level
	   IRF
	     Capability
	     Configuration
	     IRFPorts
	       []IRFPort
	     Members
	       []Member
	*/
	Capability    *IRFCapability    `xml:"Capability"`
	Configuration *IRFConfiguration `xml:"Configuration"`
	IRFPorts      *IRFPorts         `xml:"IRFPorts"`
	Members       *Members          `xml:"Members"`
}

type IRFCapability struct {
	XMLName            xml.Name `xml:"Capability"`
	MaxMemberCount     int      `xml:"MaxMemberCount"`
	MaxPriority        int      `xml:"MaxPriority"`
	MaxIfNumPerIRFPort int      `xml:"MaxIfNumPerIRFPort"`
}

type IRFConfiguration struct {
	XMLName     xml.Name `xml:"Configuration"`
	AutoUpgrade string   `xml:"AutoUpgrade"`
	Domain      int      `xml:"Domain"`
	LinkDelay   int      `xml:"LinkDelay"`
	MacPersist  int      `xml:"MacPersist"`
	BridgeMac   string   `xml:"BridgeMac"`
	TopoType    int      `xml:"TopoType"`
	StackMode   bool     `xml:"StackMode"`
	MemberCount int      `xml:"MemberCount"`
}

// IRFPorts table contains configuration of the IRF ports.
type IRFPorts struct {
	XMLName  xml.Name  `xml:"IRFPorts"`
	IRFPorts []IRFPort `xml:"IRFPort"`
}

type IRFPort struct {
	XMLName   xml.Name          `xml:"IRFPort"`
	MemberID  int               `xml:"MemberID,omitempty"`
	Port      int               `xml:"Port,omitempty"`
	Neighbor  int               `xml:"Neighbor"`
	State     int               `xml:"State"`
	Interface *IRFPortInterface `xml:"Interface"`
}

type IRFPortInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfName  string   `xml:"IfName"`
	//Physical IRF port mode
	// 1 Normal, 2 - Enhanced 3 - Extended
	Mode      int `xml:"Mode"`
	LinkState int `xml:"LinkState"`
}

// Members table contains configuration of the IRF members.
type Members struct {
	XMLName xml.Name `xml:"Members"`
	Members []Member `xml:"Member"`
}

type Member struct {
	XMLName     xml.Name `xml:"Member"`
	MemberID    int      `xml:"MemberID,omitempty"`
	Description string   `xml:"Description,omitempty"`
	NewMemberID int      `xml:"NewMemberID,omitempty"`
	Priority    int      `xml:"Priority,omitempty"`
}
