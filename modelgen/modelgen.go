package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ispec-inc/civgen-go/modelgen/generator"
	"github.com/ispec-inc/civgen-go/modelgen/value"
)

var (
	// Required
	name        = flag.String("name", "", "Model name")
	fields      = flag.String("fields", "", "Fields of the model (e.g. ID:string,Name:string,CreatedAt:time.Time,Update:time.Time")
	projectPath = flag.String("project_path", "", "Go package path of this project root")

	// Optional
	entityPath     = flag.String("entity_path", "pkg/infra/entity", "Relative path to the entity directory from 'project_root'")
	modelPath      = flag.String("model_path", "pkg/domain/model", "Relative path to the model directory from 'project_root'")
	viewPath       = flag.String("view_path", "pkg/view", "Relative path to the view directory from 'project_root'")
	repositoryPath = flag.String("repository_path", "pkg/domain/repository", "Relative path to the repository directory from 'project_root'")
	daoPath        = flag.String("dao_path", "pkg/infra/dao", "Relative path to the dao directory from 'project_root'")
	errorPath      = flag.String("error_path", "pkg/apperror", "Relative path to the error directory from 'project_root'")

	createEntity     = flag.Bool("create_entity", true, "Create entity file, if true")
	createModel      = flag.Bool("create_model", true, "Create model file, if true")
	createView       = flag.Bool("create_view", true, "Create view file, if true")
	createRepository = flag.Bool("create_repository", true, "Create repository file, if true")
	createDao        = flag.Bool("create_dao", true, "Create dao file, if true")
	createDaoTest    = flag.Bool("create_dao_test", true, "Create dao test file, if true")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	validateFlag()

	value.PackageEntity = value.NewLocalPackage(*projectPath, *entityPath)
	value.PackageModel = value.NewLocalPackage(*projectPath, *modelPath)
	value.PackageView = value.NewLocalPackage(*projectPath, *viewPath)
	value.PackageRepository = value.NewLocalPackage(*projectPath, *repositoryPath)
	value.PackageDao = value.NewLocalPackage(*projectPath, *daoPath)
	value.PackageError = value.NewLocalPackage(*projectPath, *errorPath)

	gen := generator.NewGenerator(*name, *fields)

	generate(gen, value.FiletypeEntity)
	generate(gen, value.FiletypeModel)
	generate(gen, value.FiletypeView)
	generate(gen, value.FiletypeRepository)
	generate(gen, value.FiletypeDao)
	generate(gen, value.FiletypeDaoTest)
}

func generate(gen generator.Generator, ftype value.Filetype) {
	var err error
	var path value.Filepath

	switch ftype {
	case value.FiletypeEntity:
		if !*createEntity {
			return
		}
		path = value.NewFilepath(*entityPath, *name, "")
		err = gen.Model(path, value.LayerEntity)

	case value.FiletypeModel:
		if !*createModel {
			return
		}
		path = value.NewFilepath(*modelPath, *name, "")
		err = gen.Model(path, value.LayerModel)

	case value.FiletypeView:
		if !*createView {
			return
		}
		path = value.NewFilepath(*viewPath, *name, "")
		err = gen.Model(path, value.LayerView)

	case value.FiletypeRepository:
		if !*createRepository {
			return
		}
		path = value.NewFilepath(*repositoryPath, *name, "")
		err = gen.Repository(path)

	case value.FiletypeDao:
		if !*createDao {
			return
		}
		path = value.NewFilepath(*daoPath, *name, "")
		err = gen.Dao(path)

	case value.FiletypeDaoTest:
		if !*createDaoTest {
			return
		}
		path = value.NewFilepath(*daoPath, *name, "_test")
		err = gen.DaoTest(path)

	}

	if err != nil {
		fmt.Printf("Failed to generate %s file: %v\n", ftype.String(), err)
		return
	}

	fmt.Printf("Generate %s file successfully to '%s'\n", ftype.String(), path.String())
}

func validateFlag() {
	if *name == "" {
		log.Fatal(errors.New("'name' cannot be empty."))
	}
	if *fields == "" {
		log.Fatal(errors.New("'fields' cannot be empty."))
	}
	if *projectPath == "" {
		log.Fatal(errors.New("'project_path' cannot be empty."))
	}
}

func usage() {
	_, _ = io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

const usageText = `modelgen should be executed on the root directory of your go project.
Example:
	go run github.com/ispec-inc/civgen-go/modelgen --name User --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time --project_path github.com/ispec-inc/civgen-go/example [other options]
`
