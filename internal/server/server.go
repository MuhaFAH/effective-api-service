package server

import (
	"context"
	"errors"
	_ "github.com/MuhaFAH/effective-api-service/docs"
	"github.com/MuhaFAH/effective-api-service/internal/logger"
	"github.com/MuhaFAH/effective-api-service/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine  *gin.Engine
	service *service.Service

	logger  logger.MinimalisticLogger
	context context.Context
}

func NewServer(ctx context.Context, srv *service.Service, lggr logger.MinimalisticLogger) *Server {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	return &Server{engine: router, service: srv, logger: lggr, context: ctx}
}

func (s *Server) InitHandlers() {
	s.engine.Use(s.LoggingMiddleware())

	s.engine.GET("/", s.helloHandler)
	s.engine.GET("/user/get/:id", s.getUserHandler)
	s.engine.DELETE("/user/delete/:id", s.deleteUserHandler)

	s.engine.POST("/users/get", s.getUsersHandler)
	s.engine.POST("/user/create", s.createUserHandler)
	s.engine.PATCH("/user/update/:id", s.updateUserHandler)

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) Run(addr string) error {
	if s.engine == nil {
		return errors.New("engine not initialized")
	}

	s.logger.Infof("server started work on address: %s", addr)
	err := s.engine.Run(addr)
	if err != nil {
		return err
	}

	return nil
}
