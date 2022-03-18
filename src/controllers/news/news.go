package news

import (
	"net/http"
	"strconv"

	common "go_api/src/models/common"
	newsModel "go_api/src/models/news"
	news "go_api/src/services/news"

	"github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
	var param common.PaginationParams
	if err := c.ShouldBindQuery(&param); err == nil {
		res := news.GetNews(param.Page, param.Perpage)
		if res.Status != 200 {
			c.JSON(res.Status, gin.H{
				"status":  res.Status,
				"message": res.Message,
			})
			return
		}
		c.JSON(res.Status, gin.H{
			"status":  res.Status,
			"message": res.Message,
			"data":    res.Data,
			"total":   res.Total,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

}

func AddNews(c *gin.Context) {
	var data newsModel.Post
	if err := c.ShouldBindJSON(&data); err == nil {
		res := news.AddNews(data.Title, data.Content)
		c.JSON(res.Status, gin.H{
			"status":  res.Status,
			"message": res.Message,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}

func GetNewsById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := news.GetNewsById(id)
	if res.Status != 200 {
		c.JSON(res.Status, gin.H{
			"status":  res.Status,
			"message": res.Message,
		})
		return
	}
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
		"data":    res.Data,
	})
}

func UpdateNews(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data newsModel.Post
	if err := c.ShouldBindJSON(&data); err == nil {
		res := news.UpdateNews(id, data.Title, data.Content)
		c.JSON(res.Status, gin.H{
			"status":  res.Status,
			"message": res.Message,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
}

func DeleteNews(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := news.DeleteNews(id)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
}
