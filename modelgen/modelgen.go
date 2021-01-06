package main

import (
	"flag"
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/ispec-inc/civgen-go/modelgen/model"
	"github.com/ispec-inc/civgen-go/modelgen/pkg"
)

var (
	// Required
	name        = flag.String("name", "", "Model name")
	fields      = flag.String("fields", "", "Fields of the model (e.g. ID:string,Name:string,CreatedAt:time.Time,Update:time.Time")
	projectPath = flag.String("project_path", "", "Go package path of this project root")

	// Optional
	entityPath = flag.String("entity_path", "pkg/infra/entity", "Relative path to the entity directory from 'project_root'")
	modelPath  = flag.String("model_path", "pkg/domain/model", "Relative path to the model directory from 'project_root'")
	viewPath   = flag.String("view_path", "pkg/view", "Relative path to the view directory from 'project_root'")

	createEntity = flag.Bool("create_entity", true, "Create entity file, if true")
	createModel  = flag.Bool("create_model", true, "Create model file, if true")
	createView   = flag.Bool("create_view", true, "Create view file, if true")
)

func main() {
	flag.Parse()

	if err := setPackages(); err != nil {
		fmt.Println(err.Error())
		return
	}

	generateModelFile(model.LayerEntity)
	generateModelFile(model.LayerModel)
	generateModelFile(model.LayerView)
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

	err := model.GenerateModelFile(
		model.GenerateModelFileInput{
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

func setPackages() error {
	ipt := pkg.SetPkgsInput{
		ProjectRoot: *projectPath,
		EntityPath:  *entityPath,
		ModelPath:   *modelPath,
		ViewPath:    *viewPath,
	}
	return pkg.SetPkgs(ipt)
}
