package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/msssrp/go-learn/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	collection *mongo.Collection
}

func NewUserCollection(c *mongo.Client) *UserController {
	return &UserController{
		collection: models.NewUserCollection(c),
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := user.InsertUser(uc.collection.Database().Client())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, res)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {

	userId := c.Param("id")
	if userId != "" {
		id, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		user := models.User{}
		filter := bson.M{"_id": id}
		err = uc.collection.FindOne(context.Background(), filter).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	} else {
		users := []models.User{}

		cursor, err := uc.collection.Find(context.Background(), bson.M{})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := cursor.All(context.Background(), &users); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}

}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}
	var user models.User
	res, err := uc.collection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error deleting"})
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user %s deleted success", user.FirstName)})
}
