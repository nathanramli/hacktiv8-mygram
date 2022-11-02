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

	user *controllers.UserController
	photo *controllers.PhotoController
}

func NewRouter(r *gin.Engine, user *controllers.UserController, photo *controllers.PhotoController) *router {
	return &router{
		router: r,
		user:   user,
		photo: photo,
	}
}

func (r *router) Start(port string) {
	r.router.POST("/v1/users/register", r.user.Register)
	r.router.POST("/v1/users/login", r.user.Login)
	r.router.PUT("/v1/users/:userId", r.verifyToken, r.user.Update)

	r.router.GET("/v1/validate", r.verifyToken, r.user.TestValidate)

	//photo
	r.router.POST("/v1/photos", r.photo.CreatePhoto)
	r.router.GET("/v1/photos", r.photo.GetPhotos)
	r.router.PUT("/v1/photos/:photoId", r.verifyToken, r.photo.UpdatePhoto)

	r.router.Run(port)
}

func (r *router) verifyToken(ctx *gin.Context) {
	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
	// fmt.Println(bearerToken[2])
	if len(bearerToken) != 3 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid bearer token",
		})
		return
	}
	claims, err := common.ValidateToken(bearerToken[2])
	// fmt.Println(bearerToken[1])
	// fmt.Println(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("userData", claims)
	ctx.Set("photoData", claims)
}
