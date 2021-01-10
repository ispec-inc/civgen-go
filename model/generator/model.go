package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/ispec-inc/civgen-go/model/value"
)

func (g generator) Model(path value.Filepath, layer value.Layer) error {
	f := jen.NewFile(layer.String())

	f.ImportName("time", "time")

	switch layer {
	case value.LayerEntity:
		f.ImportName(value.PackageModel.Path(), value.PackageModel.Name())
	case value.LayerView:
		f.ImportName(value.PackageModel.Path(), value.PackageModel.Name())
	}

	if layer == value.LayerEntity {
		f.Const().Id(fmt.Sprintf("%sModelName", g.name)).Op("=").Lit(g.name)
	}

	fs := value.NewFields(g.fields)

	var jenFields []jen.Code
	for _, f := range fs {
		jenFields = append(jenFields, f.ToStructField(layer))
	}

	f.Type().Id(g.name).Struct(jenFields...)

	switch layer {
	case value.LayerEntity:
		g.addEntityModelParser(f, fs)
	case value.LayerView:
		g.addViewConstructor(f, fs)
	}

	return f.Save(path.String())
}

func (g generator) addEntityModelParser(f *jen.File, fields []value.Field) {
	parsers := make(jen.Dict)
	for _, f := range fields {
		parsers[jen.Id(f.Name)] = jen.Id(fmt.Sprintf("e.%s", f.Name))
	}

	f.Func().Params(jen.Id("e").Id(g.name)).Id("ToModel").Params().Qual(value.PackageModel.Path(), g.name).
		Block(jen.Return(jen.Qual(value.PackageModel.Path(), g.name).Block(parsers)))
}

func (g generator) addViewConstructor(f *jen.File, fields []value.Field) {
	parsers := make(jen.Dict)
	for _, f := range fields {
		parsers[jen.Id(f.Name)] = jen.Id(fmt.Sprintf("m.%s", f.Name))
	}

	funcName := fmt.Sprintf("New%s", g.name)

	f.Func().Id(funcName).Params(jen.Id("m").Qual(value.PackageModel.Path(), g.name)).Id(g.name).
		Block(jen.Return(jen.Id(g.name).Block(parsers)))
}
