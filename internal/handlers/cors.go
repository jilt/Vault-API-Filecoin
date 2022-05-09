package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cors struct{}

func (c Cors) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")

		ctx.Header("Access-Control-Allow-Headers", "Authorization")

		ctx.Header("Access-Control-Allow-Headers", "Content-Type")

		if ctx.Request.Method == http.MethodOptions {
			ctx.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		ctx.Next()

	}
}
