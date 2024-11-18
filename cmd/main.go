package main

import (
	"fmt"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
}
