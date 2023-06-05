package main

import (
	"log"
	"os"

	"github.com/jamesdearlove/willow/apps/upClient/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	upToken := os.Getenv("UP_TOKEN")
	api.MakeTransactionListRequest(upToken)

	

}
