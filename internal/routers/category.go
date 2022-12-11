package routers

import (
	"blog/internal/api"

	"github.com/gin-gonic/gin"
)

func addCategoryRoute(r *gin.RouterGroup) {
	r.POST("/category", api.GetAllCategory)
}
