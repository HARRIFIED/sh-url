package handler

import (
	"net/http"

	"url-shortner-go/api/v1/dto"
	"url-shortner-go/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	urlService *service.URLService
}

func NewURLHandler(urlService *service.URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

func (handler *URLHandler) Shorten(context *gin.Context) {
	var req dto.ShortenRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortenUrl, err := handler.urlService.Shorten(req.URL)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, dto.ShortenResponse{Code: shortenUrl.Code})
}

func (handler *URLHandler) Resolve(context *gin.Context) {
	code := context.Param("code")
	url, err := handler.urlService.Resolve(code)
	if err != nil || url == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	context.Redirect(http.StatusFound, url.Original)
}
