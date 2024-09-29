package setup

import (
	service "payment/app/server"
)

func Start(address string) {
	// setup up tracer
	// call the run
	router := service.RouteInitilize()
	service.Run(address, router)
}
