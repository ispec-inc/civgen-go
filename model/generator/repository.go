package generator

import (
	"github.com/dave/jennifer/jen"
	"github.com/ispec-inc/civgen-go/model/value"
)

func (g generator) Repository(path value.Filepath) error {
	f := jen.NewFile(value.PackageRepository.Name())

	f.ImportName(value.PackageModel.Path(), value.PackageModel.Name())
	f.ImportName(value.PackageError.Path(), value.PackageError.Name())

	f.Type().Id(g.name).Interface(
		jen.Id("Get").Params(
			jen.Id("id").Int64(),
		).Params(
			jen.Qual(value.PackageModel.Path(), g.name),
			jen.Qual(value.PackageError.Path(), "Error"),
		),
		jen.Id("List").Params(
			jen.Id("ids").Index().Int64(),
		).Params(
			jen.Index().Qual(value.PackageModel.Path(), g.name),
			jen.Qual(value.PackageError.Path(), "Error"),
		),
	)

	return f.Save(path.String())
}
