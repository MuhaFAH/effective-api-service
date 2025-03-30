package server

import (
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// @Router / [get]
// @Summary Базовое приветствие
// @Description Базовая и удобная проверка работоспособности ответа от сервера.
// @Tags main
// @ID hello-handler
// @Produce json
// @Success 200 {object} map[string]string
func (s *Server) helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello!"})
}

// @Router /user/create [post]
// @Summary Создание пользователя
// @Description Позволяет создать нового пользователя
// @Tags users
// @ID create-user
// @Accept json
// @Produce json
// @Param input body GetUserExample true "Информация о пользователе (обязательны только имя и фамилия)"
// @Success 200 {object} OkResponseExample "Успешная обработка данных"
// @Failure 400 {object} ErrorResponseExample "Ошибка в процессе выполнения"
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

// @Router /user/get/{id} [get]
// @Summary Получение пользователя
// @Description Позволяет получить информацию о пользователе по его ID
// @Tags users
// @ID get-user
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} entities.User "Данные пользователя"
// @Failure 400 {object} ErrorResponseExample "Ошибка в процессе выполнения"
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

// @Router /users/get [post]
// @Summary Получение списка пользователей по фильтру
// @Description Позволяет получить список пользователей по заданным критериям фильтрации
// @Tags users
// @ID get-users
// @Accept json
// @Produce json
// @Param filter body GetUsersExample true "Фильтры для поиска пользователей"
// @Success 200 {array} entities.User "Список пользователей"
// @Failure 400 {object} ErrorResponseExample "Ошибка в процессе выполнения"
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

// @Router /user/update/{id} [patch]
// @Summary Обновление пользователя
// @Description Позволяет обновить данные пользователя по ID
// @Tags users
// @ID update-user
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param input body UpdateUserExample true "Обновленные данные пользователя"
// @Success 200 {object} entities.User "Обновленные данные пользователя"
// @Failure 400 {object} ErrorResponseExample "Ошибка в процессе выполнения"
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

// @Router /user/delete/{id} [delete]
// @Summary Удаление пользователя
// @Description Позволяет удалить пользователя по его ID
// @Tags users
// @ID delete-user
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} entities.User "Пользователь успешно удалён"
// @Failure 400 {object} ErrorResponseExample "Ошибка в процессе выполнения"
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
