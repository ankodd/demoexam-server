package handlers

import (
	"github.com/ankodd/demoexam/core/internal/service"
	"github.com/ankodd/demoexam/core/internal/utils/errs"
	"github.com/ankodd/demoexam/core/internal/utils/sl"
	"log/slog"
	"net/http"
)

func (h *OrderHandler) Statistics(w http.ResponseWriter, r *http.Request) {
	// Get statistics
	statistics, err := service.Statistics(h.store)
	if err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
		return
	}

	// Write response
	if err := Write(w, &statistics, http.StatusOK); err != nil {
		http.Error(w, errs.InternalServerErr, http.StatusInternalServerError)
		sl.ReqLog(http.StatusInternalServerError, h.logger, r, slog.LevelError)
		return
	}

	sl.ReqLog(http.StatusOK, h.logger, r, slog.LevelInfo)
}
