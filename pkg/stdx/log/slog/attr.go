package slog

import "golang.org/x/exp/slog"

func ErrAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}
