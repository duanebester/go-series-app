package cmd

import (
	"fmt"
	"go-series-app/database"
	"go-series-app/models"
	"go-series-app/utils"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(seedDbCmd)
}

func seedDb(db *gorm.DB) {
	insertProduct := &models.Product{Code: "D42", Price: 100}

	db.Create(insertProduct)
	fmt.Printf("insert ID: %d, Code: %s, Prict: %d\n",
		insertProduct.ID, insertProduct.Code, insertProduct.Price)

	readProduct := &models.Product{}
	db.First(&readProduct, "code = ?", "D42") // find product with code D42

	fmt.Printf("read ID: %d, Code: %s, Prict: %d\n",
		readProduct.ID, readProduct.Code, readProduct.Price)

	hashedPassword, err := utils.HashPassword("password")
	if err != nil {
		panic("failed to seed db: failed to hash password")
	}
	insertUser := &models.User{Username: "user1", Email: "test@test.com", Password: hashedPassword}
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
