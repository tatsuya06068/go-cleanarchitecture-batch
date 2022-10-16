package main

import (
   "log"
//    "os"
   "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

   log.Println("exec batch")

}