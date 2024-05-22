package cmd

import (
	"github.com/dtrwi/datedive/config"
	"github.com/dtrwi/datedive/internal/app/appcontext"
	"github.com/dtrwi/datedive/internal/app/option"
	"github.com/dtrwi/datedive/internal/app/server"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Execute adds all child commands.
// This is called by main.main(). It only needs to happen once.
func Execute() {
	var err error
	cfg := config.Config()
	app := appcontext.NewAppContext(cfg)

	var dbPostgre *sqlx.DB
	if cfg.GetBool("postgre.is_enabled") {
		dbPostgre, err = app.GetDBPostgreConn()
		if err != nil {
			logrus.Fatalf("failed to start, error connect to DB Postgre | %v", err)
			return
		}
		defer dbPostgre.Close()
	}

	opt := option.AppOption{
		Config:     cfg,
		PostgreSQL: dbPostgre,
	}

	server := server.NewServer(opt)
	server.StartApp()
}
