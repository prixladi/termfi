package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const storageEndpoint = "storage-endpoint"
const storageAccessKey = "storage-access-key"
const storageAccessKeyId = "storage-access-key-id"
const storageBucket = "storage-bucket"

var configKeys = []string{storageEndpoint, storageAccessKey, storageAccessKeyId, storageBucket}

func GetStorageEndpoint() string {
	return viper.GetString(storageEndpoint)
}

func GetStorageAccessKey() string {
	return viper.GetString(storageAccessKey)
}

func GetStorageAccessKeyId() string {
	return viper.GetString(storageAccessKeyId)
}

func GetStorageBucket() string {
	return viper.GetString(storageBucket)
}

func BindStorageEndpoint(flag *pflag.Flag) error {
	return viper.BindPFlag(storageEndpoint, flag)
}

func BindStorageAccessKey(flag *pflag.Flag) error {
	return viper.BindPFlag(storageAccessKey, flag)
}

func BindStorageAccessKeyId(flag *pflag.Flag) error {
	return viper.BindPFlag(storageAccessKeyId, flag)
}

func BindStorageBucket(flag *pflag.Flag) error {
	return viper.BindPFlag(storageBucket, flag)
}
