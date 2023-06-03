package utils

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bwmarrin/discordgo"
)

type S3DataSource struct {
	Downloader s3manager.Downloader
}

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

func (s S3DataSource) DownloadAndParseFile(env Env, name string, mimetype string) (*discordgo.File, error) {
	var file *discordgo.File
	buffer := aws.NewWriteAtBuffer([]byte{})
	_, err := s.Downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(env.S3Bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		fmt.Println("S3 download error", err)
		return nil, err
	}
	file = &discordgo.File{Reader: bytes.NewReader(buffer.Bytes()), Name: name, ContentType: mimetype}
	return file, nil
}
