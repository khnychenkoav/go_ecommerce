// models/request.go
package models

// CreateProductRequest представляет структуру запроса для создания товара
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=3"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"desc" binding:"required"`
}

// RegisterRequest представляет структуру запроса для регистрации пользователя
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest представляет структуру запроса для входа пользователя
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}
