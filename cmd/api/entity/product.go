package entity

import (
	"github.com/shopspring/decimal"
	_ "gorm.io/gorm"
)

// @Model
type Product struct {
	BaseEntity
	Name        string          `gorm:"type:varchar(255);not null;column:name" json:"name"`
	Description string          `gorm:"type:text;not null;column:description" json:"description"`
	Price       decimal.Decimal `gorm:"type:numeric(18, 4);not null;column:price" json:"price"`
	Quantity    int             `gorm:"type:integer;not null;column:quantity" json:"quantity"`
}

func (Product) TableName() string {
	return "products"
}
