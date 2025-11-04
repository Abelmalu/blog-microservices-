package main

import (
	"fmt"
	"encoding/json"


)

type Student struct {
	ID         string `json:"id"`
	FullName   string `json:"full_name"`
	Batch      string `json:"batch"`
	Department string `json:"department"`
}

// func main() {
// 	student := Student{
// 		ID:         "C6:65:D3:EA",
// 		FullName:   "Alice Johnson",
// 		Batch:      "Batch 2025",
// 		Department: "Computer Science",
// 	}
// 	// name := "abe"

//  data, err := json.Marshal(student)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println((data))


// 	// ages := []int {1,2,3,4,5}

// 	// fmt.Println(ages)
// }


func main() {
	jsonStr := `{"id":"C6:65:D3:EA","full_name":"Alice Johnson","batch":"Batch 2025","department":"Computer Science"}`

	var student Student
	err := json.Unmarshal([]byte(jsonStr), &student)
	if err != nil {
		panic(err)
	}

	fmt.Println(student.FullName) // Output: Alice Johnson
}
