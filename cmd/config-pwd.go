package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPwdCmd = &cobra.Command{
	Use:     "pwd",
	Short:   "Prints config path",
	Example: "vim $(termfi config pwd)",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(viper.ConfigFileUsed())
	},
}
