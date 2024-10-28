package controllers

import (
	"tardis/internals/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterStoreRoutes(r *gin.Engine) {
	routergroup := r.Group("api/v1/", nil)

	//register Routes
	r.GET("/", handlers.Ping)
	r.POST("/", handlers.Create)
	routergroup.POST("", handlers.Retrieve)
	routergroup.DELETE("", handlers.Delete)
	routergroup.PATCH("", handlers.Update)
}
