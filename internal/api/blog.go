package api

import (
	"blog/internal/model/request"
	"blog/internal/model/response"
	"blog/internal/service"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlogList(c *gin.Context) {
	req := request.GetBlogListRequest{
		CategoryName: c.DefaultPostForm("categoryName", ""),
		PageNum:      utils.StrToInt(c.PostForm("pageNum")),
		PageSize:     utils.StrToInt(c.PostForm("pageSize")),
	}
	res, err := service.GetBlogList(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response.GetBlogListResponse{
		Code: 200,
		Data: res,
	})
}

func GetTargetBlog(c *gin.Context) {
	req := request.GetTargetBlogRequest{
		ID: c.Param("id"),
	}
	res, err := service.GetTargetBlog(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response.GetTargetBlogResponse{
		Code: 200,
		Data: res,
	})
}
