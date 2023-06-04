package utils

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/bwmarrin/discordgo"
)

type S3DataSource struct {
	Downloader *manager.Downloader
	Client     *s3.Client
}

func CreateS3Client(env Env) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	cfg.Region = env.AwsRegion
	if err != nil {
		return nil, err
	}
	Client := s3.NewFromConfig(cfg)
	return Client, nil
}

func CreateS3Downloader(client *s3.Client) *manager.Downloader {
	var partMiBs int64 = 10
	downloader := manager.NewDownloader(client, func(d *manager.Downloader) {
		d.PartSize = partMiBs * 1024 * 1024
	})
	return downloader
}

func (s S3DataSource) ListAllFilesInFolder(env Env, dayOfWeek string) ([]types.Object, error) {
	res, err := s.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(env.S3Bucket),
		Prefix: aws.String(dayOfWeek),
	})
	if err != nil {
		return nil, fmt.Errorf("error listing objects: %s", err.Error())
	}
	return res.Contents, nil
}

func (s S3DataSource) DownloadAndParseFile(env Env, key string) (*discordgo.File, error) {
	res, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(env.S3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching objects: %s", err.Error())
	}
	defer res.Body.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(res.Body)
	mimetype := *res.ContentType
	file := &discordgo.File{Reader: bytes.NewReader(buffer.Bytes()), Name: key, ContentType: mimetype}

	return file, nil
}

func (s S3DataSource) DownloadAndParseFileViaDownloader(env Env, key string, mimetype string) (*discordgo.File, error) {
	buffer := manager.NewWriteAtBuffer([]byte{})

	_, err := s.Downloader.Download(context.TODO(), buffer, &s3.GetObjectInput{
		Bucket: aws.String(env.S3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("S3 download error", err)
		return nil, err
	}
	file := &discordgo.File{Reader: bytes.NewReader(buffer.Bytes()), Name: key, ContentType: mimetype}
	return file, nil
}
