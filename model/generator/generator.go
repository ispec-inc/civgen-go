package generator

import (
	"github.com/ispec-inc/civgen-go/model/value"
)

type Generator interface {
	Model(path value.Filepath, layer value.Layer) error
	Repository(path value.Filepath) error
	Dao(path value.Filepath) error
	DaoTest(path value.Filepath) error
}

type generator struct {
	name   string
	fields string
}

func NewGenerator(name, fields string) Generator {
	return generator{
		name:   name,
		fields: fields,
	}
}
