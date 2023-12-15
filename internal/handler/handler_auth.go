package handler

import (
	"Interface_droch_3/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

type UserResponse struct {
	Message string `json:"message"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided JSON data
// @Tags user
// @Param user body model.User true "User data in JSON format"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный запрос (Bad Request)"})
		return
	}

	err := h.service.Set(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Внутренняя ошибка сервера (Internal Server Error)"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{Message: "Пользователь успешно создан"})
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user with the provided JSON data
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @Success 200 {string} string "User get successfully"
// @Failure 500 {object} ErrorResponse
// @Router /user/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)

	user, err := h.service.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Внутренняя ошибка сервера (Internal Server Error)"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CheckUser godoc
// @Summary Check if a user exists
// @Description Check if a user with the provided ID exists
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/check/{id} [get]
func (h *Handler) CheckUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный запрос (Bad Request)"})
		return
	}

	// Вызывайте метод Check вашего сервиса для проверки существования пользователя
	exists, err := h.service.Check(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Внутренняя ошибка сервера (Internal Server Error)"})
		return
	}

	if exists {
		c.JSON(http.StatusOK, UserResponse{Message: "Пользователь успешно найден"})
	} else {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Пользователь не найден"})
	}
}

// DeleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user with the provided ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID" format(int64)
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный запрос (Bad Request)"})
		return
	}

	err = h.service.Delete(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Внутренняя ошибка сервера (Internal Server Error)"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{Message: "Пользователь успешно удален"})
}

// GetAllUsers godoc
// @Summary Get a list of all users
// @Description Get a list of all users with their IDs
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} int "List of user IDs"
// @Router /user/get_all [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	// Получите список всех пользователей с помощью метода GetAllId вашего сервиса
	userIDs := h.service.GetAllId()

	// Дополнительная логика для получения данных о пользователях на основе их ID

	// Верните список пользователей в формате JSON
	c.JSON(http.StatusOK, userIDs)
}
