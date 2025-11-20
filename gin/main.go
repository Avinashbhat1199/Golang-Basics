package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	routes.userroute(router)
	router.Run(":8080")
}
