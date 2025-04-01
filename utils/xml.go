package utils

import (
	"encoding/xml"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func DecodeStruct(input any) any {

	if isNil(input) {
		log.Println("null value")
		return input
	}

	sourceValueType := getRawType(reflect.TypeOf(input))
	sourceValueName := removeGeneric(sourceValueType.Name())

	if sourceValueType.Kind() != reflect.Struct {
		return input
	}

	sourceValue := getRawValue(input)
	// Criando a struct destino
	var destFields []reflect.StructField

	sourceFieldValuesMap := make(map[string]any)
	//Definir outros campos
	for i := range sourceValueType.NumField() {

		sourceFieldName := sourceValueType.Field(i).Name
		sourceFieldType := getRawType(sourceValueType.Field(i).Type)
		sourceFieldValue := sourceValue.Field(i).Interface()

		if sourceFieldType.Kind() == reflect.Struct || sourceFieldType.Kind() == reflect.Interface {
			sourceFieldValuesMap[sourceFieldName] = DecodeStruct(sourceFieldValue)
		} else {
			sourceFieldValuesMap[sourceFieldName] = sourceFieldValue
		}

		fieldValueType := reflect.TypeOf(sourceFieldValuesMap[sourceFieldName])

		if strings.HasSuffix(sourceFieldName, "_") {
			if sourceFieldName == "Ns_" {
				sourceFieldValuesMap["_nameSpace_"] = sourceValue.FieldByName(sourceFieldName).String()
			} else {
				attrName := sourceFieldName[:len(sourceFieldName)-1]
				destFields = append(destFields, reflect.StructField{
					Name: sourceFieldName,
					Type: fieldValueType,
					Tag: reflect.StructTag(fmt.Sprintf(`xml:"%s,attr"`,
						UnCapitalize(attrName)),
					),
				},
				)
			}
			continue
		}

		destFields = append(destFields, reflect.StructField{
			Name: sourceFieldName,
			Type: fieldValueType,
		},
		)
	}

	// Criar um novo tipo de struct
	return newStruct(destFields, sourceValueName, sourceFieldValuesMap)

}

func newStruct(
	destFields []reflect.StructField,
	sourceValueName string,
	sourceFieldValuesMap map[string]any,
) any {

	_nameSpace, ok := sourceFieldValuesMap["_nameSpace_"]
	tagString := `xml:"%s"`
	if ok {
		tagString = fmt.Sprintf(`xml:"%s %s"`, _nameSpace, UnCapitalize(sourceValueName))
		delete(sourceFieldValuesMap, "_nameSpace_")
	} else {
		tagString = fmt.Sprintf(`xml:"%s"`, UnCapitalize(sourceValueName))

	}

	destFields = append(destFields, reflect.StructField{
		Name: "XMLName",
		Type: reflect.TypeFor[xml.Name](),
		Tag:  reflect.StructTag(tagString),
	})

	destStructType := reflect.StructOf(destFields)
	destStruct := reflect.New(destStructType).Elem()

	for _, f := range destFields[:len(destFields)-1] {
		destStruct.FieldByName(f.Name).Set(reflect.ValueOf(sourceFieldValuesMap[f.Name]))
		delete(sourceFieldValuesMap, f.Name)
	}

	return destStruct.Interface()
}

func isNil(input any) bool {
	result := reflect.ValueOf(input)
	if result.Kind() == reflect.Pointer {
		return result.IsNil()
	}
	return false
}

func getRawType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Pointer {
		return t.Elem()
	}
	return t
}

func getRawValue(input any) reflect.Value {
	result := reflect.ValueOf(input)
	if result.Kind() == reflect.Pointer {
		result = result.Elem()
	}
	return result
}

func removeGeneric(name string) string {
	if i := strings.Index(name, "["); i > -1 {
		return name[:i]
	}
	return name
}
