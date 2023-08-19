package controllers

import (
	"Go-CRUD/initializers"
	"Go-CRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostCreate(c *gin.Context) {

	//Get Data off request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"post": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	//Get Data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the posts were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Respond with it
}

func PostDelete(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Delete
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
