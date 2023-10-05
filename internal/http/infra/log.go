package infra

import (
	"github.com/evgenivanovi/gomart/pkg/std"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"golang.org/x/exp/slog"
)

func LogSuccessRequest(request any) {
	if request != nil {
		slogx.Log().Debug(
			"Incoming request",
			slog.Any("request", &request),
		)
	} else {
		slogx.Log().Debug(
			"Incoming request",
			slog.Any("request", nil),
		)
	}
}

func LogErrorRequest(err error) {
	if err != nil {
		slogx.Log().Error(
			"Incoming request",
			slog.String("error", err.Error()),
		)
	} else {
		slogx.Log().Error(
			"Incoming request",
			slog.String("error", std.Nil),
		)
	}
}

func LogSuccessResponse(response any) {
	if response != nil {
		slogx.Log().Debug(
			"Outcoming response",
			slog.Any("response", &response),
		)
	} else {
		slogx.Log().Debug(
			"Outcoming response",
			slog.Any("response", nil),
		)
	}
}

func LogErrorResponse(err error) {
	if err != nil {
		slogx.Log().Error(
			"Outcoming response",
			slog.String("error", err.Error()),
		)
	} else {
		slogx.Log().Error(
			"Outcoming response",
			slog.String("error", std.Nil),
		)
	}
}
