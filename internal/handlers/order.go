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

type OrderHandler struct {
	store  *storage.OrderStorage
	logger *slog.Logger
}

func NewOrderHandler(storage *storage.OrderStorage, logger *slog.Logger) *OrderHandler {
	return &OrderHandler{
		store:  storage,
		logger: logger,
	}
}

func (h *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	if err := service.OrderAdd(&order, h.store); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
		return
	}

	if err := Write(w, msg.New(msg.OrderCreateSuccess), http.StatusOK); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}
}

func (h *OrderHandler) FetchOrder(w http.ResponseWriter, r *http.Request) {
	// Parsing id
	id, err := requestparse.ParseID(r)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	users, err := service.OrderFetch(id, h.store)
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

func (h *OrderHandler) FetchOrders(w http.ResponseWriter, r *http.Request) {
	// FetchOrder orders in storage
	orders, err := service.OrderFetchAll(h.store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
		return
	}

	// Write Response
	if err := Write(w, &orders, http.StatusOK); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Parsing id
	id, err := requestparse.ParseID(r)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Parsing Order
	var order models.Order
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	if err := service.OrderUpdate(id, &order, h.store); err != nil {
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

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Parsing id
	id, err := requestparse.ParseID(r)
	if err != nil {
		http.Error(w, errs.BadRequestErr, http.StatusBadRequest)
		sl.ReqLog(http.StatusBadRequest, h.logger, r, slog.LevelError)
		return
	}

	// Service logic
	if err := service.OrderDelete(id, h.store); err != nil {
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
