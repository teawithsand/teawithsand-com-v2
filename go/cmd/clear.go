package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/teawithsand/webpage/domain"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/domain/db"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears database",
	Long:  `Clears entire database`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Clearing database (DO NOT RUN IN PROD!!1)")
		di, err := domain.ConstructDI()
		if err != nil {
			panic(err)
		}

		defer di.Delete()

		var db *db.TypedDB
		err = di.Fill(dikey.DBDatabaseDI, &db)
		if err != nil {
			panic(err)
		}

		err = db.Clear(cmd.Context())
		if err != nil {
			panic(err)
		}

		log.Println("Database is empty now")
	},
}

func init() {
	dbCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
