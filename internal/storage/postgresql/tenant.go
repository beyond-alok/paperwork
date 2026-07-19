package postgresql

import (
	"errors"

	"github.com/beyond-alok/paperwork/internal/models"
	"github.com/jackc/pgx"
)

type TenantRepo struct {
	db *pgx.Conn
}

func NewTenantRepo(db *pgx.Conn) *TenantRepo {
	return &TenantRepo{
		db: db,
	}
}

func (r *TenantRepo) GetAll() {

}

func (r *TenantRepo) GetById(id string) {

}

func (r *TenantRepo) GetByUsername(email string) (*models.Tenant, error) {
	var tenant models.Tenant

	query := `SELECT id,name,username FROM users WHERE username= $1`

	row := r.db.QueryRow(query, email)
	err := row.Scan(&tenant.ID, &tenant.Name, &tenant.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &tenant, nil
}

func (r *TenantRepo) Create(tenant *models.Tenant) error {

	query := `INSERT INTO users (name,username) VALUES($1,$2,$3) RETURNING id `

	row := r.db.QueryRow(query, tenant.Name, tenant.Username)
	err := row.Scan(&tenant.ID)
	if err != nil {
		return err
	}

	return nil

}

func (r *TenantRepo) Update() {

}

func (r *TenantRepo) Delete() {

}
