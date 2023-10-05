package loyalty

import (
	"github.com/evgenivanovi/gomart/internal/boot/app"
	"github.com/evgenivanovi/gomart/internal/loyalty/infra"
	"github.com/evgenivanovi/gomart/pkg/fw"
	"github.com/evgenivanovi/gomart/pkg/stdx/net/restyx"
	"github.com/go-resty/resty/v2"
)

/* __________________________________________________ */

var (
	LoyaltyHTTPClient    *resty.Client
	LoyaltyClientService *infra.LoyaltyClientService
)

/* __________________________________________________ */

func BootLoyaltyInfrastructure() {
	initLoyaltyHTTPClient()
	initLoyaltyClient()
}

func initLoyaltyHTTPClient() {

	address, err := app.AccrualProperty.CalcElse(fw.FirstStringNotEmptyElse())
	if err != nil {
		panic(err)
	}

	LoyaltyHTTPClient = resty.
		New().
		SetBaseURL(address).
		SetRetryAfter(restyx.RetryOnManyRequestsAfterSeconds())

}

func initLoyaltyClient() {
	LoyaltyClientService = infra.ProvideLoyaltyClientService(LoyaltyHTTPClient)
}

/* __________________________________________________ */
