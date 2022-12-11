package api

import (
	"blog/internal/model/request"
	"blog/internal/model/response"
	"blog/internal/service"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	req := request.GetAllCategoryRequest{
		PageNum:  utils.StrToInt(c.DefaultPostForm("pageNum", "1")),
		PageSize: utils.StrToInt(c.DefaultPostForm("pageSize", "10")),
	}
	res, err := service.GetAllCategory(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response.GetAllCategoryResponse{
		Code: 200,
		Data: res,
	})
}
