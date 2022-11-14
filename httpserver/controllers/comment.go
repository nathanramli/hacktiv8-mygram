package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	// "github.com/dgrijalva/jwt-go"
	"net/http"

	"github.com/nathanramli/hacktiv8-mygram/common"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/services"
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
	userId := userData.Id
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
	response := c.svc.GetComments(ctx)
	WriteJsonResponse(ctx, response)
}

func (c *CommentController) EditComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req params.UpdateComment
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	comment, err := c.svc.GetCommentByID(ctx, id)
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
	if userData.Id != comment.UserId {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	response := c.svc.EditComments(ctx, &req, uint(id))
	WriteJsonResponse(ctx, response)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	comment, err := c.svc.GetCommentByID(ctx, id)
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
	if userData.Id != comment.UserId {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	response := c.svc.DeleteComment(ctx, uint(id))
	WriteJsonResponse(ctx, response)
}
