// models/request.go
package models

// CreateProductRequest представляет структуру запроса для создания товара
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required,min=3"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Desc  string  `json:"desc" binding:"required"`
}
