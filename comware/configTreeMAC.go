package comware

import "encoding/xml"

type MAC struct {
	/*top level
	MAC
	  MacAging
	  MacFwdSrcCheck
	    []FwdSrcCheck
	  MacUnicastTable
	    []MacTableEntry
	  MacPort
	    []PortLearn
	  MacAging
	  MacSpecification

	*/
	MacUnicastTable  *MacUnicastTable  `xml:"MacUnicastTable"`
	MacPort          *MacPort          `xml:"MacPort"`
	MacVLAN          *MacVLAN          `xml:"MacVLAN"`
	MacAging         *MacAging         `xml:"MacAging"`
	MacSpecification *MacSpecification `xml:"MacSpecification"`
}

// MacUnicastTable table contains unicast MAC address table information.
type MacUnicastTable struct {
	Unicast []MacTableEntry `xml:"Unicast"`
}

type MacTableEntry struct {
	XMLName    xml.Name `xml:"Unicast"`
	VLANID     string   `xml:"VLANID"`
	MacAddress string   `xml:"MacAddress"`
	PortIndex  int      `xml:"PortIndex"`
	Status     string   `xml:"Status"`
	Aging      string   `xml:"Aging"`
}

// MacPort table contains the information of MAC learning on an interface.
type MacPort struct {
	XMLName    xml.Name    `xml:"MacPort"`
	PortsLearn []PortLearn `xml:"PortLearn"`
}

type PortLearn struct {
	XMLName              xml.Name `xml:"PortLearn"`
	PortIndex            int      `xml:"PortIndex"`
	LearnEnable          bool     `xml:"LearnEnable,omitempty"`
	PortForwardingEnable bool     `xml:"PortForwardingEnable,omitempty"`
}

// MacVLAN table contains the information of MAC learning for a VLAN.
type MacVLAN struct {
	XMLName    xml.Name    `xml:"MacVLAN"`
	VLANsLearn []VLANLearn `xml:"VLANLearn"`
}

type VLANLearn struct {
	XMLName     xml.Name `xml:"VLANLearn"`
	VLANID      int      `xml:"VLANID"`
	LearnEnable bool     `xml:"LearnEnable"`
}

type MacAging struct {
	XMLName      xml.Name `xml:"MacAging"`
	AgingTimeMin int      `xml:"AgingTimeMin,omitempty"`
	AgingTimeMax int      `xml:"AgingTimeMax,omitempty"`
	AgingTime    int      `xml:"AgingTime,omitempty"`
}

type MacSpecification struct {
	XMLName                  xml.Name `xml:"MacSpecification"`
	PortLearnMaxNumLimit     int      `xml:"PortLearnMaxNumLimit,omitempty"`
	SupportMacGroup          bool     `xml:"SupportMacGroup,omitempty"`
	SupportMacVLANLearnLimit bool     `xml:"SupportMacVLANLearnLimit,omitempty"`
	SupportPortBridgeEnable  bool     `xml:"SupportPortBridgeEnable,omitempty"`
}