package server

import (
	"context"
	"errors"
	"github.com/MuhaFAH/effective-api-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine  *gin.Engine
	service *service.Service

	context context.Context
}

func NewServer(ctx context.Context, srv *service.Service) *Server {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	return &Server{engine: router, service: srv, context: ctx}
}

func (s *Server) InitHandlers() {
	s.engine.GET("/", s.helloHandler)
	s.engine.GET("/user/get/:id", s.getUserHandler)
	s.engine.GET("/user/delete/:id", s.deleteUserHandler)

	s.engine.POST("/user/create", s.createUserHandler)
	s.engine.PATCH("/user/update/:id", s.updateUserHandler)
}

func (s *Server) Run(addr string) error {
	if s.engine == nil {
		return errors.New("engine not initialized")
	}

	err := s.engine.Run(addr)
	if err != nil {
		return err
	}

	return nil
}
