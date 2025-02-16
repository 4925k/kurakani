package kurakani

import (
	"net/http"

	"github.com/4925k/kurakani/foundation/logger"
	"github.com/4925k/kurakani/foundation/web"
)

func Routes(app *web.App, log *logger.Logger) {
	api := newApp()

	app.HandlerFunc(http.MethodGet, "", "/status", api.status)

}
