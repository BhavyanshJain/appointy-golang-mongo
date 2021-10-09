package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	models "github.com/BhavyanshJain/appointy-golang-mongo/models"
)

func CreatePostEndpoint(response http.ResponseWriter, request *http.Request, client *mongo.Client) {
	response.Header().Set("content-type", "application/json")
	var post models.Post
	_ = json.NewDecoder(request.Body).Decode(&post)
	collection := client.Database("appointy-tech-task").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, bson.D{
		{Key: "userId", Value: post.UserId},
		{Key: "caption", Value: post.Caption},
		{Key: "imageURL", Value: post.ImageURL},
		{Key: "postedTimeStamp", Value: time.Now().Format(time.RFC3339)},
	})

	defer cancel()
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	json.NewEncoder(response).Encode(result)

}

func GetPostsEndpoint(response http.ResponseWriter, request *http.Request, client *mongo.Client) {
	if request.URL.Path != "/posts" {
		http.NotFound(response, request)
		return
	}
	response.Header().Set("content-type", "application/json")
	var posts []models.Post
	collection := client.Database("appointy-tech-task").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(posts)
}

func GetPostsOfUserEndpoint(response http.ResponseWriter, request *http.Request, client *mongo.Client) {
	response.Header().Set("content-type", "application/json")
	id := strings.TrimPrefix(request.URL.Path, "/posts/users/")
	var posts []models.Post
	collection := client.Database("appointy-tech-task").Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.D{{Key: "userId", Value: id}}
	cursor, err := collection.Find(ctx, filter, options.Find())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(posts)
}
