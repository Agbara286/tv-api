package main

import (
	"state-tv-api/config"
	"state-tv-api/controllers"
	"state-tv-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize Database
	config.ConnectDB()

	// 2. AutoMigrate the Database Tables
	config.DB.AutoMigrate(&models.Article{})

	// 3. Setup Router
	router := gin.Default()

	// 4. Setup CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 5. Register Routes
	router.GET("/api/articles", controllers.GetArticles)
	router.POST("/api/articles", controllers.CreateArticle)
	router.POST("/api/subscribe", controllers.SubscribeToNewsletter)
	router.GET("/api/articles/:id", controllers.GetArticle)
	router.PUT("/api/articles/:id", controllers.UpdateArticle)
	router.DELETE("/api/articles/:id", controllers.DeleteArticle)

	// 6. Start Server
	router.Run(":8080")
}
