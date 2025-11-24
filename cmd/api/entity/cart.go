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

/*
	for filtering field use like this for [carts] table:
	- carts.quantity -> even for current table filtering, always call the table name like this
	- products.name -> filter using products table with field name -> 
	remember to not using struct field -> always use real tables and field name
*/

func (Cart) TableName() string {
	return "carts"
}