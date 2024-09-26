package sl

import (
	"log/slog"
	"net/http"
	"os"
)

func New() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}

func ReqLog(status int, sl *slog.Logger, r *http.Request, logLevel slog.Level) {
	switch logLevel {
	case slog.LevelError:
		sl.Error(
			"incoming request",
			slog.String("method", r.Method),
			slog.Int("status", status),
			slog.String("url", r.URL.String()),
			slog.String("user_agent", r.UserAgent()),
		)
	case slog.LevelInfo:
		sl.Info(
			"incoming request",
			slog.String("method", r.Method),
			slog.Int("status", status),
			slog.String("url", r.URL.String()),
			slog.String("user_agent", r.UserAgent()),
		)
	default:
		return
	}
}

func Err(sl *slog.Logger, err error) {
	sl.Error("error", slog.String("error", err.Error()))
}
