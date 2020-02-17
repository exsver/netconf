package junos

import "encoding/xml"

type PolicyOptions struct {
	XMLName          xml.Name          `xml:"policy-options"`
	PrefixLists      []PrefixList      `xml:"prefix-list,omitempty"`
	PolicyStatements []PolicyStatement `xml:"policy-statement,omitempty"`
	Communities      []Community       `xml:"community,omitempty"`
	ASPaths          []ASPath          `xml:"as-path,omitempty"`
	ASPathGroups     []ASPathGroup     `xml:"as-path-group,omitempty"`
}

type PrefixList struct {
	XMLName         xml.Name         `xml:"prefix-list"`
	Name            string           `xml:"name"`
	PrefixListItems []PrefixListItem `xml:"prefix-list-item,omitempty"`
}

type PrefixListItem struct {
	XMLName xml.Name `xml:"prefix-list-item"`
	Name    string   `xml:"name"` // Address prefix
}

type PolicyStatement struct {
	XMLName xml.Name `xml:"policy-statement"`
	Name    string   `xml:"name"`
}

type Community struct {
	XMLName xml.Name `xml:"community"`
	Name    string   `xml:"name"`
}

type ASPathGroup struct {
	XMLName xml.Name `xml:"as-path-group"`
	Name    string   `xml:"name"`
	ASPaths []ASPath `xml:"as-path"`
}

type ASPath struct {
	XMLName xml.Name `xml:"as-path"`
	Name    string   `xml:"name"`
	Path    string   `xml:"path"`
}
