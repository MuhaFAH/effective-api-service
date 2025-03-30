package server

import (
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Server) LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()
		method := ctx.Request.Method

		s.logger.Info("new request", "method", method, "URL", ctx.Request.URL.String())
		ctx.Next()

		s.logger.Info("request completed", "code", ctx.Writer.Status(), "duration", time.Since(now))
		ctx.Next()
	}
}
