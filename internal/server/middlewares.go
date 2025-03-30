package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Server) LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()

		method := ctx.Request.Method
		s.logger.Info("new request", "method", method, "URL", ctx.Request.URL.String())
		ctx.Next()

		statusCode := ctx.Writer.Status()

		s.logger.Info("request completed", "code", statusCode, "text", http.StatusText(statusCode), "duration", time.Since(now))
		ctx.Next()
	}
}
