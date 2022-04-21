package cmd

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/teawithsand/webpage/domain/webapp"
)

func getAllFilenames(fs *embed.FS, path string) (out []string) {
	if len(path) == 0 {
		path = "."
	}
	entries, err := fs.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		fp := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			res := getAllFilenames(fs, fp)
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return
}

// debugassetsCmd represents the debugassets command
var debugassetsCmd = &cobra.Command{
	Use:   "debugassets",
	Short: "Prints all embedded assets",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files := getAllFilenames(&webapp.EmbeddedAssets, "")
		fmt.Printf("Got %d entries\n", len(files))
		for _, f := range files {
			fmt.Println(f)
		}
	},
}

func init() {
	rootCmd.AddCommand(debugassetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debugassetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugassetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
