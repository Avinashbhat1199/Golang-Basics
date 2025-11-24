package main

import (
	"fmt"

	"github.com/Avinashbhat1199/Golang-Basics/config"
	"github.com/Avinashbhat1199/Golang-Basics/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("tst")
	router := gin.New()
	config.Connect()
	routes.UserRoutes(router)
	router.Run(":8080")
}
