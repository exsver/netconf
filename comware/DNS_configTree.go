package comware

import "encoding/xml"

type DNS struct {
	/* top level
	   DNS
	     DNSConfig
	     DNSServer
	       []Server
	     DNSDomain
	       []Domain
	     DNSHost
	       []Host
	     DNSSpecification
	*/
	DNSConfig        *DNSConfig        `xml:"DNSConfig"`
	DNSServer        *DNSServer        `xml:"DNSServer"`
	DNSDomain        *DNSDomain        `xml:"DNSDomain"`
	DNSHost          *DNSHost          `xml:"DNSHost"`
	DNSSpecification *DNSSpecification `xml:"DNSSpecification"`
}

type DNSConfig struct {
	XMLName        xml.Name `xml:"DNSConfig"`
	DNSProxyEnable bool     `xml:"DNSProxyEnable"`
	DNSDSCP        int      `xml:"DNSDSCP"`  // min: 0, max: 63
	DNSDSCP6       int      `xml:"DNSDSCP6"` // min: 0, max: 63
}

type DNSSpecification struct {
	XMLName         xml.Name `xml:"DNSSpecification"`
	SupportDNSProxy bool     `xml:"SupportDNSProxy"`
}

type DNSServer struct {
	XMLName xml.Name           `xml:"DNSServer"`
	Servers []DNSServerAddress `xml:"Server"`
}

type DNSServerAddress struct {
	XMLName     xml.Name `xml:"Server"`
	IpAddress   string   `xml:"IpAddress"`
	VRF         int      `xml:"VRF"`
	Type        int      `xml:"Type"`
	AddressType int      `xml:"AddressType"`
	IfIndex     int      `xml:"IfIndex"`
	Priority    int      `xml:"Priority"`
}

type DNSDomain struct {
	XMLName xml.Name        `xml:"DNSDomain"`
	Domains []DNSDomainName `xml:"Domain"`
}

type DNSDomainName struct {
	XMLName    xml.Name `xml:"Domain"`
	VRF        int      `xml:"VRF"`
	Type       int      `xml:"Type"`
	Priority   int      `xml:"Priority"`
	DomainName string   `xml:"DomainName"`
}

type DNSHost struct {
	XMLName xml.Name      `xml:"DNSHost"`
	Hosts   []DNSHostName `xml:"Host"`
}

type DNSHostName struct {
	XMLName   xml.Name `xml:"Host"`
	VRF       int      `xml:"VRF"`
	Type      int      `xml:"Type"`
	QueryType int      `xml:"QueryType"`
	HostName  string   `xml:"HostName"`
	Result    string   `xml:"Result"`
}
