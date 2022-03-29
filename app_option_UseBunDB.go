package pac

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func UseBunDB(dsn string) AppOption {
	return func(a *App) {
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		a.RegisterService("db", bun.NewDB(sqldb, pgdialect.New()))
	}
}
