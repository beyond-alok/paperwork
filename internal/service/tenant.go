package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/beyond-alok/paperwork/internal/models"
	"github.com/beyond-alok/paperwork/internal/repository"
	"github.com/go-playground/validator/v10"
)

type TenantService struct {
	store repository.TenantStore
}

type TenantRegisterReq struct {
	Name     string `json:"name" validate:"required"`
	// TODO : add more validation feild
	Username string `json:"username" validate:"required"`
}

func NewTenantService(store repository.TenantStore) *TenantService {
	return &TenantService{
		store: store,
	}
}

func (s *TenantService) Register(req *TenantRegisterReq) (string, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		var validationErr validator.ValidationErrors
		if errors.As(err, &validationErr) {
			msgs := make([]string, 0, len(validationErr))
			for _, feildErr := range validationErr {
				msgs = append(msgs, fmt.Sprintf("%s faild on %s validation", feildErr.Field(), feildErr.Tag()))
			}
			return "", Error{
				Code: http.StatusBadRequest,
				Msg:  strings.Join(msgs, ";"),
				Err:  err,
			}
		}
		return "", Error{
			Code: http.StatusInternalServerError,
			Msg:  "internal server error",
			Err:  err,
		}

	}
	user, err := s.store.GetByUsername(req.Username)
	if user != nil {
		return "", Error{
			Code: http.StatusBadRequest,
			Msg:  "tenant already exist",
			Err:  err,
		}
	}

	usr := models.Tenant{
		Name:         req.Name,
		Username:        req.Username,
	}

	err = s.store.Create(&usr)
	if err != nil {
		return "", Error{
			Code: http.StatusInternalServerError,
			Msg:  "failed to create user",
			Err:  err,
		}
	}
	return "tenant registered successfull", nil

}
