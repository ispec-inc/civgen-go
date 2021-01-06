package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/ispec-inc/civgen-go/modelgen/dao"
	"github.com/ispec-inc/civgen-go/modelgen/model"
	"github.com/ispec-inc/civgen-go/modelgen/pkg"
	"github.com/ispec-inc/civgen-go/modelgen/repository"
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
)

const usageText = `modelgen should be executed on the root directory of your go project.
Example:
	go run github.com/ispec-inc/civgen-go/modelgen --name User --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time --project_path github.com/ispec-inc/civgen-go/example [other options]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	if err := setPackages(); err != nil {
		fmt.Println(err.Error())
		return
	}

	generateModelFile(model.LayerEntity)
	generateModelFile(model.LayerModel)
	generateModelFile(model.LayerView)
	generateRepositoryFile()
	generateDaoFile()
}

func generateModelFile(layer model.Layer) {
	var doCreate bool
	switch layer {
	case model.LayerEntity:
		doCreate = *createEntity
	case model.LayerModel:
		doCreate = *createModel
	case model.LayerView:
		doCreate = *createView
	}

	var path string
	switch layer {
	case model.LayerEntity:
		path = *entityPath
	case model.LayerModel:
		path = *modelPath
	case model.LayerView:
		path = *viewPath
	}

	filepath := fmt.Sprintf("%s/%s.go", path, strcase.ToSnake(*name))

	if !doCreate {
		return
	}

	err := model.GenerateFile(
		model.GenerateFileInput{
			Name:   *name,
			Path:   filepath,
			Fields: *fields,
			Layer:  layer,
		},
	)
	if err != nil {
		fmt.Printf("Failed to generate %s file: %v\n", layer.String(), err)
		return
	}

	fmt.Printf("Generate %s file successfully to '%s'\n", layer.String(), filepath)
}

func generateRepositoryFile() {
	if !*createRepository {
		return
	}

	filepath := fmt.Sprintf("%s/%s.go", *repositoryPath, strcase.ToSnake(*name))

	err := repository.GenerateFile(
		repository.GenerateFileInput{
			Name: *name,
			Path: filepath,
		},
	)
	if err != nil {
		fmt.Printf("Failed to generate repository file: %v\n", err)
		return
	}

	fmt.Printf("Generate repository file successfully to '%s'\n", filepath)
}

func generateDaoFile() {
	if !*createDao {
		return
	}

	filepath := fmt.Sprintf("%s/%s.go", *daoPath, strcase.ToSnake(*name))

	err := dao.GenerateFile(
		dao.GenerateFileInput{
			Name: *name,
			Path: filepath,
		},
	)
	if err != nil {
		fmt.Printf("Failed to generate dao file: %v\n", err)
		return
	}

	fmt.Printf("Generate dao file successfully to '%s'\n", filepath)
}

func setPackages() error {
	ipt := pkg.SetPkgsInput{
		ProjectRoot:    *projectPath,
		EntityPath:     *entityPath,
		ModelPath:      *modelPath,
		ViewPath:       *viewPath,
		RepositoryPath: *repositoryPath,
		DaoPath:        *daoPath,
		ErrorPath:      *errorPath,
	}
	return pkg.SetPkgs(ipt)
}

func usage() {
	_, _ = io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}
