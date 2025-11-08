package main

import (
	"database/sql"
	"encoding/json"

	"strings"
	"strconv"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	//
	// "strconv"
	// "strings"
)

type Post struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MyHandler struct {
}

var db *sql.DB
var sqlError error

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	print(path)

	switch {
	case path == "/":
		fmt.Fprint(w, "homepage")

	case path == "/posts":

		var id int
		var name, description string

		var posts []Post

		rows, err := db.Query(`SELECT * FROM posts`)
		defer rows.Close()

		if err != nil {

			fmt.Println(err, "error from the select statement ")
		}

		for rows.Next() {

			rows.Scan(&id, &name, &description)
			newPost := Post{
				Id:          id,
				Name:        name,
				Description: description,
			}
			posts = append(posts, newPost)
		}
		
		data, _ := json.MarshalIndent(posts, "", " ")

		w.Header().Set("Content-Type", "application/json ")
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
				w.Header().Set("Content-Type", "text/html  ")
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

	// username:password@protocol(address)/dbname?param=value

	dsn := "root:@tcp(127.0.0.1:3306)/blogtest?parseTime=true"

	// dsn := "root:@tcp(127.0.0.1:3306)/blog?parseTime=true"

	db, sqlError = sql.Open("mysql", dsn)

	result, syntaxErro := db.Exec(

		` CREATE TABLE IF NOT EXISTS posts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50),
      
		description VARCHAR(250)
    )`,
	)

	if syntaxErro != nil {

		fmt.Println("the error from syntax is ")

		fmt.Println(syntaxErro)
	}
	fmt.Println(result)

	if sqlError != nil {

		fmt.Println("the error with the db connection pool is")

		fmt.Println(sqlError)
	} else {

		fmt.Println("no error")
	}

	defer db.Close()

	handler := MyHandler{}

	fmt.Printf("Server started at %v \n", os.Getenv("PORT"))

	err := http.ListenAndServe(":"+portAdd, handler)

	if err != nil {

		fmt.Println(err.Error())

	}

}
