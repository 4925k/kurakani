// Package mux provides routing support.
package mux

import (
	"context"
	"net/http"

	"github.com/4925k/kurakani/app/domain/kurakani"
	"github.com/4925k/kurakani/app/sdk/mid"
	"github.com/4925k/kurakani/foundation/logger"
	"github.com/4925k/kurakani/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log *logger.Logger
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config) http.Handler {
	logger := func(ctx context.Context, msg string, args ...any) {
		cfg.Log.Info(ctx, msg, args...)
	}

	app := web.NewApp(
		logger,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	kurakani.Routes(app, cfg.Log)

	return app
}
