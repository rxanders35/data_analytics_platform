package storage

import (
	"github.com/minio/minio-go"
)

type Client struct {
	minioClient *minio.Client
}

func NewClient() {}
