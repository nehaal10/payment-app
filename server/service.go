package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type HttpServer struct {
	Address string
	Stop    context.CancelFunc
	ErrGrp  *errgroup.Group
	Ctx     context.Context
	Router  http.Handler
	Server  *http.Server
}

func (srv *HttpServer) Run(address string, l *zap.Logger) {
	server := &http.Server{
		Addr:    ":" + address,
		Handler: srv.Router,
	}

	srv.Server = server

	go func() {
		if err := server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-srv.Ctx.Done()
	srv.Stop()
	srv.shutDown()
}

func CreateServer(address string) (srv *HttpServer) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	errGrp, errCtx := errgroup.WithContext(ctx)

	srv = &HttpServer{
		Address: address,
		Stop:    stop,
		ErrGrp:  errGrp,
		Ctx:     errCtx,
	}

	return srv
}

func (srv *HttpServer) shutDown() {
	_, cancel := context.WithTimeout(srv.Ctx, time.Duration(5*time.Second))
	defer func() {
		cancel()
	}()
	err := srv.Server.Shutdown(srv.Ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("SERVER IS SHUT")
}
