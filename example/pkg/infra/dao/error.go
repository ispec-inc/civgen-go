package dao

import (
	"fmt"
	"github.com/ispec-inc/civgen-go/example/pkg/apperror"
	"gorm.io/gorm"
)

func newGormError(err error, msg string) apperror.Error {
	switch err {
	case gorm.ErrRecordNotFound:
		return apperror.New(apperror.CodeNotFound, fmt.Errorf("%s:%s", msg, err.Error()))
	default:
		return apperror.New(apperror.CodeError, fmt.Errorf("%s:%s", msg, err.Error()))
	}
}
