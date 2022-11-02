package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nathanramli/hacktiv8-mygram/common"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
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

func (c *SocialMediaControllers) UpdateSocialMedia(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req params.UpdateSocialMedia
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	claims, exists := ctx.Get("socialMediaData")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}
	socialMediaData := claims.(*common.CustomClaims)
	if socialMediaData.Id != id {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}
	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{Message: "OK"})
}