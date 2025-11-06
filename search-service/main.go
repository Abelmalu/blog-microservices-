package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	// "reflect"

	"github.com/joho/godotenv"
)

type SearchHandler struct{}

func (h SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hellow")

	switch r.URL.Path {

	case "/search":

		// fmt.Fprintf(w, "welcome to the search endpoint ")

		fmt.Println(r.URL.Path)

		resp, err := http.Get("http://localhost:8080/posts")

		if err != nil {

			fmt.Println(err)
		}

		defer resp.Body.Close()
		ResponseBody, _ := io.ReadAll(resp.Body)
		type Posts struct {
			Id          int
			Name        string
			Description string
		}

		var result []Posts

		err = json.Unmarshal(ResponseBody, &result)

		// fmt.Println(result, "this is the result in golang map")
		// fmt.Println(reflect.TypeOf(result[0].Id))

		if err != nil {

			fmt.Println(err)
		}

		values := r.URL.Query()
		QueryId := values["id"][0]

		for _, value := range result {

		
			id := strconv.Itoa(value.Id)
			//for some reason string(value.id) didn't work CHECK IT OUT

			if id == QueryId {
				fmt.Println(id)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				JsonData, _:= json.Marshal(value)
				w.Write(JsonData)
				

			} else{

				fmt.Println("come on")


			}

			

		}

	}

}

func main() {

	godotenv.Load(".env")

	handler := SearchHandler{}

	portString := os.Getenv("PORT")

	http.ListenAndServe(":"+portString, handler)

}
