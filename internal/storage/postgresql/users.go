package postgresql

import (
	"errors"

	"github.com/beyond-alok/paperwork/internal/models"
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

func (r *UserRepo) GetByEmail(email string) (*models.User,error)  {
	var user models.User

	query := `SELECT id,name,email,password_hash FROM users WHERE email = $1`

	row := r.db.QueryRow(query,email)
	err := row.Scan(&user.ID,&user.Name,&user.Email,&user.PasswordHash)
	if err != nil {
		if errors.Is(err,pgx.ErrNoRows) {
			return nil,err
		}
		return nil,err
	}
	return &user,nil
}

func (r *UserRepo) Create(user *models.User) error {

	query := `INSERT INTO users (name,email,password_hash) VALUES($1,$2,$3) RETURNING id `

	row := r.db.QueryRow(query,user.Name,user.Email,user.PasswordHash)
	err := row.Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil

}

func (r *UserRepo) Update() {

}

func (r *UserRepo) Delete() {

}
