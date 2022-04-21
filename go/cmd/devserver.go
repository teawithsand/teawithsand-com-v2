package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/teawithsand/webpage/domain"
	"github.com/teawithsand/webpage/domain/webapp"
	"github.com/teawithsand/webpage/util"
)

// devserverCmd represents the devserver command
var devserverCmd = &cobra.Command{
	Use:   "devserver",
	Short: "Runs dev server serving static files from __dist dir",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Initializing DEV web server")
		di, err := domain.ConstructDI()
		if err != nil {
			panic(err)
		}

		defer di.Delete()

		log.Println("DI constraction ok")

		r := di.Get(webapp.DevRunnerDI).(util.Runner)
		err = r.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(devserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
