package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"context"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var collection *mongo.Collection

func main() {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://saanjeev:saanjeev@cluster0.iqret.mongodb.net/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("test").Collection("items")

	// Set up Gin
	r := gin.Default()
	r.POST("/items", createItem)
	r.GET("/items/:id", getItem)
	r.Run()
}

func createItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	collection.InsertOne(context.TODO(), newItem)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Item created successfully!", "resourceId": newItem.ID})
}

func getItem(c *gin.Context) {
	itemID := c.Param("id")

	var item Item
	err := collection.FindOne(context.TODO(), bson.M{"id": itemID}).Decode(&item)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No item found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": item})
}