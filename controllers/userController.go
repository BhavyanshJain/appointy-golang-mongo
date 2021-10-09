package controllers

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	models "github.com/BhavyanshJain/appointy-golang-mongo/models"
)

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request, client *mongo.Client) {
	response.Header().Set("content-type", "application/json")
	var user models.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("appointy-tech-task").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: md5.Sum([]byte(user.Password))},
	})
	defer cancel()
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	json.NewEncoder(response).Encode(result)
}

func GetUserEndpoint(response http.ResponseWriter, request *http.Request, client *mongo.Client) {
	response.Header().Set("content-type", "application/json")
	idi := strings.TrimPrefix(request.URL.Path, "/users/")
	id, _ := primitive.ObjectIDFromHex(idi)
	var user models.User
	collection := client.Database("appointy-tech-task").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, models.User{Id: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

func GetUsersEndpoint(response http.ResponseWriter, request *http.Request, client *mongo.Client) {
	if request.URL.Path != "/users" {
		http.NotFound(response, request)
		return
	}
	response.Header().Set("content-type", "application/json")
	var users []models.User
	collection := client.Database("appointy-tech-task").Collection("users")
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
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}
