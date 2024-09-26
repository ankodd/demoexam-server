package handlers

import (
	"encoding/json"
	"github.com/ankodd/demoexam/core/internal/service"
	"github.com/ankodd/demoexam/core/internal/storage"
	"github.com/ankodd/demoexam/core/internal/utils/errs"
	"github.com/ankodd/demoexam/core/internal/utils/msg"
	"github.com/ankodd/demoexam/core/internal/utils/parse/requestparse"
	"github.com/ankodd/demoexam/core/internal/utils/sl"
	"github.com/ankodd/demoexam/core/pkg/models"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	store  *storage.UserStorage
	logger *slog.Logger
}

func NewUserHandler(storage *storage.UserStorage, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		store:  storage,
		logger: logger,
	}
}

func (h *UserHandler) FetchUser(w http.ResponseWriter, r *http.Request) {
	// Parsing id
	id, err := requestparse.ParseID(r)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	users, err := service.UserFetch(id, h.store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		sl.ReqLog(http.StatusNotFound, h.logger, r, slog.LevelError)
		return
	}

	// Write response
	if err := Write(w, &users, http.StatusOK); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}

func (h *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	// FetchOrder users in storage
	users, err := service.UserFetchAll(h.store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
		return
	}

	// Write Response
	if err := Write(w, &users, http.StatusOK); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Parsing id
	id, err := requestparse.ParseID(r)
	if err != nil {
		http.Error(w, errs.InvalidIDErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Parsing user
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	if err := service.UserUpdate(id, &user, h.store); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		sl.ReqLog(http.StatusConflict, h.logger, r, slog.LevelError)
		return
	}

	// Write Response
	if err := Write(w, msg.New(msg.UpdateSuccess), http.StatusOK); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Parsing id
	id, err := requestparse.ParseID(r)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	if err := service.UserDelete(id, h.store); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
		return
	}

	// Write Response
	if err := Write(w, msg.New(msg.DeleteSuccess), http.StatusOK); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}
