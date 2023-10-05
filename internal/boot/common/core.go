package common

import (
	"github.com/evgenivanovi/gomart/internal/boot/pg"
	"github.com/evgenivanovi/gomart/internal/domain/common"
)

/* __________________________________________________ */

var (
	Transactor common.Transactor
)

/* __________________________________________________ */

func BootCommon() {
	Transactor = pg.PostgresTransactor
}

/* __________________________________________________ */
