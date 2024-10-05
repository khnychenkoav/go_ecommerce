// models/cart.go
package models

type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
