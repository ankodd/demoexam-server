package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Write(w http.ResponseWriter, obj interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("failed marshalling response: %w", err)
	}

	if _, err := w.Write(bytes); err != nil {
		return fmt.Errorf("failed writing response: %w", err)
	}

	return nil
}
