package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Post struct {
	Id          int
	Name        string
	Description string
}

var posts = []Post{

	{Id: 1, Name: "Microservices ", Description: "An in-depth look at why Go's concurrency model is the best in the businness "},
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

	fmt.Fprint(w, "welcome to my website")

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
