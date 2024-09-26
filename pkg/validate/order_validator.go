package validate

import (
	"errors"
	"github.com/ankodd/demoexam/core/pkg/models"
)

func OrderStatus(status models.Status) error {
	switch status {
	case "waiting":
		return nil
	case "working":
		return nil
	case "done":
		return nil
	default:
		return errors.New("invalid order status")
	}
}
