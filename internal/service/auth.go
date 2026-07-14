package service

import (
	"net/http"

	"github.com/beyond-alok/paperwork/internal/repository"
	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	store repository.UserStore
}

type RegisterReq struct {
	Name     string `json:"name" validate:"max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password_hash" validate:"required,min=6,max=50,containsany=@!#$%&,containsany=0987654321"`
}

func NewAuthService(store repository.UserStore) *AuthService {
	return &AuthService{
		store: store,
	}
}

func (s *AuthService) Register(req RegisterReq) (string, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return "", Error{
			Code: http.StatusBadRequest,
			Msg:  "Validation Error",
			Err:  err,
		}
	}
	// s.store.GetByEmail(req.Email)
	return "", nil
}
