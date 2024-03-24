package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lrnholanda/gopportinities/handler"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"
	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
