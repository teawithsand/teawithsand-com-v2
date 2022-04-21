package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/teawithsand/webpage/domain"
	"github.com/teawithsand/webpage/domain/common/dikey"
	"github.com/teawithsand/webpage/domain/db"
	"github.com/teawithsand/webpage/domain/domtestutil/fixture"
)

// fixtureCmd represents the fixture command
var fixtureCmd = &cobra.Command{
	Use:   "fixture",
	Short: "Applies fixtures to database database",
	Long:  `Clears entire database`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Populating database with fixtures")
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

		err = fixture.Apply(cmd.Context(), db)
		if err != nil {
			panic(err)
		}

		log.Println("Fixtures applied successfully")
	},
}

func init() {
	dbCmd.AddCommand(fixtureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixtureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixtureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
