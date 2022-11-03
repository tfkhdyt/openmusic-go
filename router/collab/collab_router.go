package collab

import (
	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/openmusic-go/controller/collab"
	"github.com/tfkhdyt/openmusic-go/middleware/jwt"
)

type Router struct {
	controller *collab.Controller
}

func NewRouter(controller *collab.Controller) *Router {
	return &Router{controller}
}

func (r Router) Route(router *gin.RouterGroup) {
	router.Use(jwt.VerifyJWT())

	router.POST("/", r.controller.Create)
}
