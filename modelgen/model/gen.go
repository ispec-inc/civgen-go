package model

import (
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

type GenerateModelFileInput struct {
	Name   string
	Path   string
	Fields string
	Layer  Layer
}

func GenerateModelFile(ipt GenerateModelFileInput) error {
	f := jen.NewFile(strcase.ToSnake(ipt.Name))

	f.ImportName("time", "time")

	fs := NewFields(ipt.Fields)

	var jenFields []jen.Code
	for _, f := range fs {
		jenFields = append(jenFields, f.ToJenStatement(ipt.Layer))
	}

	f.Type().Id(ipt.Name).Struct(jenFields...)

	return f.Save(ipt.Path)
}
