package pg

import (
	"context"

	"github.com/evgenivanovi/gomart/internal/boot/app"
	"github.com/evgenivanovi/gomart/pkg/fw"
	"github.com/evgenivanovi/gomart/pkg/pg"
	slogx "github.com/evgenivanovi/gomart/pkg/stdx/log/slog"
	"github.com/gookit/goutil/strutil"
)

/* __________________________________________________ */

var (
	DSN        string
	Datasource *pg.Datasource
)

/* __________________________________________________ */

func BootPGDatasource() {

	err := initDSN()
	if err != nil {
		panic(err)
	}

	err = initDatasource()
	if err != nil {
		panic(err)
	}

	initTasks()

}

func initDSN() error {

	dsn, err := app.DSNPostgresProperty.CalcElse(fw.FirstStringNotEmptyElse())
	if err != nil {
		// We didn't find property flag or environment variable for connection to a datasource.
		return err
	}

	DSN = dsn
	return nil

}

func initDatasource() error {

	src, err := pg.NewDatasource(
		context.Background(),
		DSN,
		*pg.NewConnectionSettings(),
	)
	if err != nil {
		return err
	}

	Datasource = src
	return nil

}

func initTasks() {

	app.Application.RegisterOnClose(
		func() {
			slogx.Log().Debug(
				"Closing PostgreSQL connection.",
			)
			Datasource.Pool.Close()
			slogx.Log().Debug(
				"Closed PostgreSQL connection.",
			)
		},
	)

}

/* __________________________________________________ */

func IsInitializedPG() bool {
	if strutil.IsEmpty(DSN) || Datasource == nil {
		return false
	}
	return true
}

/* __________________________________________________ */
