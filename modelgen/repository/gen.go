package repository

import (
	"github.com/dave/jennifer/jen"
	"github.com/ispec-inc/civgen-go/modelgen/pkg"
)

type GenerateFileInput struct {
	Name string
	Path string
}

func GenerateFile(ipt GenerateFileInput) error {
	f := jen.NewFile(pkg.Pkgs.Repository.Name())

	f.ImportName("time", "time")
	f.ImportName(pkg.Pkgs.Model.Path(), pkg.Pkgs.Model.Name())
	f.ImportName(pkg.Pkgs.Error.Path(), pkg.Pkgs.Error.Name())

	f.Type().Id(ipt.Name).Interface(
		jen.Id("Get").Params(
			jen.Id("id").Int64(),
		).Params(
			jen.Qual(pkg.Pkgs.Model.Path(), ipt.Name),
			jen.Qual(pkg.Pkgs.Error.Path(), "Error"),
		),
		jen.Id("List").Params(
			jen.Id("ids").Index().Int64(),
		).Params(
			jen.Index().Qual(pkg.Pkgs.Model.Path(), ipt.Name),
			jen.Qual(pkg.Pkgs.Error.Path(), "Error"),
		),
	)

	return f.Save(ipt.Path)
}
