package response

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/beyond-alok/paperwork/internal/service"
)



func write(w http.ResponseWriter, status int, body any) error {
	// marshal before writing headers, so we can catch encoding errors
	// before commiting to the response
	v, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(v)

	return nil
}

func writeStream(w http.ResponseWriter, status int, body any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return err
	}
	return nil
}

func WriteJson(w http.ResponseWriter, status int, successRes any) error {
	err := write(w, status, successRes)
	if err != nil {
		slog.Error(" WriteJson: failed to marshal response", "error", err, "body", successRes)
		return err
	}
	return nil
}

func WriteJsonStream(w http.ResponseWriter, status int, successRes service.Success) error {
	err := writeStream(w, status, successRes)
	if err != nil {
		slog.Error(" WriteJsonStream: failed to stream response", "error", err, "body", successRes)
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, status int, errRes any) error {
	err := write(w, status, errRes)
	if err != nil {
		slog.Error(" WriteError: failed to marshal response", "error", err, "body", errRes)
		return err
	}
	return nil
}
