package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	r := gin.Default()

	// Initialize Router
	initializeRoutes(r)

	r.Run()
}
