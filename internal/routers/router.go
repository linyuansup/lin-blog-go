package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	addBlogRoute(api)
	addCategoryRoute(api)
	r.Static("/static", "./static")
	return r
}
