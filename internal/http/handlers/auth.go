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
	var req service.UserRegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			decodeErr := service.Error{
				Code: http.StatusBadRequest,
				Msg:  "Decode Error",
				Err:  err,
			}
			slog.Error("registration failed", "error", decodeErr)
			response.WriteError(w, decodeErr.Code, decodeErr.Msg)
			return
		}

		slog.Error("registration failed", "error", err)
		response.WriteError(w, http.StatusBadRequest, "unexpected error")
		return
	}

	msg, err := h.authService.Register(&req)
	if err != nil {
		var svcError service.Error
		if errors.As(err, &svcError) {
			slog.Error("registration failed", "err", svcError)
			response.WriteError(w, svcError.Code, svcError.Msg)
			return
		}
		slog.Error("unexpected error", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.WriteJson(w, http.StatusAccepted, msg)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req service.LoginReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			decodeErr := service.Error{
				Code: http.StatusBadRequest,
				Msg:  "Decode Error",
				Err:  err,
			}
			slog.Error("login failed", "error", decodeErr)
			response.WriteError(w, decodeErr.Code, decodeErr.Msg)
			return
		}

		slog.Error("login failed", "error", err)
		response.WriteError(w, http.StatusBadRequest, "unexpected error")
		return
	}

	signedToken,err := h.authService.Login(&req)
	if err != nil {
		var svcError service.Error
		if errors.As(err, &svcError) {
			slog.Error("login failed", "err", svcError)
			response.WriteError(w, svcError.Code, svcError.Msg)
			return
		}
		slog.Error("unexpected error", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: signedToken,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteLaxMode,
		Path: "/",
		MaxAge: 72*60*60,
	})

	response.WriteJson(w,http.StatusOK,"login successful")

}

func (store *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,&http.Cookie{
		Name: "session_token",
		Value: "",
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteLaxMode,
		Path: "/",
		MaxAge: -1,
	})

	response.WriteJson(w,http.StatusOK,"logout successfully")
}
