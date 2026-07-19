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

type TenantHandler struct {
	tenantService *service.TenantService
}

func NewTenantHandler(tenantService *service.TenantService) *TenantHandler {
	return &TenantHandler{
		tenantService: tenantService,
	}
}

func (h *TenantHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req service.TenantRegisterReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			decodeErr := service.Error{
				Code: http.StatusBadRequest,
				Msg:  "Decode Error",
				Err:  err,
			}
			slog.Error("tenant registration failed", "error", decodeErr)
			response.WriteError(w, decodeErr.Code, decodeErr.Msg)
			return
		}

		slog.Error("tenant registration failed", "error", err)
		response.WriteError(w, http.StatusBadRequest, "unexpected error")
		return
	}

	msg, err := h.tenantService.Register(&req)
	if err != nil {
		var svcError service.Error
		if errors.As(err, &svcError) {
			slog.Error("tenant registration failed", "err", svcError)
			response.WriteError(w, svcError.Code, svcError.Msg)
			return
		}
		slog.Error("unexpected error", "error", err)
		response.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.WriteJson(w, http.StatusAccepted, msg)
}
