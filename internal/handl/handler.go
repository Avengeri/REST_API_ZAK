package handl

import (
	"Interface_droch_3/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/user")
	{
		auth.POST("/", h.CreateUser)
		auth.GET("/:id", h.GetUser)
		auth.GET("/check/:id", h.CheckUser)
		auth.DELETE("/:id", h.DeleteUser)
		auth.GET("/get_all", h.GetAllUsers)
	}
	return router
}
