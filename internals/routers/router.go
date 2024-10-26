package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func NewRouterGroup(prefix string) *gin.RouterGroup {
	router := gin.Default()

	sub := router.Group(prefix, nil)

	return sub
}
