package option

import (
	"github.com/dtrwi/datedive/config"

	"github.com/jmoiron/sqlx"
)

type AppOption struct {
	Config     config.Provider
	PostgreSQL *sqlx.DB
}
