package utils

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3Client() (*s3.S3, error) {
	// Create an AWS session
	s3Region := GetEnvValue("S3_REGION", "ap-south-1")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(s3Region),
		// Using server based access control for security purposes
	})
	if err != nil {
		return nil, errors.New("error creating AWS session")
	}

	// Create an S3 client, and return it
	return s3.New(awsSession), nil
}
