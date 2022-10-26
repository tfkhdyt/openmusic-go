package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/user"
)

type Router struct {
	controller *user.Controller
}

func NewRouter(controller *user.Controller) *Router {
	return &Router{controller}
}

func (r Router) Route(router *gin.RouterGroup) {
	router.POST("/", r.controller.Create)
}
