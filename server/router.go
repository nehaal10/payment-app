package service

import (
	"context"
	"payment/app/handlers"
	"payment/app/utils"

	"github.com/gin-gonic/gin"
)

// handler when getting the first middle ware then pass the server context set the server cont to gin context and from other middle ware the context from
// gin contect and add data

func RouteInitilize(serverCtx context.Context) *gin.Engine {
	router := gin.New()
	routerSec := router.Group("/api/v1/sec/")
	routerSec.GET("/world", GinTopGoHttp(handlers.LoginHello))
	return router
}

type GohttpType func(*gin.Context, context.Context)

func GinTopGoHttp(handler GohttpType) func(ctx *gin.Context) {
	return func(gctx *gin.Context) {
		ctx := utils.GetCTXFromGinCtx(gctx)

		handler(gctx, ctx)
	}
}
