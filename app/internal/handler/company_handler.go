package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/youruser/order-app/internal/models"
	"github.com/youruser/order-app/internal/service"
)

type CompanyHandler struct {
	Svc *service.CompanyService
}

func NewCompanyHandler(svc *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{Svc: svc}
}

func (h *CompanyHandler) List(c *gin.Context) {
	companies, err := h.Svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *CompanyHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	company, err := h.Svc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) Create(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := h.Svc.Create(c.Request.Context(), company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *CompanyHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	company.ID = id

	if err := h.Svc.Update(c.Request.Context(), company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}

func (h *CompanyHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	if err := h.Svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

func (h *CompanyHandler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/companies")
	r.GET("/", h.List)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
