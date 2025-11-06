package main

import (
	"fmt"
	"net/http"
	//"os"
)

func main() {

	resp, err := http.Get("http://localhost:8080/posts")

	if err != nil {

		fmt.Println(err)
	}
	defer resp.Body.Close()

	buffer := make([]byte, 1000)

	n, err := resp.Body.Read(buffer)
	length := resp.ContentLength

	fmt.Println(length)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(string(buffer), n)

}
