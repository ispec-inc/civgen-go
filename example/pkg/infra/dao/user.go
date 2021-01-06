package dao

import (
	"github.com/ispec-inc/civgen-go/example/pkg/apperror"
	"github.com/ispec-inc/civgen-go/example/pkg/domain/model"
	"github.com/ispec-inc/civgen-go/example/pkg/infra/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return User{
		db,
	}
}
func (d User) Get(id int64) (model.User, apperror.Error) {
	var e entity.User
	if err := d.db.First(&e, id).Error; err != nil {
		return model.User{}, apperror.NewGormFind(err, entity.UserModelName)
	}
	return e.ToModel(), nil
}

func (d User) List(ids []int64) (ms []model.User, aerr apperror.Error) {
	query := d.db
	if len(ids) > 0 {
		query = query.Where("id in (?)", ids)
	}

	var es []entity.User
	if err := query.Find(&es).Error; err != nil {
		return []model.User{}, apperror.NewGormFind(err, entity.UserModelName)
	}

	for _, e := range es {
		ms = append(ms, e.ToModel())
	}
	return ms, nil
}
