package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/beyond-alok/paperwork/internal/http/response"
	"github.com/beyond-alok/paperwork/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req service.RegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("registration failed", "error", err)
		if errors.Is(err, io.EOF) {
			decodeErr := service.Error{
				Code: http.StatusBadRequest,
				Msg: "Decode Error",
				Err: err,
			}
			response.WriteError(w,decodeErr.Code,decodeErr )
		}
		return
	}

	msg,err := h.authService.Register(req)
	if err != nil {
		var svcError service.Error
		if errors.As(err,&svcError) {
			slog.Error("registration failed","err",svcError)
			response.WriteError(w,svcError.Code,svcError)
		}
		slog.Error("unexpected error","error",err)
		response.WriteError(w,http.StatusInternalServerError,"internal server error")
	}
}
	


func (store *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login page"))
}

func (store *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}
