package option

import "github.com/dtrwi/datedive/internal/app/appcontext"

type HandlerOption struct {
	AppOption
	Services *appcontext.Services
}
