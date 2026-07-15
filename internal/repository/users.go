package repository

import (
	"github.com/beyond-alok/paperwork/internal/models"
	"github.com/beyond-alok/paperwork/internal/storage/postgresql"
)

type UserStore interface {
	GetAll()
	GetById(string)
	GetByEmail(string) (*models.User,error)
	Create(*models.User) error
	Update()
	Delete()
}

var _ UserStore = (*postgresql.UserRepo)(nil)
