package cmd

import (
	"go-series-app/api"
	"go-series-app/database"
	"go-series-app/services"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runApiCmd)
}

var runApiCmd = &cobra.Command{
	Use:   "run-api",
	Short: "Hi",
	Long:  `Hi`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDB(false)
		services := services.NewService(db)
		api := api.NewApi(services)
		log.Fatal(api.Listen(":3000"))
	},
}
