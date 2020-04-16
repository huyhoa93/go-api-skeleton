package news

import (
	"net/http"
	"strconv"

	"gopkg.in/validator.v2"

	commentsModel "../../models/comments"
	common "../../models/common"
	newsModel "../../models/news"
	news "../../services/news"
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
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	if invalid := validator.Validate(data); invalid != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": invalid.Error(),
		})
		return
	}
	res := news.AddNews(data.Title, data.Content)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
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

func AddComment(c *gin.Context) {
	var data commentsModel.CommentData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	if invalid := validator.Validate(data); invalid != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": invalid.Error(),
		})
		return
	}
	newsRes := news.GetNewsById(data.NewsId)
	if newsRes.Status != 200 {
		c.JSON(newsRes.Status, gin.H{
			"status":  newsRes.Status,
			"message": newsRes.Message,
		})
		return
	}
	res := news.AddComment(data.NewsId, data.Comment)
	c.JSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
	})
}
