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

	claims, exist := ctx.Get("userData")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}
	userData := claims.(*common.CustomClaims)
	userId := uint(userData.Id)
	err := validator.New().Struct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := c.svc.CreateSocialMedia(ctx, &request, userId)
	WriteJsonResponse(ctx, response)
}
func (c *SocialMediaControllers) GetSocialMedia(ctx *gin.Context) {
	response := c.svc.GetSocialMedia(ctx)
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
	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims, exists := ctx.Get("userData")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}

	if socialMediaData := claims.(*common.CustomClaims); socialMediaData.Id != id {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized to update other socialmedia data",
		})
		return
	}

	if err = validator.New().Struct(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp := c.svc.UpdateSocialMedia(ctx, &req, uint(id))
	WriteJsonResponse(ctx, resp)
}

func (c *SocialMediaControllers) DeleteSocialMedia(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims, exists := ctx.Get("userData")
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

	_, err = c.svc.DeleteSocialMedia(ctx, uint(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{Message: "Your Social media has been successfully deleted"})
}
