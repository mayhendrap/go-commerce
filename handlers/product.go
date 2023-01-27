package handlers

import (
	"github.com/gin-gonic/gin"
	"go-commerce/dtos/request"
	"go-commerce/dtos/response"
	"go-commerce/helpers"
	"go-commerce/services"
	"net/http"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{productService: productService}
}

func (h *productHandler) Create(c *gin.Context) {
	var input request.ProductRequest
	if c.ShouldBindJSON(&input) != nil {

	}
}

func (h *productHandler) FindAll(c *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ResponseInternalServerError)
	}
	rsp := response.DefaultResponse{
		Status:  200,
		Message: "Success",
		Data:    products,
	}
	c.JSON(http.StatusOK, rsp)
}
