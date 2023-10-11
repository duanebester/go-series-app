package models

import (
	"go-series-app/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUser(opts *User) *User {
	hashedPassword, err := utils.HashPassword(opts.Password)
	if err != nil {
		panic("failed to seed db: failed to hash password")
	}
	return &User{
		ID:        uuid.New(),
		Username:  opts.Username,
		Email:     opts.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type User struct {
	ID        uuid.UUID      `gorm:"primaryKey,type:uuid"`
	Username  string         `gorm:"uniqueIndex;not null;size:50;" json:"username"`
	Email     string         `gorm:"uniqueIndex;not null;size:255;" json:"email"`
	Password  string         `gorm:"not null;" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func NewProduct(opts *Product) *Product {
	return &Product{
		ID:        uuid.New(),
		Name:      opts.Name,
		Price:     opts.Price,
		SKU:       opts.SKU,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Product struct {
	ID        uuid.UUID      `gorm:"primaryKey,type:uuid"`
	Name      string         `json:"name"`
	Price     uint           `json:"price"`
	SKU       string         `gorm:"uniqueIndex" json:"sku"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Report struct {
	NumberUsers int
}
