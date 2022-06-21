package comware

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

func (b STPBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	reflStruct := reflect.Indirect(reflect.ValueOf(b))
	structType := reflStruct.Type()

	// Collector is a special public field that should exists per struct
	collectorField := reflStruct.FieldByName("Collector")
	var res string
	if collectorField.IsValid() {
		res = fmt.Sprintf("%v", reflStruct.FieldByName("Collector").Interface())
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		name := field.Name
		value := reflStruct.Field(i).Interface()

		// Skip special and unused field
		if collectorField.IsValid() {
			if name == "Collector" || !strings.Contains(res, name) {
				continue
			}
		}

		if value == nil {
			continue
		}

		e.EncodeElement(value, xml.StartElement{Name: xml.Name{Local: name}})
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
