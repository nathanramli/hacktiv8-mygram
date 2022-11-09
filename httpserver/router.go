package httpserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nathanramli/hacktiv8-mygram/common"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers"
)

type router struct {
	router *gin.Engine

	user        *controllers.UserController
	photo       *controllers.PhotoController
	comment     *controllers.CommentController
	socialMedia *controllers.SocialMediaControllers
}

func NewRouter(r *gin.Engine, user *controllers.UserController, photo *controllers.PhotoController, comment *controllers.CommentController, socialMedia *controllers.SocialMediaControllers) *router {
	return &router{
		router:      r,
		user:        user,
		photo:       photo,
		comment:     comment,
		socialMedia: socialMedia,
	}
}

func (r *router) Start(port string) {
	r.router.POST("/v1/users/register", r.user.Register)
	r.router.POST("/v1/users/login", r.user.Login)
	r.router.PUT("/v1/users/:userId", r.verifyToken, r.user.Update)
	r.router.DELETE("/v1/users", r.verifyToken, r.user.Delete)

	// Test
	r.router.GET("/v1/validate", r.verifyToken, r.user.TestValidate)

	// social media
	r.router.POST("/v1/socialmedias", r.verifyToken, r.socialMedia.CreateSocialMedia)
	r.router.GET("/v1/socialmedias", r.verifyToken, r.socialMedia.GetSocialMedia)
	r.router.PUT("/v1/socialmedias/:socialMediaId", r.verifyToken, r.socialMedia.UpdateSocialMedia)
	r.router.DELETE("/v1/socialmedias/:socialMediaId", r.verifyToken, r.socialMedia.DeleteSocialMedia)

	//photo
	r.router.POST("/v1/photos", r.verifyToken, r.photo.CreatePhoto)
	r.router.GET("/v1/photos", r.verifyToken, r.photo.GetPhotos)
	r.router.PUT("/v1/photos/:photoId", r.verifyToken, r.photo.UpdatePhoto)
	r.router.DELETE("/v1/photos/:photoId", r.verifyToken, r.photo.DeletePhoto)

	//comment
	r.router.POST("/v1/comments", r.verifyToken, r.comment.CreateComment)
	r.router.GET("/v1/comments", r.verifyToken, r.comment.GetComments)
	r.router.PUT("/v1/comments/:commentId", r.verifyToken, r.comment.EditComment)

	r.router.Run(port)
}

func (r *router) verifyToken(ctx *gin.Context) {
	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
	if len(bearerToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid bearer token",
		})
		return
	}
	claims, err := common.ValidateToken(bearerToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("userData", claims)
	// ctx.Set("socialMediaData", claims)
}
