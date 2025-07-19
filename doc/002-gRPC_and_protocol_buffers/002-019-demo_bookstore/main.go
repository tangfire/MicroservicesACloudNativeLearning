package main

import "fmt"

func main() {
	db, err := NewDB("bookstore.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	// db

}
