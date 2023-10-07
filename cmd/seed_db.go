package cmd

import (
	"fmt"
	"go-series-app/database"
	"go-series-app/models"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(seedDbCmd)
}

func seedDb(db *gorm.DB) {
	db.AutoMigrate(&models.Product{})

	insertProduct := &models.Product{Code: "D42", Price: 100}

	db.Create(insertProduct)
	fmt.Printf("insert ID: %d, Code: %s, Prict: %d\n",
		insertProduct.ID, insertProduct.Code, insertProduct.Price)

	readProduct := &models.Product{}
	db.First(&readProduct, "code = ?", "D42") // find product with code D42

	fmt.Printf("read ID: %d, Code: %s, Prict: %d\n",
		readProduct.ID, readProduct.Code, readProduct.Price)
}

var seedDbCmd = &cobra.Command{
	Use:   "seed-db",
	Short: "Hi",
	Long:  `Hi`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDB()
		seedDb(db)
	},
}
