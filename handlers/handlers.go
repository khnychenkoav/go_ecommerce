// handlers/handlers.go
package handlers

import (
	"go-ecommerce/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler структура, содержащая объект базы данных
type Handler struct {
	DB        *gorm.DB
	JWTSecret string
}

// SetupRoutes настраивает маршруты и передаёт объект базы данных
func SetupRoutes(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	h := &Handler{DB: db, JWTSecret: jwtSecret}

	// Маршруты аутентификации
	router.POST("/api/register", h.Register)
	router.POST("/api/login", h.Login)

	// Маршруты продуктов (общедоступные)
	router.GET("/api/products", h.GetProducts)
	router.GET("/api/products/:id", h.GetProduct)

	// Защищённые маршруты
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware(jwtSecret))
	{
		// Маршруты продуктов (требуют аутентификации)
		auth.POST("/api/products", h.CreateProduct)
		auth.PUT("/api/products/:id", h.UpdateProduct)
		auth.DELETE("/api/products/:id", h.DeleteProduct)

		// Маршруты корзины
		auth.GET("/api/cart", h.GetCartItems)
		auth.POST("/api/cart/add", h.AddToCart)
		auth.DELETE("/api/cart/remove/:id", h.RemoveFromCart)
	}
}
