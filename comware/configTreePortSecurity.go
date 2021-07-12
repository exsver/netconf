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
	XMLName xml.Name `xml:"Common"`
	Enable  bool     `xml:"Enable"`
	// MACMove - Allows 802.1X or MAC authenticated users to move between ports on a device.
	MACMove bool `xml:"MACMove,omitempty"`
	// AuthorStrict - Allows 802.1X or MAC authenticated users to be online or not while authorizing failed.
	AuthorStrict bool `xml:"AuthorStrict,omitempty"`
	// SecMACAgingInterval - Period of secure MAC address aging timer.
	SecMACAgingInterval int `xml:"SecMACAgingInterval,omitempty"`
	// IntrusionShutdownInterval - Period value of Intrusion Protection's port shutdown timer.
	IntrusionShutdownInterval int `xml:"IntrusionShutdownInterval,omitempty"`
	SecMACCnt                 int `xml:"SecMACCnt,omitempty"`
}

type PortSecurityInterfaces struct {
	Interfaces []PortSecurityInterface `xml:"Interface"`
}

type PortSecurityInterface struct {
	XMLName xml.Name `xml:"Interface"`
	IfIndex int      `xml:"IfIndex"`
	// AuthenMode - Authentication Mode:
	//  1 - noRestrictions
	//  2 - autolearn
	//  3 - mac-authentication
	//  4 - mac-else-userlogin-secure
	//  5 - mac-else-userlogin-secure-ext
	//  6 - secure
	//  7 - userlogin
	//  8 - userlogin-secure
	//  9 - userlogin-secure-ext
	//  10 - userlogin-secure-or-mac
	//  11 - userlogin-secure-or-mac-ext
	//  12 - userlogin-withoui
	AuthenMode int `xml:"AuthenMode,omitempty"`
	// ProtectionMode - Protection action while intrusion detected:
	//  1 - no action
	//  2 - block user mac
	//  3 - shutdown port
	//  4 - shutdown port for a time
	ProtectionMode int `xml:"ProtectionMode,omitempty"`
	// SecMACAgeMode - Secure MAC aging mode:
	//  0 - Timeout
	//  1 - Inactivity
	SecMACAgeMode    int  `xml:"SecMACAgeMode,omitempty"`
	SecMACCurrentCnt int  `xml:"SecMACCurrentCnt,omitempty"`
	NTKMode          int  `xml:"NTKMode,omitempty"`
	AuthorIgnore     bool `xml:"AuthorIgnore,omitempty"`
	SecMACSticky     bool `xml:"SecMACSticky,omitempty"`
	SecMACDynamic    bool `xml:"SecMACDynamic,omitempty"`
}
