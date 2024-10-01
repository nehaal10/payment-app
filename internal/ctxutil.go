package internal

import (
	"context"

	"github.com/gin-gonic/gin"
)

const (
	Key_Ctx string = "Key_CTX"
)

func GetCTXFromGinCtx(c *gin.Context) context.Context {
	ctx, isExist := c.Get(Key_Ctx)
	if !isExist {
		ctx := context.Background()
		return ctx
	}

	return ctx.(context.Context)
}
