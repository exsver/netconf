package junos

import "encoding/xml"

type Chassis struct {
	XMLName           xml.Name           `xml:"chassis"`
	AggregatedDevices *AggregatedDevices `xml:"aggregated-devices,omitempty"`
	Alarm             *Alarm             `xml:"alarm,omitempty"`
	NetworkServices   string             `xml:"network-services,omitempty"`
}

type AggregatedDevices struct {
	XMLName  xml.Name  `xml:"aggregated-devices"`
	Ethernet *Ethernet `xml:"ethernet"`
}

type Ethernet struct {
	XMLName     xml.Name `xml:"ethernet"`
	DeviceCount int      `xml:"device-count"`
}

type Alarm struct {
	XMLName            xml.Name                 `xml:"alarm"`
	ManagementEthernet *AlarmManagementEthernet `xml:"management-ethernet"`
	Ethernet           *AlarmEthernet           `xml:"ethernet"`
}

type AlarmManagementEthernet struct {
	XMLName  xml.Name `xml:"management-ethernet"`
	LinkDown string   `xml:"link-down"` // ignore | red | yellow
}

type AlarmEthernet struct {
	XMLName  xml.Name `xml:"ethernet"`
	LinkDown string   `xml:"link-down"` // ignore | red | yellow
}
