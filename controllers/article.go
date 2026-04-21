package controllers

import (
	"state-tv-api/config"
	"state-tv-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticles handles GET /api/articles
func GetArticles(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	var articles []models.Article 
	var total int64

	config.DB.Model(&models.Article{}).Count(&total)
	config.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&articles)

	c.JSON(200, gin.H{"data": articles, "meta": gin.H{"total": total, "page": page, "limit": limit}})
}

// CreateArticle handles POST /api/articles
func CreateArticle(c *gin.Context) {
	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(201, input)
}

// UpdateArticle handles PUT /api/articles/:id
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}

	var updateData models.Article
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	config.DB.Model(&article).Updates(updateData)
	c.JSON(200, article)
}

// DeleteArticle handles DELETE /api/articles/:id
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}

	config.DB.Delete(&article)
	c.JSON(200, gin.H{"message": "Article deleted successfully"})
}
// GetArticle handles GET /api/articles/:id
func GetArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(200, article)
}
