package kurakani

import (
	"context"
	"net/http"

	"github.com/4925k/kurakani/foundation/web"
)

type app struct {
}

func newApp() *app {
	return &app{}
}

func (a *app) status(_ context.Context, _ *http.Request) web.Encoder {
	return status{
		Status: "ok",
	}
}
