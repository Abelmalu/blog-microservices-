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

	resp, err := http.Get("http://localhost:8080/posts")

	if err != nil {

		http.Error(w, "failed to fetch posts", http.StatusInternalServerError)
	}

	defer resp.Body.Close()
	// status := resp.Status
	// fmt.Println(resp.Body.Read())
	// io.Copy(os.Stdout, resp.Body)
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	_, err = io.Copy(os.Stdout, resp.Body)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error copying: %v\n", err)
	// 	os.Exit(1)

	// }
}

func main() {

	godotenv.Load(".env")

	handler := SearchHandler{}

	portString := os.Getenv("PORT")

	http.ListenAndServe(":"+portString, handler)

}
