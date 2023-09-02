package main

import (
	"fmt"
	"log"

	"github.com/AkashGit21/typeface-assignment/models"
	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the persistent DB for metadata storage
	metdataDBLayer, err := utils.NewPersistenceDBLayer()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Test the persistent DB by inserting a new row.
	fileMetadata := models.Metadata{
		Filename:    "abc.txt",
		SizeInBytes: int64(4780090),
		S3ObjectKey: "s3://url.in/obj1",
		Description: "some description",
		Status:      1,
	}

	id, err := metdataDBLayer.SaveRecord(fileMetadata)
	if err != nil {
		log.Fatalf("Error saving record: %v", err)
	}
	fmt.Println("Record saved successfully with id: ", id)
}
