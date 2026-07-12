package postgresql

import (
	"github.com/jackc/pgx"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{
		db : db,
	}
}

func (r *UserRepo) GetAll() {

}

func (r *UserRepo) GetById(id string) {

}

func (r *UserRepo) GetByEmail(email string)  {

}

func (r *UserRepo) Create() {

}

func (r *UserRepo) Update() {

}

func (r *UserRepo) Delete() {

}
