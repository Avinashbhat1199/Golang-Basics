package controller

import "github.com/gin-gonic/gin"

func UserContoller(ctx *gin.Context) {

	ctx.String(200, "test")
}
