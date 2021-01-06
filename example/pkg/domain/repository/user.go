package repository

import (
	"github.com/ispec-inc/civgen-go/example/pkg/apperror"
	"github.com/ispec-inc/civgen-go/example/pkg/domain/model"
)

type User interface {
	Get(id int64) (model.User, apperror.Error)
	List(ids []int64) ([]model.User, apperror.Error)
}
