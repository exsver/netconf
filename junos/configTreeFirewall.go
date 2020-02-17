package junos

import "encoding/xml"

type Firewall struct {
	XMLName xml.Name `xml:"firewall"`
	Filters []Filter `xml:"filter"`
}

type Filter struct {
	XMLName xml.Name    `xml:"filter"`
	Name    string      `xml:"name"`
	From    *FilterFrom `xml:"from"`
	Then    *FilterThen `xml:"then"`
}

type FilterFrom struct {
	XMLName xml.Name `xml:"from"`
}

type FilterThen struct {
	XMLName xml.Name `xml:"then"`
}
