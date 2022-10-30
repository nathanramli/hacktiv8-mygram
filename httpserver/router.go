package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers"
)

type router struct {
	router *gin.Engine

	user *controllers.UserController
}

func NewRouter(r *gin.Engine, user *controllers.UserController) *router {
	return &router{
		router: r,
		user:   user,
	}
}

func (r *router) Start(port string) {
	r.router.POST("/v1/users/register", r.user.Register)
	r.router.POST("/v1/users/login", r.user.Login)
	r.router.Run(port)
}
