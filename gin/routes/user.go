package routes

import "github.com/gin-gonic/gin"

func userroute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})
}
