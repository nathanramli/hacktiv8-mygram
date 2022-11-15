package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/dgrijalva/jwt-go"
	"github.com/nathanramli/hacktiv8-mygram/common"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/services"
	"net/http"
	"strconv"
)

type PhotoController struct {
	svc services.PhotoSvc
}

func NewPhotoController(svc services.PhotoSvc) *PhotoController {
	return &PhotoController{
		svc: svc,
	}
}

func (c *PhotoController) CreatePhoto(ctx *gin.Context) {
	var req params.CreatePhoto

	if err := ctx.ShouldBindJSON(&req); err != nil {
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
	userId := int(userData.Id)
	err := validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := c.svc.CreatePhoto(ctx, &req, userId)
	WriteJsonResponse(ctx, response)
}

func (c *PhotoController) GetPhotos(ctx *gin.Context) {
	response:= c.svc.GetPhotos(ctx)
	WriteJsonResponse(ctx, response)
}

func (c *PhotoController) UpdatePhoto(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req params.UpdatePhoto
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	photo, err := c.svc.GetPhotoByID(ctx, id)

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

	userData := claims.(*common.CustomClaims)
	if userData.Id != photo.UserId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized User",
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

	response := c.svc.UpdatePhoto(ctx, &req, id)
	
	WriteJsonResponse(ctx, response)
}

func (c *PhotoController) DeletePhoto(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	photo, err := c.svc.GetPhotoByID(ctx, id)

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

	userData := claims.(*common.CustomClaims)
	if userData.Id != photo.UserId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized User",
			})
		return
	}
	response := c.svc.DeletePhoto(ctx, id)
	
	WriteJsonResponse(ctx, response)
}