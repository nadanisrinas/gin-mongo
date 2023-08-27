package controllers

import (
	"context"
	"log"
	"net/http"
	"sesi7-challenge/database"
	"sesi7-challenge/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatePosts godoc
// @Summary create post
// @Description create post coresponding on user input
// @Tags posts
// @Accept json
// @Produce json
// @Param models.Post body models.Post true "create car"
// @Success 200 {object} models.Post
// @Router / [post]
func CreatePost(c *gin.Context) {
	var DB = database.ConnectDB()
	var postCollection = GetCollection(DB, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(models.Post)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	postPayload := models.Post{
		ID:      primitive.NewObjectID(),
		Title:   post.Title,
		Article: post.Article,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}

// GetPosts go doc
// @Summary get details
// @Description get details post
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} models.Post
// @Router / [get]
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("myGoappDB").Collection("Posts")
	return collection
}

// GetOnePosts go doc
// @Summary get detail for a given id
// @Description get detail of post coresponding on user input
// @Tags posts
// @Accept json
// @Produce json
// @Param postId path int true "id of post"
// @Success 200 {object} models.Post
// @Router /getOne/{postId} [get]
func ReadOnePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var result models.Post

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}

// Update Posts godoc
// @Summary update post
// @Description update post coresponding on id of post at user input
// @Tags posts
// @Accept json
// @Produce json
// @Param postId path int true "id of post to be updated"
// @Success 200 {object} models.Post
// @Router /update/{postId} [post]
func UpdatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var post models.Post

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"title": post.Title, "article": post.Article}

	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}

// Delete Posts godoc
// @Summary delete post
// @Description delete post coresponding on id of post at user input
// @Tags posts
// @Accept json
// @Produce json
// @Param postId path int true "id of post to be deleted"
// @Success 204 "no content"
// @Router /delete/{postId} [delete]
func DeletePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	postId := c.Param("postId")

	var postCollection = GetCollection(DB, "Posts")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(postId)
	result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": res})
}
