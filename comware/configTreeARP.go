package comware

import "encoding/xml"

type ARP struct {
	/* top level
		ARP
	      ArpAttackProtect
	      ArpAuthorized
	      ArpDetection
	        []Detect
	      ArpDetectionTrust
	        []Interface
	      ArpFilterSource
	        []FilterSource
	      ArpGratuitous
	      ArpGratuitousInterval
	        []GratuitousInterval
	      ArpLearnLimit
	        []LearnLimit
	      ArpProxy
	        []Proxy
	      ArpRateLimit
	        []RateLimit
	      ArpRateLimitLog
	      ArpSnooping
	        []Snooping
	      ArpSpecification
	      ArpTable
	        []ArpEntry
	*/
	ArpDetection      *ArpDetection      `xml:"ArpDetection"`
	ArpDetectionTrust *ArpDetectionTrust `xml:"ArpDetectionTrust"`
	ArpFilterSource   *ArpFilterSource   `xml:"ArpFilterSource"`
	ArpGratuitous     *ArpGratuitous     `xml:"ArpGratuitous"`
	ArpRateLimit      *ArpRateLimit      `xml:"ArpRateLimit"`
	ArpRateLimitLog   *ArpRateLimitLog   `xml:"ArpRateLimitLog"`
	ArpSpecification  *ArpSpecification  `xml:"ArpSpecification"`
	ArpTable          *ArpTable          `xml:"ArpTable"`
}

// ArpDetection table contains the information about arp detections.
type ArpDetection struct {
	XMLName   xml.Name `xml:"ArpDetection"`
	ArpDetect []Detect `xml:"Detect"`
}

type Detect struct {
	XMLName                    xml.Name `xml:"Detect"`
	VLANID                     int      `xml:"VLANID"`
	DetectionEnable            bool     `xml:"DetectionEnable"`
	RestrictedForwardingEnable bool     `xml:"RestrictedForwardingEnable"`
}

// ArpDetectionTrust table contains the information of interface that arp trusted.
type ArpDetectionTrust struct {
	XMLName    xml.Name                     `xml:"ArpDetectionTrust"`
	Interfaces []ArpDetectionTrustInterface `xml:"Interface"`
}

type ArpDetectionTrustInterface struct {
	XMLName     xml.Name `xml:"Interface"`
	IfIndex     int      `xml:"IfIndex"`
	TrustEnable bool     `xml:"TrustEnable"`
}

// ArpFilterSource table contains the information about arp gateway protection and protected ip address.
type ArpFilterSource struct {
	XMLName       xml.Name       `xml:"ArpFilterSource"`
	FilterSources []FilterSource `xml:"FilterSource"`
}

// ArpGratuitous contains gratuitous ARP packets enable information.
type ArpGratuitous struct {
	XMLName     xml.Name `xml:"ArpGratuitous"`
	LearnEnable bool     `xml:"LearnEnable"`
	SendEnable  bool     `xml:"SendEnable"`
}

type FilterSource struct {
	XMLName     xml.Name `xml:"FilterSource"`
	IfIndex     int      `xml:"IfIndex"`
	Ipv4Address string   `xml:"Ipv4Address"`
}

// ArpSpecification table contains ARP specification information.
type ArpSpecification struct {
	XMLName              xml.Name `xml:"ArpSpecification"`
	SupportMultiport     bool     `xml:"SupportMultiport"`
	SupportDetection     bool     `xml:"SupportDetection"`
	SupportFilter        bool     `xml:"SupportFilter"`
	SupportSnooping      bool     `xml:"SupportSnooping"`
	SupportModeUNI       bool     `xml:"SupportModeUNI"`
	RateLimitDefault     int      `xml:"RateLimitDefault"`
	RateLimitLowerLimit  int      `xml:"RateLimitLowerLimit"`
	RateLimitUpperLimit  int      `xml:"RateLimitUpperLimit"`
	SMACDefaultThreshold int      `xml:"SMACDefaultThreshold"`
	SMACMinThreshold     int      `xml:"SMACMinThreshold"`
	SMACMaxThreshold     int      `xml:"SMACMaxThreshold"`
	IntervalDefault      int      `xml:"IntervalDefault"`
	LearnMaxNumLimit     int      `xml:"LearnMaxNumLimit"`
	LearnMaxNumDefault   int      `xml:"LearnMaxNumDefault"`
}

// ArpRateLimit table contains ARP rate limit information.
type ArpRateLimit struct {
	XMLName             xml.Name                `xml:"ArpRateLimit"`
	RateLimitInterfaces []ArpRateLimitInterface `xml:"RateLimit"`
}

type ArpRateLimitInterface struct {
	XMLName         xml.Name `xml:"RateLimit"`
	IfIndex         int      `xml:"IfIndex"`
	RateLimitEnable bool     `xml:"RateLimitEnable"`
	RateLimitNum    int      `xml:"RateLimitNum"`
}

// ArpRateLimitLog contains ARP rate limit log information.
type ArpRateLimitLog struct {
	XMLName     xml.Name `xml:"ArpRateLimitLog"`
	LogEnable   bool     `xml:"LogEnable"`
	LogInterval int      `xml:"LogInterval,omitempty"`
}

// ArpTable table contains ARP table information.
type ArpTable struct {
	XMLName    xml.Name   `xml:"ArpTable"`
	ArpEntries []ArpEntry `xml:"ArpEntry"`
}

type ArpEntry struct {
	XMLName     xml.Name `xml:"ArpEntry"`
	IfIndex     int      `xml:"IfIndex"`
	VLANID      int      `xml:"VLANID"`
	Ipv4Address string   `xml:"Ipv4Address"`
	MacAddress  string   `xml:"MacAddress"`
	PortIndex   int      `xml:"PortIndex"`
	VrfIndex    int      `xml:"VrfIndex"`
	ArpType     int      `xml:"ArpType"`
}
