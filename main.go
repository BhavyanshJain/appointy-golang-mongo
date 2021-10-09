package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	controllers "github.com/BhavyanshJain/appointy-golang-mongo/controllers"
)

func main() {
	fmt.Println("Starting the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(rw http.ResponseWriter, r *http.Request) { controllers.CreateUserEndpoint(rw, r, client) })
	mux.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) { controllers.CreatePostEndpoint(rw, r, client) })
	mux.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) { controllers.GetUsersEndpoint(rw, r, client) })
	mux.HandleFunc("/posts", func(rw http.ResponseWriter, r *http.Request) { controllers.GetPostsEndpoint(rw, r, client) })
	mux.HandleFunc("/users/", func(rw http.ResponseWriter, r *http.Request) { controllers.GetUserEndpoint(rw, r, client) })
	mux.HandleFunc("/posts/users/", func(rw http.ResponseWriter, r *http.Request) { controllers.GetPostsOfUserEndpoint(rw, r, client) })
	log.Fatal(http.ListenAndServe(":12345", mux))

}
