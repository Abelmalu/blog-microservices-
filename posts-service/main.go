package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"strings"
)

type Post struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var posts = []Post{

	{
		Id:          1,
		Name:        "Microservices",
		Description: "An in-depth look at why Go's concurrency model is the best in the businness ",
	},
	{
		Id:          2,
		Name:        "Mastering Net/HTTP: No Frameworks Needed",
		Description: "A tutorial on building a high-performance HTTP server using only the Go Standard Library, focusing on the http.Handler interface and ServeMux.",
	},
	{
		Id:          3,
		Name:        "Optimizing PostgreSQL Queries for Scale",
		Description: "Best practices for indexing, query planning, and connection pooling to ensure your application's database performance scales with your traffic.",
	},
	{
		Id:          4,
		Name:        "Understanding Message Brokers (Kafka vs. RabbitMQ)",
		Description: "A comparison of the two leading message brokers and a guide on choosing the right one for asynchronous task processing in your service-oriented architecture.",
	},
}

type MyHandler struct {
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case path == "/":
		fmt.Fprint(w, "homepage")

	case path == "/posts":
		// Return all posts
		data, _ := json.MarshalIndent(posts, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	case strings.HasPrefix(path, "/posts/"):
		// Handle /posts/{id}
		id := strings.TrimPrefix(path, "/posts/")
		// Parse the id to an integer
		pid, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid post id", http.StatusBadRequest)
			return
		}
		// Now find the post with this id
		for _, post := range posts {
			if post.Id == pid {
				data, _ := json.MarshalIndent(post, "", "  ")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(data)
				return
			}
		}
		http.Error(w, "Post not found", http.StatusNotFound)
		

	default:
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}

func main() {

	godotenv.Load(".env")
	portAdd := os.Getenv("PORT")

	handler := MyHandler{}

	fmt.Printf("Server started at %v \n", os.Getenv("PORT"))

	err := http.ListenAndServe(":"+portAdd, handler)

	if err != nil {

		fmt.Println(err.Error())

	}

}
