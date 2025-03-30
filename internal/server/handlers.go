package server

import (
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) helloHandler(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Hello!"})
}

func (s *Server) createUserHandler(c *gin.Context) {
	var userData e.User
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.service.CreateUser(s.context, userData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (s *Server) getUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.service.ReadUser(s.context, uint(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
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
	userData.ID = uint(id)
	updatedUser, err := s.service.UpdateUser(s.context, userData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": updatedUser})
}

//TODO хендлер обновления пользователя

//TODO хендлер удаления пользователя
