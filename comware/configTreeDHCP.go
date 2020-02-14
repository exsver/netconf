package comware

import "encoding/xml"

type DHCP struct {
	/* top level
	DHCP
	  DHCPConfig
	  DHCPIfMode
	  DHCPServerIpPool
	    []IpPool
	  DHCPServerPoolStatic
	    []IPPoolStatic
	  DHCPServerIPInUse         ***ReadOnly***
	    []IpInUse               ***ReadOnly***
	*/
	DHCPConfig           *DHCPConfig           `xml:"DHCPConfig"`
	DHCPIfMode           *DHCPIfMode           `xml:"DHCPIfMode"`
	DHCPServerIPPool     *DHCPServerIPPool     `xml:"DHCPServerIpPool"`
	DHCPServerPoolStatic *DHCPServerPoolStatic `xml:"DHCPServerPoolStatic"`
	DHCPServerIPInUse    *DHCPServerIPInUse    `xml:"DHCPServerIpInUse"`
}

type DHCPConfig struct {
	XMLName      xml.Name          `xml:"DHCPConfig"`
	DHCPEnable   bool              `xml:"DHCPEnable,omitempty"`
	ServerConfig *DHCPServerConfig `xml:"ServerConfig"`
}

type DHCPServerConfig struct {
	XMLName           xml.Name `xml:"ServerConfig"`
	AlwaysBroadcast   bool     `xml:"AlwaysBroadcast,omitempty"`
	IgnoreBOOTP       bool     `xml:"IgnoreBOOTP,omitempty"`
	BOOTPReplyRFC1048 bool     `xml:"BOOTPReplyRFC1048,omitempty"`
	Opt82Enabled      bool     `xml:"Opt82Enabled,omitempty"`
	PingNumber        int      `xml:"PingNumber,omitempty"`  //Valid values are:0-10
	PingTimeout       int      `xml:"PingTimeout,omitempty"` //Valid values are:0-10000
}

type DHCPRelayConfig struct {
	XMLName           xml.Name `xml:"RelayConfig"`
	UserInfoRecord    bool     `xml:"UserInfoRecord,omitempty"`
	UserInfoRefresh   bool     `xml:"UserInfoRefresh,omitempty"`
	UserInfoFlushTime int      `xml:"UserInfoFlushTime,omitempty"` //Valid values are:0-120
}

// DHCPIfMode table contains DHCP interface mode information.
type DHCPIfMode struct {
	XMLName    xml.Name `xml:"DHCPIfMode"`
	Interfaces []IfMode `xml:"IfMode"`
}

type IfMode struct {
	XMLName xml.Name `xml:"IfMode"`
	IfIndex int      `xml:"IfIndex"`
	Mode    int      `xml:"Mode"`
}

// DHCPServerIPPool table contains DHCP server IP pool information.
type DHCPServerIPPool struct {
	XMLName xml.Name `xml:"DHCPServerIpPool"`
	IPPools []IPPool `xml:"IpPool"`
}

type IPPool struct {
	XMLName            xml.Name `xml:"IpPool"`
	PoolIndex          int      `xml:"PoolIndex"`
	PoolName           string   `xml:"PoolName,omitempty"`
	NetworkIpv4Address string   `xml:"NetworkIpv4Address,omitempty"`
	NetworkIpv4Mask    string   `xml:"NetworkIpv4Mask,omitempty"`
	GatewayIpv4Address string   `xml:"GatewayIpv4Address,omitempty"`
	DNSIpv4Address     string   `xml:"DNSIpv4Address,omitempty"`
	StartIpv4Address   string   `xml:"StartIpv4Address,omitempty"`
	EndIpv4Address     string   `xml:"EndIpv4Address,omitempty"`
	LeaseDay           int      `xml:"LeaseDay,omitempty"`
	LeaseHour          int      `xml:"LeaseHour,omitempty"`
	LeaseMinute        int      `xml:"LeaseMinute,omitempty"`
	LeaseSecond        int      `xml:"LeaseSecond,omitempty"`
	LeaseUnlimit       bool     `xml:"LeaseUnlimit,omitempty"`
}

// DHCPServerPoolStatic table contains static binding information of DHCP server IP pool.
type DHCPServerPoolStatic struct {
	XMLName     xml.Name       `xml:"DHCPServerPoolStatic"`
	StaticHosts []IPPoolStatic `xml:"IPPoolStatic"`
}

type IPPoolStatic struct {
	XMLName     xml.Name `xml:"IPPoolStatic"`
	PoolIndex   int      `xml:"PoolIndex"`
	Ipv4Address string   `xml:"Ipv4Address,omitempty"`
	Ipv4Mask    string   `xml:"Ipv4Mask,omitempty"`
	HAddress    string   `xml:"HAddress,omitempty"`
	HType       int      `xml:"HType,omitempty"`
}

// DHCPServerIPInUse table contains binding information about assigned IP addresses
type DHCPServerIPInUse struct {
	XMLName xml.Name  `xml:"DHCPServerIpInUse"`
	IPInUse []IPInUse `xml:"IpInUse"`
}

type IPInUse struct {
	XMLName     xml.Name `xml:"IpInUse"`
	PoolIndex   int      `xml:"PoolIndex"`
	VLANID      int      `xml:"VLANID"`
	HType       int      `xml:"HType"`
	Type        int      `xml:"Type"`
	IfIndex     int      `xml:"IfIndex"`
	Ipv4Address string   `xml:"Ipv4Address"`
	CID         string   `xml:"CID"`
	HAddress    string   `xml:"HAddress"`
	EndLease    string   `xml:"EndLease"`
}
