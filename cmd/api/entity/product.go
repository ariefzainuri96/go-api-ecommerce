package entity

import (
	_ "gorm.io/gorm"
)

// @Model
type Product struct {
	BaseEntity
	Name        string `gorm:"type:varchar(255);not null;column:name" json:"name"`
	Description string `gorm:"type:text;not null;column:description" json:"description"`
	Price       int64  `gorm:"type:bigint;not null;column:price" json:"price"`
	Quantity    int    `gorm:"type:integer;not null;column:quantity" json:"quantity"`
}
