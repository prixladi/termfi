package cmd

import (
	"github.com/prixladi/termfi/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"cf"},
	Short:   "Set and persists configuration keys for termfi",
	Long: `Set and persists configuration keys for termfi

supported configuration keys: 
 		- 'storageEndpoint', 
 		- 'storageAccessKeyId', 
 		- 'storageAccessKey', 
 		- 'storageBucket'
`,
	Example: "termfi config --storageEndpoint https://example-storage.com",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.Write()
	},
}

func init() {
	configCmd.Flags().String("storageEndpoint", "", "Sets url of storage endpoint (S3 compatible storage)")
	configCmd.Flags().String("storageAccessKeyId", "", "Sets access key id to storage")
	configCmd.Flags().String("storageAccessKey", "", "Sets access key to storage")
	configCmd.Flags().String("storageBucket", "", "Sets storage bucket")

	config.BindStorageEndpoint(configCmd.Flags().Lookup("storageEndpoint"))
	config.BindStorageAccessKeyId(configCmd.Flags().Lookup("storageAccessKeyId"))
	config.BindStorageAccessKey(configCmd.Flags().Lookup("storageAccessKey"))
	config.BindStorageBucket(configCmd.Flags().Lookup("storageBucket"))

	configCmd.AddCommand(configPwdCmd)
}
