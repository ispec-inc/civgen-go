# mockio

**This code is forked from https://github.com/golang/mock/tree/master/mockgen**

## Usage

Generate io struct of interface method.

```
go run github.com/ispec-inc/civgen-go/mockio \
	--source {source file path}
	--destination {destination file path}
	[other options which are same as github.com/golang/mock/mockgen]
```

## Example

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

See [example/pkg/domain/repository/mockio/user.go](../example/pkg/domain/repository/mockio/user.go) for more details.