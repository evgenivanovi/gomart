package pg

import (
	"github.com/evgenivanovi/gomart/pkg/pg"
)

/* __________________________________________________ */

var (
	PostgresReadRequester  pg.ReadRequester
	PostgresWriteRequester pg.WriteRequester
	PostgresTransactor     pg.TrxRequester
)

/* __________________________________________________ */

func BootPGInfrastructure() {

	PostgresReadRequester = *pg.ProvideReadRequester(Datasource.Pool)
	PostgresWriteRequester = *pg.ProvideWriteRequester(Datasource.Pool)
	PostgresTransactor = *pg.ProvideTrxRequester(Datasource.Pool)

}
