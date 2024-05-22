package option

import "github.com/dtrwi/datedive/internal/app/appcontext"

type ServiceOption struct {
	AppOption
	Repositories *appcontext.Repositories
}
