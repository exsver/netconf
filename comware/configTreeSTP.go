package comware

import "encoding/xml"

type STP struct {
	/* top level
	   STP
	     Base
	     Region
	     Interfaces
	       []Interface
	*/
	Base       *STPBase                    `xml:"Base"`
	Interfaces *STPInterfacesConfiguration `xml:"Interfaces"`
}

type STPBase struct {
	XMLName xml.Name `xml:"Base"`
	// Mode - The spanning-tree working modes:
	//  0 - STP
	//  2 - RSTP
	//  3 - MSTP
	//  4 - PVST
	Mode int `xml:"Mode"`
	// The max TCs that will be processed within the TC-protection period.
	// values: 1 - 255
	TcThreshold int `xml:"TcThreshold"`
	// PathCostMethod - The path cost calculation method:
	//  0 - Legacy
	//  1 - IEEE 802.1D-1998
	//  2 - IEEE 802.1t
	PathCostMethod int  `xml:"PathCostMethod,omitempty"`
	HelloTime      int  `xml:"HelloTime,omitempty"`    // The intervals in seconds. Valid values are: 1-10.
	MaxHops        int  `xml:"MaxHops,omitempty"`      // Valid values are: 1-40.
	MaxAge         int  `xml:"MaxAge,omitempty"`       // Valid values are: 6-40.
	ForwardDelay   int  `xml:"ForwardDelay,omitempty"` // The forward delay timer in seconds. Valid values are: 4-30.
	TcSnooping     bool `xml:"TcSnooping"`
	DigestSnooping bool `xml:"DigestSnooping"`
	BPDUProtect    bool `xml:"BPDUProtect"`
	TcProtect      bool `xml:"TcProtect"`
	Enable         bool `xml:"Enable"`
}

// STPInterfacesConfiguration table contains information about interface-level STP functions.
type STPInterfacesConfiguration struct {
	XMLName                    xml.Name                    `xml:"Interfaces"`
	STPInterfacesConfiguration []STPInterfaceConfiguration `xml:"Interface"`
}

type STPInterfaceConfiguration struct {
	XMLName           xml.Name `xml:"Interface"`
	IfIndex           int      `xml:"IfIndex"`
	PointToPoint      int      `xml:"PointToPoint,omitempty"`      //1 - Force-true; 2 - Force-false; 3 - Auto-negotiated by the link;
	TransmitHoldCount int      `xml:"TransmitHoldCount,omitempty"` //Valid values are: 1-255.
	Enable            bool     `xml:"Enable,omitempty"`
	EdgedPort         bool     `xml:"EdgedPort,omitempty"`
	RootProtect       bool     `xml:"RootProtect,omitempty"`
	LoopProtect       bool     `xml:"LoopProtect,omitempty"`
	RoleRestrict      bool     `xml:"RoleRestrict,omitempty"`
	TcRestrict        bool     `xml:"TcRestrict,omitempty"`
	DigestSnooping    bool     `xml:"DigestSnooping,omitempty"`
}
