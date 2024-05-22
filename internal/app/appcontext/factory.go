package appcontext

import (
	"github.com/dtrwi/datedive/config"
	"github.com/dtrwi/datedive/internal/app/driver"
	"github.com/jmoiron/sqlx"
)

// AppContext the app context struct
type AppContext struct {
	config config.Provider
}

// NewAppContext initiate appcontext object
func NewAppContext(config config.Provider) *AppContext {
	return &AppContext{
		config: config,
	}
}

func (a *AppContext) GetDBPostgreConn() (*sqlx.DB, error) {
	return driver.NewDBPostgre(a.getDBPostgreOption())
}

func (a *AppContext) getDBPostgreOption() driver.DBPostgreOption {
	return driver.DBPostgreOption{
		URL:         a.config.GetString("postgre.url"),
		MaxIdleConn: a.config.GetInt("postgre.max_idle_connections"),
		MaxOpenConn: a.config.GetInt("postgre.max_open_connections"),
	}
}
