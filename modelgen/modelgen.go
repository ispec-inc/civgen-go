package main

import (
	"flag"
	"fmt"
	"log"

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

	if *createEntity {
		err := model.GenerateModelFile(
			model.GenerateModelFileInput{
				Name:   *name,
				Path:   filepath(*entityPath, *name),
				Fields: *fields,
				Layer:  model.LayerEntity,
			},
		)
		if err != nil {
			log.Fatalf("failed to generate entity file: %v", err)
			return
		}
		log.Printf("Generate entity file successfully to '%s'", filepath(*entityPath, *name))
	}

	if *createModel {
		err := model.GenerateModelFile(
			model.GenerateModelFileInput{
				Name:   *name,
				Path:   filepath(*modelPath, *name),
				Fields: *fields,
				Layer:  model.LayerModel,
			},
		)
		if err != nil {
			log.Fatalf("failed to generate model file: %v", err)
			return
		}
		log.Printf("Generate model file successfully to '%s'", filepath(*modelPath, *name))
	}

	if *createView {
		err := model.GenerateModelFile(
			model.GenerateModelFileInput{
				Name:   *name,
				Path:   filepath(*viewPath, *name),
				Fields: *fields,
				Layer:  model.LayerView,
			},
		)
		if err != nil {
			log.Fatalf("failed to generate view file: %v", err)
			return
		}
		log.Printf("Generate view file successfully to '%s'", filepath(*viewPath, *name))
	}
}

func filepath(path, name string) string {
	return fmt.Sprintf("%s/%s/%s.go", *baseDir, path, strcase.ToSnake(name))
}
