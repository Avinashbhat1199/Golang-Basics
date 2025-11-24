package routes

import (
	"github.com/Avinashbhat1199/Golang-Basics/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.UserContoller)
}
