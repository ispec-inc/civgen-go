package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/ispec-inc/civgen-go/model/value"
)

func (g generator) Dao(path value.Filepath) error {
	f := jen.NewFile(value.PackageDao.Name())

	f.ImportName(value.PackageEntity.Path(), value.PackageEntity.Name())
	f.ImportName(value.PackageModel.Path(), value.PackageModel.Name())
	f.ImportName(value.PackageError.Path(), value.PackageError.Name())
	f.ImportName(value.PackageGorm.Path(), "gorm")

	f.Type().Id(g.name).Struct(
		jen.Id("db").Add(jen.Op("*")).Qual(value.PackageGorm.Path(), "DB"),
	)

	f.Func().Id(fmt.Sprintf("New%s", g.name)).Params(
		jen.Id("db").Add(jen.Op("*")).Qual(value.PackageGorm.Path(), "DB"),
	).Id(g.name).Block(
		jen.Return().Id(g.name).Block(jen.Id("db").Id(",")),
	)

	f.Line()

	f.Func().Params(
		jen.Id("d").Id(g.name),
	).Id("Get").Params(
		jen.Id("id").Int64(),
	).Params(
		jen.Qual(value.PackageModel.Path(), g.name),
		jen.Qual(value.PackageError.Path(), "Error"),
	).Block(
		jen.Var().Id("e").Qual(value.PackageEntity.Path(), g.name),
		jen.If(jen.Id("err := d.db.First(&e, id).Error; err != nil")).Block(
			jen.Return(
				jen.Qual(value.PackageModel.Path(), g.name).Block(),
				jen.Id("newGormFind").Params(
					jen.Id("err"),
					jen.Qual(value.PackageEntity.Path(), fmt.Sprintf("%sModelName", g.name)),
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
		jen.Id("d").Id(g.name),
	).Id("List").Params(
		jen.Id("ids").Index().Int64(),
	).Params(
		jen.Id("ms").Index().Qual(value.PackageModel.Path(), g.name),
		jen.Id("aerr").Qual(value.PackageError.Path(), "Error"),
	).Block(
		jen.Id("query").Op(":=").Id("d.db"),
		jen.If(jen.Len(jen.Id("ids")).Op(">").Lit(0)).Block(
			jen.Id("query").Op("=").Id("query.Where(\"id in (?)\", ids)"),
		),
		jen.Line(),
		jen.Var().Id("es").Index().Qual(value.PackageEntity.Path(), g.name),
		jen.If(jen.Id("err := query.Find(&es).Error; err != nil")).Block(
			jen.Return(
				jen.Index().Qual(value.PackageModel.Path(), g.name).Block(),
				jen.Id("newGormFind").Params(
					jen.Id("err"),
					jen.Qual(value.PackageEntity.Path(), fmt.Sprintf("%sModelName", g.name)),
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

	return f.Save(path.String())
}

func (g generator) DaoTest(path value.Filepath) error {
	f := jen.NewFile(value.PackageDao.Name())

	f.ImportName(value.PackageTesting.Path(), "testing")
	// Import assert by Id() method since the test code is written by plain text
	f.Id(fmt.Sprintf("import \"%s\"", value.PackageAssert.Path()))

	f.Func().Id(fmt.Sprintf("Test%sDao_Get", g.name)).Params(
		jen.Id("t").Op("*").Qual(value.PackageTesting.Path(), "T"),
	).Block(
		jen.Id("t.Helper()"),
		jen.Id(fmt.Sprintf("d := New%s(db)", g.name)),
		jen.Id(g.getTableTest()),
	)

	f.Line()

	f.Func().Id(fmt.Sprintf("Test%sDao_List", g.name)).Params(
		jen.Id("t").Op("*").Qual(value.PackageTesting.Path(), "T"),
	).Block(
		jen.Id("t.Helper()"),
		jen.Id(fmt.Sprintf("d := New%s(db)", g.name)),
		jen.Id(g.listTableTest()),
	)

	return f.Save(path.String())
}

func (g generator) getTableTest() string {
	return fmt.Sprintf(getTableTestCode, g.name)
}

func (g generator) listTableTest() string {
	return fmt.Sprintf(listTableTestCode, g.name)
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
