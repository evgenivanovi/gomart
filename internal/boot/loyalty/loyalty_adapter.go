package loyalty

import (
	"github.com/evgenivanovi/gomart/internal/boot/order"
	"github.com/evgenivanovi/gomart/internal/loyalty/adapter"
)

/* __________________________________________________ */

var (
	LoyaltyClient            adapter.LoyaltyClient
	LoyaltyOperationsService *adapter.LoyaltyOperationsService
)

/* __________________________________________________ */

func BootLoyaltyAdapter() {
	LoyaltyClient = LoyaltyClientService
	LoyaltyOperationsService = adapter.ProvideLoyaltyOperationsService(LoyaltyClient)
	order.OrderLoyaltyOperations = LoyaltyOperationsService
}

/* __________________________________________________ */
