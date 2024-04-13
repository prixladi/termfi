package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/prixladi/termfi/download"
	"github.com/prixladi/termfi/storage"
	"github.com/prixladi/termfi/utils"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"dw"},
	Short:   "Download a termfi file",
	Example: "termfi get http://localhost:9000/termfi/tf--ebbe7001-e8f5-4860-8ad4-d0aab490ba85--f--wp-4.4.2.zip",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts 1 arg(s), received %v", len(args))
		}
		if !utils.IsUrlValid(args[0]) {
			return fmt.Errorf("argument is not a valid url: %s", args[0])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]

		fileInfo := storage.GetFileInfo(url)

		if !fileInfo.IsTermfiFile {
			noCheck, err := strconv.ParseBool(cmd.Flag("no-check").Value.String())
			if !noCheck || err != nil {
				return errors.New("Provided file is not a termfi file. If you want to download it anyway you --no-check flag")
			}
		}

		download.DownloadFile(url, fmt.Sprintf("./%s", fileInfo.FileName), fileInfo.Size)
		return nil
	},
}

func init() {
	getCmd.Flags().BoolP("no-check", "n", false, "Skips some of the checks, for example allows to download file that was not uploaded using termfi")
}
