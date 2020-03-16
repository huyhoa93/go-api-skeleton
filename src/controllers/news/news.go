package news

import (
	"net/http"
	"strconv"

	newsModel "../../models/news"
	news "../../services/news"
	"github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
	res := news.GetNews()
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
		"data":    res.Data,
	})
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
