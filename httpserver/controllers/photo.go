package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/nathanramli/hacktiv8-mygram/common"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	// "github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/services"
	"net/http"
	// "strconv"
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

	err := validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := c.svc.CreatePhoto(ctx, &req)
	WriteJsonResponse(ctx, response)
}

func (c *PhotoController) GetPhotos(ctx *gin.Context) {
	response:= c.svc.GetPhotos(ctx)
	WriteJsonResponse(ctx, response)
}