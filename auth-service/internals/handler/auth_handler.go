package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"auth-service/internals/domain"
	"auth-service/internals/service"
)

type AuthHandler struct {
	userService *service.AuthService
}

func NewAuthHandler(userService *service.AuthService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) SetupRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		log.Println("Seting up routes")
		authGroup.GET("/alive", h.Alive)
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
	}
}

func (h *AuthHandler) Alive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "reached"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request domain.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.userService.RegisterUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request domain.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.userService.LoginUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "loggin succesfully",
		"token":   token,
	})
	return

}
