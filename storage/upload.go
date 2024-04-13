package storage

import (
	"fmt"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/minio/minio-go/v7"
	"github.com/prixladi/termfi/utils"
)

func (storage *Storage) UploadObject(filePath string, options minio.PutObjectOptions) (string, error) {
	fmt.Printf("Uploading file '%s' to server\n", filePath)

	objectName := createTermfiObjectName(filePath)

	fileSize, err := getLocalFileSize(filePath)
	if err != nil {
		return "", err
	}

	progress := pb.New64(fileSize)
	progress.Start()

	options.Progress = progress

	start := time.Now()

	info, err := storage.fileClient.FPutObject(
		storage.ctx,
		storage.bucket,
		objectName,
		filePath,
		options)
	if err != nil {
		return "", err
	}

	fileUrl := fmt.Sprintf("%s/%s/%s", storage.httpEndpoint, info.Bucket, info.Key)

	utils.ReplacefLine("Upload to '%s' completed in %s\n", fileUrl, time.Since(start))

	return fileUrl, nil
}
