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
		slog.Error("failed to decode body", "error :", err)
		if errors.Is(err, io.EOF) {
			decodeErr := response.Error{
				Msg: "Decode Error",
				Err: err.Error(),
			}
			response.WriteError(w, http.StatusBadRequest,decodeErr )
		}
		return
	}

	h.authService.Register(req)

}

func (store *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login page"))
}

func (store *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}
