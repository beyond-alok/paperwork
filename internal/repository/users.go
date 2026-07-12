package repository

import (
	"github.com/beyond-alok/paperwork/internal/storage/postgresql"
)

type UserStore interface {
	GetAll()
	GetById(string)
	GetByEmail(string)
	Create()
	Update()
	Delete()
}

var _ UserStore = (*postgresql.UserRepo)(nil)
