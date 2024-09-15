package engine

import (
	"net/http"

	"github.com/NoCapCbas/webadmin/data/model"
)

// Route represent a web handler with optional middlewares
type Route struct {
	// middleware
	Logger bool

	// authorization
	MinimumRole model.Roles

	Handler http.Handler
}
