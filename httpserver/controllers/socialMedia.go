package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/services"
)

type SocialMediaControllers struct {
	svc services.SocialMediaSvc
}

func NewSocialMediaController(svc services.SocialMediaSvc) *SocialMediaControllers {
	return &SocialMediaControllers{
		svc: svc,
	}
}

func (c *SocialMediaControllers) CreateSocialMedia(ctx *gin.Context) {
	var request params.CreateSocialMedia
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := validator.New().Struct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := c.svc.CreateSocialMedia(ctx, &request)
	WriteJsonResponse(ctx, response)
}
