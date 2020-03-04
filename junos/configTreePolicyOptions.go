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
	XMLName                xml.Name         `xml:"prefix-list"`
	NetconfConfigOperation string           `xml:"operation,attr,omitempty"`
	Name                   string           `xml:"name"`
	PrefixListItems        []PrefixListItem `xml:"prefix-list-item,omitempty"`
}

type PrefixListItem struct {
	XMLName xml.Name `xml:"prefix-list-item"`
	Name    string   `xml:"name"` // Address prefix
}

type PolicyStatement struct {
	XMLName                xml.Name `xml:"policy-statement"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
}

type Community struct {
	XMLName                xml.Name `xml:"community"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
}

type ASPathGroup struct {
	XMLName                xml.Name `xml:"as-path-group"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	ASPaths                []ASPath `xml:"as-path"`
}

type ASPath struct {
	XMLName                xml.Name `xml:"as-path"`
	NetconfConfigOperation string   `xml:"operation,attr,omitempty"`
	Name                   string   `xml:"name"`
	Path                   string   `xml:"path"`
}

func (prefixList *PrefixList) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			PolicyOptions: &PolicyOptions{
				PrefixLists: []PrefixList{*prefixList},
			},
		},
	}
}

func (policyStatement *PolicyStatement) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			PolicyOptions: &PolicyOptions{
				PolicyStatements: []PolicyStatement{*policyStatement},
			},
		},
	}
}

func (community *Community) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			PolicyOptions: &PolicyOptions{
				Communities: []Community{*community},
			},
		},
	}
}

func (asPathGroup *ASPathGroup) ConvertToConfig() *Config {
	return &Config{
		Configuration: &Configuration{
			PolicyOptions: &PolicyOptions{
				ASPathGroups: []ASPathGroup{*asPathGroup},
			},
		},
	}
}

func (asPath *ASPath) ConvertToConfig(asPathGroup string) *Config {
	if asPathGroup == "" {
		return &Config{
			Configuration: &Configuration{
				PolicyOptions: &PolicyOptions{
					ASPaths: []ASPath{*asPath},
				},
			},
		}
	}

	return &Config{
		Configuration: &Configuration{
			PolicyOptions: &PolicyOptions{
				ASPathGroups: []ASPathGroup{
					{
						Name:    asPathGroup,
						ASPaths: []ASPath{*asPath},
					},
				},
			},
		},
	}
}
