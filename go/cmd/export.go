package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/teawithsand/webpage/domain/common/export"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports data like translations and LIVR rules from go app to extenal files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ex := export.Exporter{
			Dir: "/workspace/go/__exported/",
		}

		ex.MustClear()
		ex.MustExportValidations()
		ex.MustExportTranslations()
		ex.MustExportTypescript()

		fmt.Println("Export ran OK")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
