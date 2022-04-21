package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/teawithsand/webpage/domain"
	"github.com/teawithsand/webpage/domain/webapp"
	"github.com/teawithsand/webpage/util"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Servers webpage over HTTP",
	Long:  `Servers webpage over HTTP. All confirugration comes from ENV variables. SSL is not supported, use NGINX instead for now.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Initializing web server")
		di, err := domain.ConstructDI()
		if err != nil {
			panic(err)
		}

		defer di.Delete()

		log.Println("DI constraction ok")

		r := di.Get(webapp.RunnerDI).(util.Runner)
		err = r.Run()
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.webpage.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
