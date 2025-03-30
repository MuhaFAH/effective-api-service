package server

import (
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// @Summary Получить приветствие от сервера
// @Description Базовая и удобная проверка работоспособности ответа от сервера
// @Tags main
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func (s *Server) helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello!"})
}

func (s *Server) createUserHandler(c *gin.Context) {
	var userData e.User
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()

	user, err := s.service.CreateUser(s.context, userData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	s.logger.Debug("user created", "id", user.ID, "duration", time.Since(now))
	c.JSON(200, gin.H{"user": user})
}

func (s *Server) getUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()

	user, err := s.service.ReadUser(s.context, uint(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	s.logger.Debug("found user data", "id", user.ID, "duration", time.Since(now))
	c.JSON(200, gin.H{"user": user})
}

func (s *Server) getUsersHandler(c *gin.Context) {
	var usersData UsersRequest
	if err := c.BindJSON(&usersData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()

	users, err := s.service.ReadUsersByFilter(s.context, convertIntoFilter(usersData))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	s.logger.Debug("found users data", "duration", time.Since(now))
	c.JSON(200, gin.H{"users": users})
}

func (s *Server) updateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userData e.User
	if err = c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()

	updatedUser, err := s.service.UpdateUser(s.context, uint(id), userData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	s.logger.Debug("updated user info", "duration", time.Since(now))
	c.JSON(200, gin.H{"user": updatedUser})
}

func (s *Server) deleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()

	user, err := s.service.DeleteUser(s.context, uint(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	s.logger.Debug("deleted user", "id", id, "duration", time.Since(now))
	c.JSON(200, gin.H{"user": user})
}
