// Code generated by MockIO. DO NOT EDIT.
// Source: example/pkg/domain/repository/user.go

// Package mockio_repository is a generated GoMock package.
package mockio_repository

import (
	apperror "github.com/ispec-inc/civgen-go/example/pkg/apperror"
	model "github.com/ispec-inc/civgen-go/example/pkg/domain/model"
)

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