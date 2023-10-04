package boot

import (
	"github.com/evgenivanovi/gomart/internal/boot/account"
	"github.com/evgenivanovi/gomart/internal/boot/app"
	"github.com/evgenivanovi/gomart/internal/boot/auth"
	"github.com/evgenivanovi/gomart/internal/boot/common"
	"github.com/evgenivanovi/gomart/internal/boot/loyalty"
	"github.com/evgenivanovi/gomart/internal/boot/order"
	"github.com/evgenivanovi/gomart/internal/boot/pg"
	"github.com/evgenivanovi/gomart/internal/boot/server"
)

/* __________________________________________________ */

func Boot() {
	app.BootApp()
	BootPG()
	BootCommon()
	BootAccount()
	BootAuth()
	BootLoyalty()
	BootOrder()
	server.BootHTTPServer()
}

func BootPG() {
	pg.BootPGDatasource()
	pg.BootPGInfrastructure()
	pg.BootPGMigrations()
}

func BootCommon() {
	common.BootCommon()
}

func BootAccount() {
	account.BootBalance()
	account.BootWithdraw()
	account.BootAccount()
	account.BootAccountHandlers()
}

func BootAuth() {
	auth.BootAuthToken()
	auth.BootAuthSession()
	auth.BootAuthUser()
	auth.BootAuthMW()
	auth.BootAuthHandlers()
}

func BootLoyalty() {
	loyalty.BootLoyaltyInfrastructure()
	loyalty.BootLoyaltyAdapter()
}

func BootOrder() {
	order.BootOrder()
	order.BootOrderHandlers()
}

/* __________________________________________________ */
