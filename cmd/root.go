package cmd

import (
	"os"

	"github.com/prixladi/termfi/config"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:          "termfi",
	Short:        "Minimal application for uploading and sharing files",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func onCobraInit() {
	err := config.Init(cfgFile)
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(onCobraInit)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.termfi.yaml)")

	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(configCmd)
}
