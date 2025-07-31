package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/service"
)

type OrderHandler struct {
	Svc *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{Svc: svc}
}

func (h *OrderHandler) List(c *gin.Context) {
	orders, err := h.Svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.Svc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Create(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := h.Svc.Create(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *OrderHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	order.ID = id
	if err := h.Svc.Update(c.Request.Context(), order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})

}

func (h *OrderHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := h.Svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func (h *OrderHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/orders")
	r.GET("", h.List)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
