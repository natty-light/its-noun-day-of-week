package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateS3Downloader(env Env) (*s3manager.Downloader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(env.AwsRegion),
	})
	if err != nil {
		fmt.Println("ERROR: Unable to connect to AWS")
		return nil, err
	}
	downloader := s3manager.NewDownloader(sess)

	return downloader, nil
}

func GenerateS3ObjectURL(bucket string, outputFileName string) string {
	return fmt.Sprintf(`https://%s.s3.amazonaws.com/%s`, bucket, outputFileName)
}
