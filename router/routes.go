package router

import (
	"github.com/FelipeScherem/goapiestudo.git/handlers/handlerOpening"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opening", handlerOpening.ShowOpeningHandler)
		v1.POST("/opening", handlerOpening.CreateOpeningHandler)
		v1.DELETE("/opening", handlerOpening.DeleteOpeningHandler)
		v1.PUT("/opening", handlerOpening.UpdateOpeningHandler)
		v1.GET("/openings", handlerOpening.ListOpeningsHandler)
	}
}
