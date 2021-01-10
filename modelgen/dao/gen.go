package dao

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/ispec-inc/civgen-go/modelgen/pkg"
)

const (
	gormPkg    = "gorm.io/gorm"
	assertPkg  = "github.com/stretchr/testify/assert"
	testingPkg = "testing"
)

type GenerateFileInput struct {
	Name string
	Path string
}

type GenerateTestFileInput struct {
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

func GenerateTestFile(ipt GenerateTestFileInput) error {
	f := jen.NewFile(pkg.Pkgs.Dao.Name())

	f.ImportName(testingPkg, "testing")
	// Import assert by Id() method since the test code is written by plain text
	f.Id(fmt.Sprintf("import \"%s\"", assertPkg))

	f.Func().Id(fmt.Sprintf("Test%sDao_Get", ipt.Name)).Params(
		jen.Id("t").Op("*").Qual(testingPkg, "T"),
	).Block(
		jen.Id("t.Helper()"),
		jen.Id(fmt.Sprintf("d := New%s(db)", ipt.Name)),
		jen.Id(getTableTest(ipt.Name)),
	)

	f.Line()

	f.Func().Id(fmt.Sprintf("Test%sDao_List", ipt.Name)).Params(
		jen.Id("t").Op("*").Qual(testingPkg, "T"),
	).Block(
		jen.Id("t.Helper()"),
		jen.Id(fmt.Sprintf("d := New%s(db)", ipt.Name)),
		jen.Id(listTableTest(ipt.Name)),
	)

	return f.Save(ipt.Path)
}

func getTableTest(name string) string {
	return fmt.Sprintf(getTableTestCode, name)
}

func listTableTest(name string) string {
	return fmt.Sprintf(listTableTestCode, name)
}

const getTableTestCode = `
	cases := []struct {
		name string
		id   int64
		want int64
		err  bool
	}{
		{
			name: "Found",
			id:   int64(1),
			want: int64(1),
			err:  false,
		},
		{
			name: "NotFound",
			id:   int64(2),
			want: int64(0),
			err:  true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareTestData("./testdata/%s/get.sql"); err != nil {
				t.Error(err)
			}

			opt, aerr := d.Get(tc.id)

			assert.Exactly(t, tc.want, opt.ID)
			if tc.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
			}
		})
	}
`

const listTableTestCode = `
	cases := []struct {
		name string
		ids  []int64
		want int
		err  bool
	}{
		{
			name: "ByIDs",
			ids:  []int64{1},
			want: 1,
			err:  false,
		},
		{
			name: "All",
			ids:  nil,
			want: 1,
			err:  false,
		},
		{
			name: "NotFound",
			ids:  []int64{2},
			want: 0,
			err:  false,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareTestData("./testdata/%s/list.sql"); err != nil {
				t.Error(err)
			}

			opt, aerr := d.List(tc.ids)

			assert.Exactly(t, tc.want, len(opt))
			if tc.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
			}
		})
	}
`
