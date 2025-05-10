package checkapi

import (
	"github.com/ardanlabs/service/foundation/web"
)

func Routes(app *web.App) {
	app.HandleFunc("GET /v1/liveness", liveness)
	app.HandleFunc("GET /v1/readiness", readiness)
}
