package comware

import "encoding/xml"

type LR struct {
	/* top level
	   LR
	     IfCapabilities
	       []IfCapability
	     Interfaces
	       []Interface
	*/
	IfCapabilities *LRIfCapabilities `xml:"IfCapabilities"`
	Interfaces     *LRInterfaces     `xml:"Interfaces"`
}

type LRIfCapabilities struct {
	XMLName        xml.Name         `xml:"IfCapabilities"`
	IfCapabilities []LRIfCapability `xml:"IfCapability"`
}

type LRIfCapability struct {
	XMLName        xml.Name `xml:"IfCapability"`
	IfIndex        int      `xml:"IfIndex"`
	CIRMin         int      `xml:"CIRMin"`
	CIRMax         int      `xml:"CIRMax"`
	CIRGranularity int      `xml:"CIRGranularity"`
	CBSMin         int      `xml:"CBSMin"`
	CBSMax         int      `xml:"CBSMax"`
	CBSGranularity int      `xml:"CBSGranularity"`
	EBSMin         int      `xml:"EBSMin"`
	EBSMax         int      `xml:"EBSMax"`
	EBSGranularity int      `xml:"EBSGranularity"`
	PIRMin         int      `xml:"PIRMin"`
	PIRMax         int      `xml:"PIRMax"`
	PIRGranularity int      `xml:"PIRGranularity"`
}

type LRInterfaces struct {
	XMLName    xml.Name      `xml:"Interfaces"`
	Interfaces []LRInterface `xml:"Interface"`
}

type LRInterface struct {
	XMLName        xml.Name `xml:"Interface"`
	IfIndex        int      `xml:"IfIndex"`
	Direction      int      `xml:"Direction"` // 0 - inbound, 1 - outbound
	CIR            int      `xml:"CIR"`
	CBS            int      `xml:"CBS"`
	PassedPkts     int      `xml:"PassedPkts"`
	DiscardedPkts  int      `xml:"DiscardedPkts"`
	DelayedPkts    int      `xml:"DelayedPkts"`
	PassedBytes    int      `xml:"PassedBytes"`
	DiscardedBytes int      `xml:"DiscardedBytes"`
	Active         bool     `xml:"Active"`
}
