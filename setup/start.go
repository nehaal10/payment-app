package setup

import (
	service "payment/app/server"

	"go.uber.org/zap"
)

func Start(address string, l *zap.Logger) {
	// setup up tracer
	// call the run
	srv := service.CreateServer(address)
	router := service.RouteInitilize(srv.Ctx)
	srv.Router = router
	srv.Run(address, l)
}
