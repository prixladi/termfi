package storage

import (
	"fmt"

	"github.com/minio/minio-go/v7"
)

const location = "us-east-1"

const policy = `{
	"Statement": [
	 {
	  "Action": [
	   "s3:GetObject"
	  ],
	  "Effect": "Allow",
	  "Principal": {
	   "AWS": [
		"*"
	   ]
	  },
	  "Resource": [
	   "arn:aws:s3:::%s/*"
	  ]
	 }
	],
	"Version": "2012-10-17"
   }`

func (storage *Storage) ensureBuckets() error {
	client := storage.fileClient
	ctx := storage.ctx
	bucketName := storage.bucket

	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists != nil {
			return errBucketExists
		}

		if !exists {
			return err
		}
	}

	err = client.SetBucketPolicy(ctx, bucketName, fmt.Sprintf(policy, bucketName))
	if err != nil {
		return err
	}

	return nil
}
