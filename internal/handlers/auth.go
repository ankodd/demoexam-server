package handlers

import (
	"encoding/json"
	"github.com/ankodd/demoexam/core/internal/service"
	"github.com/ankodd/demoexam/core/internal/utils/errs"
	"github.com/ankodd/demoexam/core/internal/utils/msg"
	"github.com/ankodd/demoexam/core/internal/utils/sl"
	"github.com/ankodd/demoexam/core/pkg/models"
	"log/slog"
	"net/http"
)

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Parsing user
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	if err := service.Login(&user, h.store); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		sl.ReqLog(http.StatusUnauthorized, h.logger, r, slog.LevelError)
		return
	}

	// Write response
	err = Write(w, map[string]int64{"id": user.ID}, http.StatusOK)
	if err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Parsing request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	if err := service.Register(&user, h.store); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		sl.ReqLog(http.StatusUnauthorized, h.logger, r, slog.LevelError)
		return
	}

	// Write response
	err = Write(w, msg.New(msg.RegistrationSuccess), http.StatusOK)
	if err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}
