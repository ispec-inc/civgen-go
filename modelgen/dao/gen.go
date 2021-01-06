package dao

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/ispec-inc/civgen-go/modelgen/pkg"
)

const (
	gormPkg = "gorm.io/gorm"
)

type GenerateFileInput struct {
	Name string
	Path string
}

func GenerateFile(ipt GenerateFileInput) error {
	f := jen.NewFile(pkg.Pkgs.Dao.Name())

	f.ImportName(pkg.Pkgs.Entity.Path(), pkg.Pkgs.Entity.Name())
	f.ImportName(pkg.Pkgs.Model.Path(), pkg.Pkgs.Model.Name())
	f.ImportName(pkg.Pkgs.Error.Path(), pkg.Pkgs.Error.Name())
	f.ImportName(gormPkg, "gorm")

	f.Type().Id(ipt.Name).Struct(
		jen.Id("db").Add(jen.Op("*")).Qual(gormPkg, "DB"),
	)

	f.Func().Id(fmt.Sprintf("New%s", ipt.Name)).Params(
		jen.Id("db").Add(jen.Op("*")).Qual(gormPkg, "DB"),
	).Id(ipt.Name).Block(
		jen.Return().Id(ipt.Name).Block(jen.Id("db").Id(",")),
	)

	f.Line()

	f.Func().Params(
		jen.Id("d").Id(ipt.Name),
	).Id("Get").Params(
		jen.Id("id").Int64(),
	).Params(
		jen.Qual(pkg.Pkgs.Model.Path(), ipt.Name),
		jen.Qual(pkg.Pkgs.Error.Path(), "Error"),
	).Block(
		jen.Var().Id("e").Qual(pkg.Pkgs.Entity.Path(), ipt.Name),
		jen.If(jen.Id("err := d.db.First(&e, id).Error; err != nil")).Block(
			jen.Return(
				jen.Qual(pkg.Pkgs.Model.Path(), ipt.Name).Block(),
				jen.Qual(pkg.Pkgs.Error.Path(), "NewGormFind").Params(
					jen.Id("err"),
					jen.Qual(pkg.Pkgs.Entity.Path(), fmt.Sprintf("%sModelName", ipt.Name)),
				),
			),
		),
		jen.Return(
			jen.Id("e.ToModel()"),
			jen.Nil(),
		),
	)

	f.Line()

	f.Func().Params(
		jen.Id("d").Id(ipt.Name),
	).Id("List").Params(
		jen.Id("ids").Index().Int64(),
	).Params(
		jen.Id("ms").Index().Qual(pkg.Pkgs.Model.Path(), ipt.Name),
		jen.Id("aerr").Qual(pkg.Pkgs.Error.Path(), "Error"),
	).Block(
		jen.Id("query").Op(":=").Id("d.db"),
		jen.If(jen.Len(jen.Id("ids")).Op(">").Lit(0)).Block(
			jen.Id("query").Op("=").Id("query.Where(\"id in (?)\", ids)"),
		),
		jen.Line(),
		jen.Var().Id("es").Index().Qual(pkg.Pkgs.Entity.Path(), ipt.Name),
		jen.If(jen.Id("err := query.Find(&es).Error; err != nil")).Block(
			jen.Return(
				jen.Index().Qual(pkg.Pkgs.Model.Path(), ipt.Name).Block(),
				jen.Qual(pkg.Pkgs.Error.Path(), "NewGormFind").Params(
					jen.Id("err"),
					jen.Qual(pkg.Pkgs.Entity.Path(), fmt.Sprintf("%sModelName", ipt.Name)),
				),
			),
		),
		jen.Line(),
		jen.For(jen.Id("_, e := range es")).Block(
			jen.Id("ms = append(ms, e.ToModel())"),
		),
		jen.Return(
			jen.Id("ms"),
			jen.Nil(),
		),
	)

	return f.Save(ipt.Path)
}
