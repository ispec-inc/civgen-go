# civgen-go

## model

### Usage

Generate model files

You should execute this command in the root directory of your go project.

```
go run github.com/ispec-inc/civgen-go/model \
	// Required
	--name {model name} \
	--fields {fields of model} \

	// Required, but can be set by .civgen-model.yaml
	--project_path {your project root package path}
	--entity_path {path to the 'entity' package from 'project_path'}
	--model_path {path to the 'model' package from 'project_path'}
	--view_path {path to the 'view' package from 'project_path'}
	--repository_path {path to the 'repository' package from 'project_path'}
	--dao_path {path to the 'dao' package from 'project_path'}
	--error_path {path to the 'error' package from 'project_path'}
	--database_path {path to the 'database (e.g. mysql)' package from 'project_path'}

	// Optional
	--create_entity {create entity file, if true} (default: true)
	--create_model {create model file, if true} (default: true)
	--create_view {create view file, if true} (default: true)
	--create_repository {create repository file, if true} (default: true)
	--create_dao {create dao file, if true} (default: true)
	--create_dao_test {create dao test file, if true} (default: true)
```

#### Recommend

Config file to avoid all *_path options.

Put `<your project root>/.civgen-model.yaml` file with the following contents.

```yaml
project_path: 
entity_path: 
model_path: 
view_path: 
repository_path: 
dao_path: 
error_path: 
database_path:
```

### Example
Set [example/.civgen-model.yaml](example/.civgen-model.yaml).

Generate `User` model.
```
go run github.com/ispec-inc/civgen-go/model \
  --name User \
  --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time \
```

See [example/](./example/) for generated files.


## mockio

**This code is forked from https://github.com/golang/mock/tree/master/mockgen**

### Usage

Generate io struct of interface method.

```
go run github.com/ispec-inc/civgen-go/mockio \
	--source {source file path}
	--destination {destination file path}
	[other options which are same as github.com/golang/mock/mockgen]
```

### Example

Generate example mockio structs.

```
go run github.com/ispec-inc/civgen-go/mockio \
	--source example/pkg/domain/repository/user.go
	--destination example/pkg/domain/repository/mockio/user.go
```

```go
// Source
type User interface {
	Get(id int64) (model.User, apperror.Error)
	List(ids []int64) ([]model.User, apperror.Error)
}

// Generated
type UserGet struct {
	Time  int
	ArgId int64
	Ret0  model.User
	Ret1  apperror.Error
}

type UserList struct {
	Time   int
	ArgIds []int64
	Ret0   []model.User
	Ret1   apperror.Error
}
```

See [example/pkg/domain/repository/mockio/user.go](example/pkg/domain/repository/mockio/user.go) for more details.
