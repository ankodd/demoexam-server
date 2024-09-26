package requestparse

import (
	"errors"
	"github.com/ankodd/demoexam/core/internal/utils/errs"
	"net/http"
	"strconv"
)

func ParseID(r *http.Request) (int64, error) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, errors.New(errs.InvalidIDErr)
	}

	return id, nil
}

func ParseChatID(r *http.Request) (int64, error) {
	idParam := r.URL.Query().Get("chat_id")
	chatID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, errors.New(errs.InvalidChatIDErr)
	}

	return chatID, nil
}
