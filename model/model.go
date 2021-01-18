package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ispec-inc/civgen-go/model/generator"
	"github.com/ispec-inc/civgen-go/model/value"
)

var (
	// Required
	name   = flag.String("name", "", "Model name")
	fields = flag.String("fields", "", "Fields of the model (e.g. ID:string,Name:string,CreatedAt:time.Time,Update:time.Time")
	// Optional
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
	loadConfig()
	validateFlag()

	gen := generator.NewGenerator(*name, *fields)

	if *createEntity {
		path := value.NewFilepath(cfg.EntityPath, *name, "")
		if err := gen.Model(path, value.LayerEntity); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "entity", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "entity", path.String())
	}
	if *createModel {
		path := value.NewFilepath(cfg.ModelPath, *name, "")
		if err := gen.Model(path, value.LayerModel); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "model", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "model", path.String())
	}
	if *createView {
		path := value.NewFilepath(cfg.ViewPath, *name, "")
		if err := gen.Model(path, value.LayerView); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "view", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "view", path.String())
	}
	if *createRepository {
		path := value.NewFilepath(cfg.RepositoryPath, *name, "")
		if err := gen.Repository(path); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "repository", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "repository", path.String())
	}
	if *createDao {
		path := value.NewFilepath(cfg.DaoPath, *name, "")
		if err := gen.Dao(path); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "dao", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "dao", path.String())
	}
	if *createDaoTest {
		path := value.NewFilepath(cfg.DaoPath, *name, "_test")
		if err := gen.DaoTest(path); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "dao_test", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "dao_test", path.String())
	}

	daoTestMainPath := value.NewFilepath(cfg.DaoPath, "dao_test", "")
	if _, err := os.Stat(daoTestMainPath.String()); os.IsNotExist(err) {
		if err := gen.DaoTestMain(daoTestMainPath); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "dao_test_main", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "dao_test_main", daoTestMainPath.String())
	}

	daoErrorPath := value.NewFilepath(cfg.DaoPath, "error", "")
	if _, err := os.Stat(daoErrorPath.String()); os.IsNotExist(err) {
		if err := gen.DaoError(daoErrorPath); err != nil {
			fmt.Printf("Failed to generate %s file: %v\n", "dao_error", err)
			return
		}
		fmt.Printf("Generate %s file successfully to '%s'\n", "dao_error", daoErrorPath.String())
	}
}

func validateFlag() {
	if *name == "" {
		log.Fatal(errors.New("'name' cannot be empty."))
	}
	if *fields == "" {
		log.Fatal(errors.New("'fields' cannot be empty."))
	}
}

func usage() {
	_, _ = io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

const usageText = `model should be executed on the root directory of your go project.
Example:
	go run github.com/ispec-inc/civgen-go/model --name User --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time --project_path github.com/ispec-inc/civgen-go/example [other options]
`
