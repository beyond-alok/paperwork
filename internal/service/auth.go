package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/beyond-alok/paperwork/internal/models"
	"github.com/beyond-alok/paperwork/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	store repository.UserStore
}

type RegisterReq struct {
	Name     string `json:"name" validate:"max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50,containsany=@!#$%&,containsany=0987654321"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	UserId string `json:"id"`
	jwt.RegisteredClaims
}

func NewAuthService(store repository.UserStore) *AuthService {
	return &AuthService{
		store: store,
	}
}

func (s *AuthService) Register(req *RegisterReq) (string, error) {
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
	user, err := s.store.GetByEmail(req.Email)
	if user != nil {
		return "", Error{
			Code: http.StatusBadRequest,
			Msg:  "user already exist",
			Err: err,
		}
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return "", Error{
			Code: http.StatusInternalServerError,
			Msg:  "internal server error",
			Err:  err,
		}
	}

	usr := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashBytes),
	}
	err = s.store.Create(&usr)
	if err != nil {
		return "", Error{
			Code: http.StatusInternalServerError,
			Msg:  "failed to create user",
			Err:  err,
		}
	}
	return "user registered successfull", nil
}

func (s *AuthService) Login(req *LoginReq) (string, error) {
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
	user, err := s.store.GetByEmail(req.Email)
	if err != nil {
		return "", Error{
			Code: http.StatusBadRequest,
			Msg:  "username or password doesn't exist",
			Err: err,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", Error{
			Code: http.StatusBadRequest,
			Msg:  "username or password doesn't exist",
			Err: err,
		}
	}

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	claims := Claims {
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72*time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signedToken,err := token.SignedString(jwtSecret)
	if err != nil {
		return "",Error{
			Code: http.StatusInternalServerError,
			Msg: "internal server error",
			Err: err,
		}
	}

	return signedToken,nil
}
