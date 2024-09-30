package setup

import (
	service "payment/app/server"
)

func Start(address string) {
	// setup up tracer
	// call the run
	srv := service.CreateServer(address)
	router := service.RouteInitilize(srv.Ctx)
	srv.Router = router
	srv.Run(address)
}
