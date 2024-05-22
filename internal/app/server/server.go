package server

import (
	"fmt"
	"net/http"

	"github.com/dtrwi/datedive/internal/app/appcontext"
	"github.com/dtrwi/datedive/internal/app/option"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IServer interface {
	StartApp()
}

type server struct {
	opt option.AppOption
}

// NewServer create object server
func NewServer(opt option.AppOption) IServer {
	return &server{
		opt: opt,
	}
}

func (s *server) StartApp() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderAccept, "x-app-token"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut},
	}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == s.opt.Config.GetString("app.secret"), nil
		},
	}))

	repositories := wiringRepository(option.RepositoryOption{
		AppOption: s.opt,
	})

	services := wiringService(option.ServiceOption{
		AppOption:    s.opt,
		Repositories: repositories,
	})

	handlers := option.HandlerOption{
		AppOption: s.opt,
		Services:  services,
	}

	Router(handlers, e)

	address := fmt.Sprintf(":%d", s.opt.Config.GetInt("app.port"))
	e.Logger.Fatal(e.Start(address))
}

func wiringRepository(opt option.RepositoryOption) *appcontext.Repositories {
	// wiring up all your repos here
	repo := appcontext.Repositories{}

	return &repo
}

func wiringService(opt option.ServiceOption) *appcontext.Services {
	// wiring up all services
	svc := appcontext.Services{}
	return &svc
}
