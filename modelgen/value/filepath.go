package value

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

type Filepath string

func (p Filepath) String() string {
	return string(p)
}

func NewFilepath(path, name, suffix string) Filepath {
	return Filepath(fmt.Sprintf("%s/%s%s.go", path, strcase.ToSnake(name), suffix))
}
