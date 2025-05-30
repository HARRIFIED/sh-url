package handler

import (
	"net/http"

	"url-shortner-go/api/v1/dto"
	"url-shortner-go/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	svc *service.URLService
}

func NewURLHandler(svc *service.URLService) *URLHandler {
	return &URLHandler{svc: svc}
}

func (h *URLHandler) Register(rg *gin.RouterGroup) {
	urls := rg.Group("/urls")
	urls.POST("/shorten", h.Shorten)
	urls.GET("/:code", h.Resolve)
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req dto.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.svc.Shorten(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.ShortenResponse{Code: u.Code})
}

func (h *URLHandler) Resolve(c *gin.Context) {
	code := c.Param("code")
	u, err := h.svc.Resolve(code)
	if err != nil || u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Redirect(http.StatusFound, u.Original)
}
