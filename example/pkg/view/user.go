package view

import (
	"github.com/ispec-inc/civgen-go/example/pkg/domain/model"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewUser(m model.User) User {
	return User{

		CreatedAt: m.CreatedAt,
		Email:     m.Email,
		ID:        m.ID,
		Name:      m.Name,
		UpdateAt:  m.UpdateAt,
	}
}
