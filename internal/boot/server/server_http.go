package server

import (
	"net/http"

	"github.com/evgenivanovi/gomart/internal/boot/account"
	"github.com/evgenivanovi/gomart/internal/boot/app"
	"github.com/evgenivanovi/gomart/internal/boot/auth"
	"github.com/evgenivanovi/gomart/internal/boot/order"
	"github.com/evgenivanovi/gomart/internal/http/infra"
	"github.com/evgenivanovi/gomart/pkg/fw"
	"github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"github.com/evgenivanovi/gomart/pkg/stdx/mw"
	"github.com/go-chi/chi/v5"
)

/* __________________________________________________ */

var (
	Router chi.Mux
)

/* __________________________________________________ */

func init() {
	Router = *chi.NewRouter()
}

/* __________________________________________________ */

func BootHTTPServer() {
	bootstrapHTTPRouter()
	bootstrapHTTPServer()
}

func bootstrapHTTPServer() {

	err := fw.ConfigureServer(&app.Application, &Router)
	if err != nil {
		slog.Log().Debug("bootstrap server failed", slog.ErrAttr(err))
	}

}

func bootstrapHTTPRouter() {

	Router.NotFound(defaultErrorHandler)
	Router.MethodNotAllowed(defaultErrorHandler)

	Router.Post(
		"/api/user/login",
		mw.Conveyor(
			infra.AsHandlerFunc(auth.SignInController),
		).ServeHTTP,
	)

	Router.Post(
		"/api/user/register",
		mw.Conveyor(
			infra.AsHandlerFunc(auth.SignUpController),
		).ServeHTTP,
	)

	Router.Post(
		"/api/user/orders",
		mw.Conveyor(
			infra.AsHandlerFunc(order.OrderLoadController),
			auth.AuthMW,
		).ServeHTTP,
	)

	Router.Get(
		"/api/user/orders",
		mw.Conveyor(
			infra.AsHandlerFunc(order.GetOrdersController),
			auth.AuthMW,
		).ServeHTTP,
	)

	Router.Get(
		"/api/user/balance",
		mw.Conveyor(
			infra.AsHandlerFunc(account.GetBalanceController),
			auth.AuthMW,
		).ServeHTTP,
	)

	Router.Post(
		"/api/user/balance/withdraw",
		mw.Conveyor(
			infra.AsHandlerFunc(account.WithdrawController),
			auth.AuthMW,
		).ServeHTTP,
	)

	Router.Get(
		"/api/user/withdrawals",
		mw.Conveyor(
			infra.AsHandlerFunc(account.WithdrawalHistoryController),
			auth.AuthMW,
		).ServeHTTP,
	)

}

func defaultErrorHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusBadRequest)
	_, _ = writer.Write(nil)
}

/* __________________________________________________ */
