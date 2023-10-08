package models

import "gorm.io/gorm"

type User struct {
	ID    int
	Email string
}

type Report struct {
	NumberUsers int
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
	SKU   string
}
