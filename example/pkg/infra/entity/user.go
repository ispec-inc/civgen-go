package entity

import (
	"github.com/ispec-inc/civgen-go/example/pkg/domain/model"
	"time"
)

const UserModelName = "User"

type User struct {
	ID        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdateAt  time.Time `gorm:"column:update_at"`
}

func (e User) ToModel() model.User {
	return model.User{

		CreatedAt: e.CreatedAt,
		Email:     e.Email,
		ID:        e.ID,
		Name:      e.Name,
		UpdateAt:  e.UpdateAt,
	}
}
