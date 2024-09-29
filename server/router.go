package service

import (
	"context"

	"github.com/gin-gonic/gin"
)

func RouteInitilize() *gin.Engine {
	router := gin.New()
	routerSec := router.Group("/api/v1/sec/")
	routerSec.GET("/world", GinTopGoHttp(HelloFunc))
	return router
}

type GohttpType func(*gin.Context, context.Context)

func GinTopGoHttp(handler GohttpType) func(ctx *gin.Context) {
	return func(gctx *gin.Context) {
		ctx := getFromCtx(gctx)

		handler(gctx, ctx)
	}
}

func getFromCtx(c *gin.Context) context.Context {
	ctx, isExist := c.Get("Key_CTX")
	if !isExist {
		ctx := context.Background()
		return ctx
	}

	return ctx.(context.Context)
}

func HelloFunc(c *gin.Context, ctx context.Context) {
	c.JSON(200, struct {
		Messaage string `json:"message"`
	}{Messaage: "Worked"})
}
