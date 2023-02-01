package handlers

import (
	"github.com/gin-gonic/gin"
	"go-commerce/dtos/request"
	"go-commerce/dtos/response"
	"go-commerce/services"
	"net/http"
)

type authenticationHandler struct {
	authenticationService services.AuthenticationService
}

func NewAuthenticationHandler(authenticationService services.AuthenticationService) *authenticationHandler {
	return &authenticationHandler{authenticationService: authenticationService}
}

func (h *authenticationHandler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.DefaultResponse{
			Status:  400,
			Message: "Failed to register",
			Data:    nil,
		})
		return
	}

	user, err := h.authenticationService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.DefaultResponse{
			Status:  500,
			Message: "Failed to register",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.DefaultResponse{
		Status:  200,
		Message: "Success",
		Data:    user,
	})
}

func (h *authenticationHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.DefaultResponse{
			Status:  400,
			Message: "Failed to register",
			Data:    nil,
		})
		return
	}

	token, err := h.authenticationService.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.DefaultResponse{
			Status:  500,
			Message: "Failed to login",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.DefaultResponse{
		Status:  200,
		Message: "Success",
		Data:    token,
	})
}
