package main

import (
	//"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

type SearchHandler struct{}

func (h SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {

	case "/search":

		fmt.Fprintf(w, "welcome to the search endpoint ")

		fmt.Println(r.URL.Path)

		resp, err := http.Get("http://localhost:8080/posts")

		if err != nil {

			fmt.Println(err)
		}

		defer resp.Body.Close()
		ResponseBody, err := (io.ReadAll(resp.Body))
		fmt.Println(string(ResponseBody))

		values := r.URL.Query()
		id := values["id"]

		fmt.Println("the the query parameters values i.e what Query method returns ", values)

		fmt.Println("the ", id[0])
	}

}

func main() {

	godotenv.Load(".env")

	handler := SearchHandler{}

	portString := os.Getenv("PORT")

	http.ListenAndServe(":"+portString, handler)

}
