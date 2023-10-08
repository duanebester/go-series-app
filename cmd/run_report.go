package cmd

import (
	"fmt"
	"go-series-app/database"
	"go-series-app/services"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runReportCmd)
}

var runReportCmd = &cobra.Command{
	Use:   "run-report",
	Short: "Hi",
	Long:  `Hi`,
	Run: func(cmd *cobra.Command, args []string) {
		// new db
		db := database.NewDB(false)
		service := services.NewService(db)
		report := service.GetReport()
		fmt.Println("report: ", report)
	},
}
