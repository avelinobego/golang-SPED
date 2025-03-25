package builder

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"sped/utils"
	"strings"
)

type values struct {
	FieldName   string
	Position    int
	Value       any
	IsAttribute bool
}

type mapValues map[string]values

type Builder struct {
	instance  *reflect.Value
	name      string
	namespace string
	values    mapValues
	errors    []error
}

func NewBuilder(name string) *Builder {
	return &Builder{
		name:     name,
		values:   make(mapValues, 0),
		instance: nil,
	}
}

func (b *Builder) SetNameSpace(ns string) {
	b.namespace = ns
}

func (b *Builder) AddChild(builder *Builder) *Builder {
	temp := builder.GetStruct()

	s := reflect.TypeOf(temp).Elem()
	if field, ok := s.FieldByName("XMLName"); ok {
		n := field.Tag.Get("xml")
		return b.Add(n, temp)
	} else {
		return b
	}

}

func (b *Builder) Add(n string, v any) *Builder {

	if _, found := b.values[n]; found {
		return b
	}

	var temp any

	switch typeDiscovered := v.(type) {
	case Builder:
		temp = typeDiscovered.GetStruct()
		if !b.validateName(temp, n) {
			b.GetStruct()
			return b
		}
	case *Builder:
		temp = typeDiscovered.GetStruct()
		if !b.validateName(temp, n) {
			b.GetStruct()
			return b
		}
	default:
		temp = typeDiscovered
	}

	b.values[n] = values{
		FieldName:   n,
		Position:    len(b.values),
		Value:       temp,
		IsAttribute: false,
	}

	return b
}

func (b *Builder) validateName(r any, n string) bool {
	s := reflect.TypeOf(r)
	if field, ok := s.FieldByName("XMLName"); ok {
		if n != field.Tag.Get("xml") {
			b.errors = append(b.errors, fmt.Errorf(`nomes diferentes. esperava "%s" e recebi "%s"`, field.Tag.Get("xml"), n))
			return false
		}
	}
	return true
}

func (b Builder) GetErros() (result []error) {
	result = append(result, b.errors...)
	return
}

func (b *Builder) AddAttribute(n string, v any) *Builder {

	if _, found := b.values[n]; found {
		return b
	}

	b.values[n] = values{
		FieldName:   n,
		Position:    len(b.values),
		Value:       v,
		IsAttribute: true,
	}

	return b
}

func (b *Builder) BuildStruct() {

	var fields []reflect.StructField

	fields = append(fields, reflect.StructField{
		Name: "XMLName",
		Type: reflect.TypeOf(xml.Name{}),
		Tag:  reflect.StructTag(fmt.Sprintf(`xml:"%s"`, b.name)),
	},
	)

	if b.namespace != "" {
		fields = append(fields, reflect.StructField{
			Name: "XMLnamespace",
			Type: reflect.TypeOf(""),
			Tag:  reflect.StructTag(`xml:"xmlms,attr"`),
		},
		)
	}

	// Criar campos
	tagName := strings.Builder{}

	values := sorteAndTransforMap(b.values)

	for _, v := range values {

		tagName.WriteString(`xml:"%s`)
		if v.IsAttribute {
			tagName.WriteString(`,attr`)
		}
		tagName.WriteString(`"`)

		fields = append(fields, reflect.StructField{
			Name: utils.Capitalize(v.FieldName),
			Type: reflect.TypeOf(v.Value),
			Tag:  reflect.StructTag(fmt.Sprintf(tagName.String(), v.FieldName)),
		})

		tagName.Reset()
	}

	// Criar um novo tipo de struct
	structType := reflect.StructOf(fields)

	name := xml.Name{
		Local: b.name,
	}

	// Criar uma instância da struct
	instance := reflect.New(structType).Elem()
	instance.FieldByName("XMLName").Set(reflect.ValueOf(name))
	if instance.FieldByName("XMLnamespace").IsValid() {
		instance.FieldByName("XMLnamespace").SetString(b.namespace)
	}

	b.instance = &instance

	b.fillValues()
}

func sorteAndTransforMap(mValues mapValues) (result []values) {
	result = make([]values, len(mValues))
	for _, v := range mValues {
		result[v.Position] = v
	}
	return
}

func (b *Builder) fillValues() {
	for _, v := range b.values {
		b.instance.FieldByName(utils.Capitalize(v.FieldName)).Set(reflect.ValueOf(v.Value))
	}

}

func (b *Builder) GetStruct() any {
	if b.instance == nil {
		b.BuildStruct()
	}
	result := b.instance.Interface()
	return &result
}

func (b Builder) String() string {
	return "oi"
}
