package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const (
	configPath = "./.civgen-model.yaml"
)

type config struct {
	ProjectPath    string `yaml:"project_path"`
	EntityPath     string `yaml:"entity_path"`
	ModelPath      string `yaml:"model_path"`
	ViewPath       string `yaml:"view_path"`
	RepositoryPath string `yaml:"repository_path"`
	DaoPath        string `yaml:"dao_path"`
	ErrorPath      string `yaml:"error_path"`
	DatabasePath   string `yaml:"database_path"`
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

	var cfg config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("cannot parse %s correctly: %v\n", configPath, err)
		os.Exit(1)
	}

	if cfg.ProjectPath != "" {
		*projectPath = cfg.ProjectPath
	}
	if cfg.EntityPath != "" {
		*entityPath = cfg.EntityPath
	}
	if cfg.ModelPath != "" {
		*modelPath = cfg.ModelPath
	}
	if cfg.ViewPath != "" {
		*viewPath = cfg.ViewPath
	}
	if cfg.RepositoryPath != "" {
		*repositoryPath = cfg.RepositoryPath
	}
	if cfg.DaoPath != "" {
		*daoPath = cfg.DaoPath
	}
	if cfg.ErrorPath != "" {
		*errorPath = cfg.ErrorPath
	}
	if cfg.DatabasePath != "" {
		*databasePath = cfg.DatabasePath
	}
}
