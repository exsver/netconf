package comware

import "encoding/xml"

type PoE struct {
	/* top level
		   PoE
	         Capabilities
	*/
	Capabilities *PoECapabilities `xml:"Capabilities"`
}

type PoECapabilities struct {
	XMLName     xml.Name `xml:"Capabilities"`
	SupportPoE  bool     `xml:"SupportPoE"`
	PSEPolicy   bool     `xml:"PSEPolicy"`
	PDPolicy    bool     `xml:"PDPolicy"`
	PowerManage bool     `xml:"PowerManage"`
	MaxPower    bool     `xml:"MaxPower"`
	PSEUpdate   bool     `xml:"PSEUpdate"`
	PowerMode   bool     `xml:"PowerMode"`
	PSEMaxPower bool     `xml:"PSEMaxPower"`
}
