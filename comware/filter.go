package comware

import (
	"bytes"
	"fmt"
)

func convertToEmptyTag(in []byte) []byte {
	position := bytes.Index(in, []byte(">"))
	if position != -1 {
		return append(in[0:position], []byte("/>")...)
	}

	return in
}

type XMLFilter struct {
	Key      string
	Value    string
	IsRegExp bool
}

func (filter *XMLFilter) ConvertToXML() []byte {
	if filter.IsRegExp {
		return []byte(fmt.Sprintf(`<%s xmlns:re="http://www.hp.com/netconf/base:1.0" re:regExp="%s"/>`, filter.Key, filter.Value))
	}

	return []byte(fmt.Sprintf("<%s>%s</%s>", filter.Key, filter.Value, filter.Key))
}

func convertFiltersToXML(filters []XMLFilter) (xml []byte) {
	for _, v := range filters {
		xml = append(xml, v.ConvertToXML()...)
	}

	return xml
}
