package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	ctx          context.Context
	fileClient   *minio.Client
	bucket       string
	httpEndpoint string
	options      StorageOptions
}

type StorageOptions struct {
	Endpoint    string
	AccessKeyId string
	AccessKey   string
	Bucket      string
}

func New(ctx context.Context, options StorageOptions) (*Storage, error) {
	endpoint, useSsl := parseEndpoint(options.Endpoint)

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(options.AccessKeyId, options.AccessKey, ""),
		Secure: useSsl,
	})
	if err != nil {
		return nil, err
	}

	store := Storage{
		ctx:          ctx,
		fileClient:   minioClient,
		bucket:       options.Bucket,
		httpEndpoint: getHttpEndpoint(endpoint, useSsl),
	}

	err = store.ensureBuckets()
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func parseEndpoint(endpoint string) (string, bool) {
	useSsl := false

	if strings.HasPrefix(endpoint, "https://") {
		endpoint = strings.Replace(endpoint, "https://", "", 1)
		useSsl = true
	}
	if strings.HasPrefix(endpoint, "http://") {
		endpoint = strings.Replace(endpoint, "http://", "", 1)
	}

	return endpoint, useSsl
}

func getHttpEndpoint(endpoint string, useSsl bool) string {
	protocol := "http://"
	if useSsl {
		protocol = "https://"
	}

	return fmt.Sprintf("%s%s", protocol, endpoint)
}
