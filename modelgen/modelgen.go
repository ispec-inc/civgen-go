package main

import (
	"flag"
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/ispec-inc/civgen-go/modelgen/model"
)

var (
	// Required
	name   = flag.String("name", "", "Model name")
	fields = flag.String("fields", "", "Fields of the model (e.g. ID:string,Name:string,CreatedAt:time.Time,Update:time.Time")

	// Optional
	baseDir    = flag.String("base_dir", ".", "Path to base directory")
	entityPath = flag.String("entity_path", "pkg/infra/entity", "Path to entity directory")
	modelPath  = flag.String("model_path", "pkg/domain/model", "Path to model directory")
	viewPath   = flag.String("view_path", "pkg/view", "Path to view directory")

	createEntity = flag.Bool("create_entity", true, "Create entity file, if true")
	createModel  = flag.Bool("create_model", true, "Create model file, if true")
	createView   = flag.Bool("create_view", true, "Create view file, if true")
)

func main() {
	flag.Parse()
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

	filepath := fmt.Sprintf("%s/%s/%s.go", *baseDir, path, strcase.ToSnake(*name))

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
