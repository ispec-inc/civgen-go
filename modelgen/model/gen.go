package model

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/ispec-inc/civgen-go/modelgen/pkg"
)

type GenerateModelFileInput struct {
	Name   string
	Path   string
	Fields string
	Layer  Layer
}

func GenerateModelFile(ipt GenerateModelFileInput) error {
	f := jen.NewFile(ipt.Layer.String())

	f.ImportName("time", "time")

	switch ipt.Layer {
	case LayerEntity:
		f.ImportName(pkg.Pkgs.Model.Path(), pkg.Pkgs.Model.Name())
	case LayerView:
		f.ImportName(pkg.Pkgs.Model.Path(), pkg.Pkgs.Model.Name())
	}

	fs := NewFields(ipt.Fields)

	var jenFields []jen.Code
	for _, f := range fs {
		jenFields = append(jenFields, f.ToStructField(ipt.Layer))
	}

	f.Type().Id(ipt.Name).Struct(jenFields...)

	switch ipt.Layer {
	case LayerEntity:
		addEntityModelParser(f, ipt.Name, fs)
	case LayerView:
		addViewConstructor(f, ipt.Name, fs)
	}

	return f.Save(ipt.Path)
}

func addEntityModelParser(f *jen.File, name string, fields []Field) {
	parsers := make(jen.Dict)
	for _, f := range fields {
		parsers[jen.Id(f.Name)] = jen.Id(fmt.Sprintf("e.%s", f.Name))
	}

	f.Func().Params(jen.Id("e").Id(name)).Id("ToModel").Params().Qual(pkg.Pkgs.Model.Path(), name).
		Block(jen.Return(jen.Qual(pkg.Pkgs.Model.Path(), name).Block(parsers)))
}

func addViewConstructor(f *jen.File, name string, fields []Field) {
	parsers := make(jen.Dict)
	for _, f := range fields {
		parsers[jen.Id(f.Name)] = jen.Id(fmt.Sprintf("m.%s", f.Name))
	}

	funcName := fmt.Sprintf("New%s", name)

	f.Func().Id(funcName).Params(jen.Id("m").Qual(pkg.Pkgs.Model.Path(), name)).Id(name).
		Block(jen.Return(jen.Id(name).Block(parsers)))
}
