package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator"
	"github.com/ispec-inc/civgen-go/model/value"
	yaml "gopkg.in/yaml.v2"
)

const (
	configPath = "./.civgen-model.yaml"
)

var cfg *config

type config struct {
	ProjectPath    string `yaml:"project_path" validate:"required"`
	EntityPath     string `yaml:"entity_path" validate:"required"`
	ModelPath      string `yaml:"model_path" validate:"required"`
	ViewPath       string `yaml:"view_path" validate:"required"`
	RepositoryPath string `yaml:"repository_path" validate:"required"`
	DaoPath        string `yaml:"dao_path" validate:"required"`
	ErrorPath      string `yaml:"error_path" validate:"required"`
	DatabasePath   string `yaml:"database_path" validate:"required"`
}

func loadConfig() {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("cannot read %s correctly: %v\n", configPath, err)
		os.Exit(1)
	}

	cfg = &config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		fmt.Printf("cannot parse %s correctly: %v\n", configPath, err)
		os.Exit(1)
	}

	validateConfig()
	setPackages()
}

func validateConfig() {
	val := validator.New()
	if err := val.Struct(cfg); err != nil {
		fmt.Printf("invalid config: %v\n", err)
		os.Exit(1)
	}
}

func setPackages() {
	value.PackageEntity = value.NewLocalPackage(cfg.ProjectPath, cfg.EntityPath)
	value.PackageModel = value.NewLocalPackage(cfg.ProjectPath, cfg.ModelPath)
	value.PackageView = value.NewLocalPackage(cfg.ProjectPath, cfg.ViewPath)
	value.PackageRepository = value.NewLocalPackage(cfg.ProjectPath, cfg.RepositoryPath)
	value.PackageDao = value.NewLocalPackage(cfg.ProjectPath, cfg.DaoPath)
	value.PackageError = value.NewLocalPackage(cfg.ProjectPath, cfg.ErrorPath)
	value.PackageDatabase = value.NewLocalPackage(cfg.ProjectPath, cfg.DatabasePath)
}
