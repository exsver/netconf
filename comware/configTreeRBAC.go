package comware

import "encoding/xml"

type RBAC struct {
	/* top level
	   RBAC
	     Features
	       []Feature
	     Roles
	       []Role
	     Rules
	       []Rule
	*/
	Features *RBACFeatures `xml:"Features"`
}

type RBACFeatures struct {
	XMLName  xml.Name      `xml:"Features"`
	Features []RBACFeature `xml:"Feature"`
}

type RBACFeature struct {
	XMLName     xml.Name `xml:"Feature"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"`
}
