package model

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

const (
	nameTypeSep = ":"
	fieldSep    = ","
)

type Field struct {
	Name string
	Type string
}

func (f Field) ToJenStatement(t Layer) *jen.Statement {
	s := jen.Id(f.Name)

	switch f.Type {
	case "time.Time":
		s = s.Qual("time", "Time")
	default:
		s = s.Op(f.Type)
	}

	switch t {
	case LayerEntity:
		return s.Tag(gormTag(f.Name))
	case LayerModel:
		return s
	case LayerView:
		return s.Tag(jsonTag(f.Name))
	default:
		return s
	}
}

func NewField(nameAndType string) Field {
	s := strings.Split(nameAndType, nameTypeSep)
	return Field{
		Name: s[0],
		Type: s[1],
	}
}

func NewFields(nameAndTypes string) (fs []Field) {
	nats := strings.Split(nameAndTypes, fieldSep)
	for _, nat := range nats {
		fs = append(fs, NewField(nat))
	}
	return fs
}

func jsonTag(field string) map[string]string {
	return map[string]string{
		"json": strcase.ToSnake(field),
	}
}

func gormTag(field string) map[string]string {
	return map[string]string{
		"gorm": fmt.Sprintf("column:%s", strcase.ToSnake(field)),
	}
}
