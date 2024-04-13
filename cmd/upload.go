package cmd

import (
	"fmt"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/prixladi/termfi/config"
	"github.com/prixladi/termfi/storage"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"up"},
	Short:   "Upload a file",
	Example: "termfi upload ./tst/wordpress-4.4.2.zip",
	Args:    cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		missingKeys := []string{}

		if config.GetStorageEndpoint() == "" {
			missingKeys = append(missingKeys, "storageEndpoint")
		}
		if config.GetStorageAccessKey() == "" {
			missingKeys = append(missingKeys, "storageAccessKey")
		}
		if config.GetStorageAccessKeyId() == "" {
			missingKeys = append(missingKeys, "storageAccessKeyId")
		}
		if config.GetStorageBucket() == "" {
			missingKeys = append(missingKeys, "storageBucket")
		}

		if len(missingKeys) != 0 {
			return fmt.Errorf(
				"Missing some configuration keys for storage connection (%s). Call `termfi config -h` for more info.",
				strings.Join(missingKeys, ", "))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Fix windows path
		path := strings.Replace(args[0], "\\", "/", -1)

		storageOptions := storage.StorageOptions{
			Endpoint:    config.GetStorageEndpoint(),
			AccessKeyId: config.GetStorageAccessKeyId(),
			AccessKey:   config.GetStorageAccessKey(),
			Bucket:      config.GetStorageBucket(),
		}

		store, err := storage.New(cmd.Context(), storageOptions)
		if err != nil {
			return err
		}

		info, err := store.UploadObject(path, minio.PutObjectOptions{})
		if err != nil {
			return err
		}

		fmt.Printf(`You can download the file using termfi:
    
  termfi download %s

  or 

  curl -O %s

note that latter method will not preserve correct filename
`, info, info)

		return nil
	},
}
