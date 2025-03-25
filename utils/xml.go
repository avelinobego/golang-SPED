package utils

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

type xmlValue struct {
	RootName string
	Value    any
}

func ToXml(value any) (string, error) {
	temp := xmlValue{
		Value: value,
	}

	result, err := xml.Marshal(temp)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (xmlv xmlValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return decodeStruct(e, xmlv)
}

func decodeStruct(e *xml.Encoder, input xmlValue) error {
	originalValue := reflect.ValueOf(input.Value)
	if originalValue.Kind() == reflect.Pointer {
		originalValue = originalValue.Elem()
	}

	var fields []reflect.StructField
	fields = append(fields, reflect.StructField{
		Name: "XMLName",
		Type: reflect.TypeOf(xml.Name{}),
	},
	)

	originaType := getType(reflect.TypeOf(input.Value))

	if originaType.Kind() != reflect.Struct {
		return e.Encode(input.Value)
	}

	//Definir outros campos
	for i := range originaType.NumField() {

		tipo := getType(originaType.Field(i).Type)

		if tipo.Kind() == reflect.Struct {

			fields = append(fields, reflect.StructField{
				Name: tipo.Name(),
				Type: reflect.TypeOf(xmlValue{}),
				Tag:  reflect.StructTag(fmt.Sprintf(`xml:"%s"`, UnCapitalize(CleanBetween(tipo.Name(), "[", "]")))),
			},
			)
			continue
		}

		fields = append(fields, reflect.StructField{
			Name: originaType.Field(i).Name,
			Type: originaType.Field(i).Type,
			Tag:  reflect.StructTag(fmt.Sprintf(`xml:"%s"`, UnCapitalize(originaType.Field(i).Name))),
		},
		)
	}

	// Criar um novo tipo de struct
	newStructType := reflect.StructOf(fields)
	newStruct := reflect.New(newStructType).Elem()

	for i := range newStruct.NumField() {
		if i == 0 {

			newStruct.Field(i).Set(reflect.ValueOf(xml.Name{
				Local: UnCapitalize(CleanBetween(originaType.Name(), "[", "]")),
			}))
			continue
		}

		originalIndex := i - 1

		tipo := getType(newStruct.Field(i).Type())
		if tipo.Kind() == reflect.Struct {
			xmlValue := xmlValue{
				Value: originalValue.Field(originalIndex).Interface(),
			}
			newStruct.Field(i).Set(reflect.ValueOf(xmlValue))
			continue
		}

		newStruct.Field(i).Set(originalValue.Field(originalIndex))
	}

	return e.Encode(newStruct.Interface())

}

func getType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Pointer {
		return t.Elem()
	}
	return t
}

func removeGeneric(name string) string {
	return strings.TrimLeft(strings.TrimRight(name, "]"), "[]")
}
