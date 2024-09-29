package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(address string, router http.Handler) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	srv := &http.Server{
		Addr:    ":" + address,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	<-signalChan
	shutDown(srv)
}

func shutDown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer func() {
		cancel()
	}()

	err := srv.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("SERVER IS SHUT")
}
