package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/auth"
)

type Router struct {
	controller *auth.Controller
}

func NewRouter(controller *auth.Controller) *Router {
	return &Router{controller}
}

func (r Router) Route(router *gin.RouterGroup) {
	router.POST("/", r.controller.Login)
	router.PUT("/", r.controller.RefreshToken)
}
