package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := myEnv["FOO"]
	fmt.Printf("the env variable stored in Key:FOO is %s", s3Bucket)

	// now do something with s3 or whatever
}
