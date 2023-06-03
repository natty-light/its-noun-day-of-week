package utils

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/bwmarrin/discordgo"
)

type S3DataSource struct {
	Downloader *manager.Downloader
	Client     *s3.Client
}

func CreateS3Client(env Env) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	Client := s3.NewFromConfig(cfg)
	return Client, nil
}

func (s S3DataSource) CreateS3Downloader() *manager.Downloader {
	var partMiBs int64 = 10
	downloader := manager.NewDownloader(s.Client, func(d *manager.Downloader) {
		d.PartSize = partMiBs * 1024 * 1024
	})
	return downloader
}

func (s S3DataSource) DownloadAndParseFile(env Env, key string, mimetype string) (*discordgo.File, error) {
	var file *discordgo.File
	buffer := manager.NewWriteAtBuffer([]byte{})

	_, err := s.Downloader.Download(context.TODO(), buffer, &s3.GetObjectInput{
		Bucket: aws.String(env.S3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("S3 download error", err)
		return nil, err
	}
	file = &discordgo.File{Reader: bytes.NewReader(buffer.Bytes()), Name: key, ContentType: mimetype}
	return file, nil
}
