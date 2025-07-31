package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/service"
)

type FirmHandler struct {
	Svc *service.FirmService
}

func NewFirmHandler(svc *service.FirmService) *FirmHandler {
	return &FirmHandler{Svc: svc}
}

func (h *FirmHandler) List(c *gin.Context) {
	firms, err := h.Svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, firms)
}

func (h *FirmHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid firm ID"})
		return
	}

	firm, err := h.Svc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, firm)
}

func (h *FirmHandler) Create(c *gin.Context) {
	var firm models.Firm
	if err := c.ShouldBindJSON(&firm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := h.Svc.Create(c.Request.Context(), firm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *FirmHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid firm ID"})
		return
	}

	var firm models.Firm
	if err := c.ShouldBindJSON(&firm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	firm.ID = id
	if err := h.Svc.Update(c.Request.Context(), firm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Firm updated successfully"})
}

func (h *FirmHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid firm ID"})
		return
	}

	if err := h.Svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Firm deleted successfully"})
}

func (h *FirmHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/firms")
	r.GET("", h.List)
	r.GET("/:id", h.Get)
	r.POST("", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
