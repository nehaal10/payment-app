package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHello(c *gin.Context, ctx context.Context) {
	// get the context
	c.JSON(http.StatusAccepted, struct {
		Message string
	}{
		Message: "worked",
	})
}
