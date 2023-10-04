package slog

import (
	zapx "github.com/evgenivanovi/gomart/pkg/stdx/log/zap"
	"golang.org/x/exp/slog"
)

/* __________________________________________________ */

var logger *slog.Logger

/* __________________________________________________ */

func init() {
	logger = zapx.LogAsStructured(*zapx.Log())
}

func Log() *slog.Logger {
	return logger
}
