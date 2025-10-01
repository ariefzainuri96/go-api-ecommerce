package data

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);not null;column:name"`
	Description string `gorm:"type:text;not null;column:description"`
	Price       int64  `gorm:"type:bigint;not null;column:price"`
	Quantity    int    `gorm:"type:integer;not null;column:quantity"`
}
