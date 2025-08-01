package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/service"
)

type ProductCityHandler struct {
	Svc *service.ProductCityService
}

func NewProductCityHandler(svc *service.ProductCityService) *ProductCityHandler {
	return &ProductCityHandler{Svc: svc}
}

func (h *ProductCityHandler) List(c *gin.Context) {
	productCities, err := h.Svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productCities)
}

func (h *ProductCityHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product city ID"})
		return
	}

	productCity, err := h.Svc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productCity)
}

func (h *ProductCityHandler) Create(c *gin.Context) {
	var productCity models.ProductCity
	if err := c.ShouldBindJSON(&productCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := h.Svc.Create(c.Request.Context(), productCity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *ProductCityHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product city ID"})
		return
	}

	var productCity models.ProductCity
	if err := c.ShouldBindJSON(&productCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	productCity.ID = id
	if err := h.Svc.Update(c.Request.Context(), productCity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product city updated successfully"})
}

func (h *ProductCityHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product city ID"})
		return
	}

	if err := h.Svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product city deleted successfully"})
}

func (h *ProductCityHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/product-cities")
	r.GET("", h.List)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
