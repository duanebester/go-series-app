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
	// drop product table if exists
	// db.Migrator().DropTable(&models.Product{})
	insertProduct := models.NewProduct(&models.Product{Name: "D42", Price: 100, SKU: "abc123"})

	db.Create(insertProduct)
	fmt.Printf("insert ID: %d, Name: %s, Prict: %d\n",
		insertProduct.ID, insertProduct.Name, insertProduct.Price)

	readProduct := &models.Product{}
	db.First(&readProduct, "name = ?", "D42") // find product with code D42

	fmt.Printf("read ID: %d, Name: %s, Prict: %d\n",
		readProduct.ID, readProduct.Name, readProduct.Price)

	// truncate users table
	// db.Migrator().DropTable(&models.User{})
	insertUser := models.NewUser(&models.User{Username: "user1", Email: "test@test.com", Password: "password"})
	db.Create(insertUser)
}

var seedDbCmd = &cobra.Command{
	Use:   "seed-db",
	Short: "Hi",
	Long:  `Hi`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.NewDB(true)
		seedDb(db)
	},
}
