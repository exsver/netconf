package comware

import "encoding/xml"

type ResourceMonitor struct {
	/* top level
	   ResourceMonitor
		 Configuration
	     Monitors				***ReadOnly***
		   []Monitor			***ReadOnly***
		 Resources
		   []Resource
	*/
	Configuration *ResourceMonitorConfiguration `xml:"Configuration"`
	Monitors      *Monitors                     `xml:"Monitors"`
	Resources     *Resources                    `xml:"Resources"`
}

type ResourceMonitorConfiguration struct {
	XMLName                xml.Name `xml:"Configuration"`
	MinorResendEnable      bool     `xml:"MinorResendEnable,omitempty"`
	SyslogOutput           bool     `xml:"SyslogOutput,omitempty"`           //Output enable for syslog.
	SnmpNotificationOutput bool     `xml:"SnmpNotificationOutput,omitempty"` //Output enable for SNMP notification.
	NetconfEventOutput     bool     `xml:"NetconfEventOutput,omitempty"`     //Output enable for NETCONF event.
}

// Monitors table contains usage information about the resource monitor.
type Monitors struct {
	XMLName  xml.Name  `xml:"Monitors"`
	Monitors []Monitor `xml:"Monitor"`
}

type Monitor struct {
	XMLName    xml.Name    `xml:"Monitor"`
	DeviceNode *DeviceNode `xml:"DeviceNode"`
	Name       string      `xml:"Name"`
	Unit       string      `xml:"Unit"` // absolute | percentage
	Used       int         `xml:"Used"`
	Free       int         `xml:"Free"`
	Total      int         `xml:"Total"`
}

type DeviceNode struct {
	XMLName xml.Name `xml:"DeviceNode"`
	Chassis int      `xml:"Chassis"`
	Slot    int      `xml:"Slot"`
	CPUID   int      `xml:"CPUID"`
}

// Resources table contains thresholds information about the resource monitor.
type Resources struct {
	XMLName   xml.Name   `xml:"Resources"`
	Resources []Resource `xml:"Resource"`
}

type Resource struct {
	XMLName         xml.Name    `xml:"Resource"`
	DeviceNode      *DeviceNode `xml:"DeviceNode"`
	Name            string      `xml:"Name,omitempty"`
	Unit            string      `xml:"Unit,omitempty"`            //Unit of resource monitor threshold. Valid values are: absolute, percentage
	MinorThreshold  int         `xml:"MinorThreshold,omitempty"`  //The free resource threshold for minor condition.
	SevereThreshold int         `xml:"SevereThreshold,omitempty"` //The free resource threshold for severe condition.
}
