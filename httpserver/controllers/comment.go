package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// "github.com/dgrijalva/jwt-go"
	"github.com/nathanramli/hacktiv8-mygram/common"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/services"
	"net/http"
	// "strconv"
)

type CommentController struct {
	svc services.CommentSvc
}

func NewCommentController(svc services.CommentSvc) *CommentController {
	return &CommentController{
		svc: svc,
	}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var req params.CreateComment

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
	response := c.svc.CreateComment(ctx, &req, userId)
	WriteJsonResponse(ctx, response)
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	response:= c.svc.GetComments(ctx)
	WriteJsonResponse(ctx, response)
}