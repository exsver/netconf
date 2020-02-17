package junos

import "encoding/xml"

type RoutingInstances struct {
	XMLName   xml.Name          `xml:"routing-instances"`
	Instances []RoutingInstance `xml:"instance,omitempty"`
}

type RoutingInstance struct {
	XMLName      xml.Name `xml:"instance"`
	Name         string   `xml:"name"`
	Description  string   `xml:"description,omitempty"`
	InstanceType string   `xml:"instance-type"`
}
