package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicialize Router
	r := gin.Default()

	// Initialize Routes
	initializeRoutes(r)

	// Run the server
	r.Run(":8080")
}
