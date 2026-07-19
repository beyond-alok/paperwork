package repository

import (
	"github.com/beyond-alok/paperwork/internal/models"
	"github.com/beyond-alok/paperwork/internal/storage/postgresql"
)

type TenantStore interface {
	GetAll()
	GetById(string)
	GetByUsername(string) (*models.Tenant,error)
	Create(*models.Tenant) error
	Update()
	Delete()

}

var _ TenantStore = (*postgresql.TenantRepo)(nil)