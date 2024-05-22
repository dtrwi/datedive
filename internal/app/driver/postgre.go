package driver

import (
	"github.com/jmoiron/sqlx"
)

// DBPostgreOption for postgresql connection
type DBPostgreOption struct {
	URL         string
	MaxIdleConn int
	MaxOpenConn int
}

// NewDBPostgre return a client connection handle to a Postgre server.
func NewDBPostgre(option DBPostgreOption) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", option.URL)
	if err != nil {
		return db, err
	}

	db.SetMaxIdleConns(option.MaxIdleConn)
	db.SetMaxOpenConns(option.MaxOpenConn)

	return db, nil
}
