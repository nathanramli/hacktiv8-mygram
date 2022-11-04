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
	socialMedia *controllers.SocialMediaControllers
}

func NewRouter(r *gin.Engine, user *controllers.UserController, socialMedia *controllers.SocialMediaControllers) *router {
	return &router{
		router:      r,
		user:        user,
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
