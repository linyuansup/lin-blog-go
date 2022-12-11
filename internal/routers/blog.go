package routers

import (
	"blog/internal/api"

	"github.com/gin-gonic/gin"
)

func addBlogRoute(r *gin.RouterGroup) {
	r.POST("/blog/:id", api.GetTargetBlog)
	r.POST("/blog", api.GetBlogList)
}
