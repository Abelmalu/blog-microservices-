package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 100)
	fmt.Println("the length of the buffer is ", len(buf))
	for {
		n, err := file.Read(buf)
		newFile, _ := io.ReadAll(file)
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			panic(err)
		}

		fmt.Print(string(buf[:n]))
	}
}
