package bundb

import (
	"database/sql"

	"github.com/teampui/pac"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// ProviderDB will register BunDB into Pac's services
func ProvideDB(dsn string) pac.AppOption {
	if dsn == "" {
		panic("pac/bundb: cannot start, missed DSN settings")
	}

	return func(a *pac.App) {
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		a.Services.Add("db", bun.NewDB(sqldb, pgdialect.New()))
	}
}
