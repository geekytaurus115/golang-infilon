package main

import (
	"log"

	"github.com/geekytaurus115/golang-infilon/handler"
)

func main() {
	err := handler.InitDB(handler.DATA_SOURCE)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return
	}
	defer handler.CloseDB()

	handler.StartApp()
}
