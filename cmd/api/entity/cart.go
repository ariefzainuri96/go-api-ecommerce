package entity

import (
	_ "gorm.io/gorm"
)

// @Model
type Cart struct {
	BaseEntity
	ProductId int64   `gorm:"type:bigint;not null;column:product_id" json:"product_id"`
	UserId    int64   `gorm:"type:bigint;not null;column:user_id" json:"user_id"`
	Quantity  int     `gorm:"type:integer;not null;column:quantity" json:"quantity"`
	Product   Product `json:"product"`
}
