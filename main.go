package main

import (
	"log"
	"net/http"
	"os"

	auth "./src/controllers/auth"
	news "./src/controllers/news"
	utility "./src/utility"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		check := utility.ValidateToken(c)
		if check {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			return
		}
	}
}

func setupRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/public", "./public")

	authen := r.Group("/auth")
	{
		authen.POST("/login", auth.Login)
	}

	api := r.Group("/api")
	api.Use(AuthMiddleware())
	{
		api.GET("/news", news.GetNews)
		api.GET("/news/:id", news.GetNewsById)
		api.POST("/news/add", news.AddNews)
		api.PUT("/news/update/:id", news.UpdateNews)
		api.DELETE("/news/delete/:id", news.DeleteNews)
	}

	return r
}

func main() {
	//Read the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the "PORT" env variable
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := setupRouter()
	router.Run(":" + port)
}
