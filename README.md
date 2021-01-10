# civgen-go

## modelgen

### Usage

Generate model files
```
go run github.com/ispec-inc/civgen-go/modelgen \
  --name {model name} \
  --fields {fields of model} \
  --project_path {your project path}

  // Optional
	--entity_path {Relative path to the entity directory from 'project_root'} (default: pkg/infra/entity)
	--model_path {"Relative path to the model directory from 'project_root'} (default: pkg/domain/model)
	--view_path {Relative path to the view directory from 'project_root'} (default: pkg/view)
	--repository_path {Relative path to the repository directory from 'project_root'} (default: pkg/domain/repository)
	--dao_path {Relative path to the dao directory from 'project_root'} (default: pkg/infra/dao)
	--error_path {Relative path to the error directory from 'project_root'} (default: pkg/apperror)
	--create_entity {Create entity file, if true} (default: true)
	--create_model {Create model file, if true} (default: true)
	--create_view {Create view file, if true} (default: true)
	--create_repository {Create repository file, if true} (default: true)
	--create_dao {Create dao file, if true} (default: true)
	--create_dao_test {Create dao test file, if true} (default: true)
```

### Example
Command to generate example/
```
go run github.com/ispec-inc/civgen-go/modelgen \
  --name User \
  --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time \
  --project_path github.com/ispec-inc/civgen-go/example
```

See [example/](./example/) for generated files.
