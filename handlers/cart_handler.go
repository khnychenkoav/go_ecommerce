// handlers/cart_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"go-ecommerce/models"

	"github.com/gin-gonic/gin"
)

// AddToCart добавляет товар в корзину пользователя
// @Summary Добавить товар в корзину
// @Description Добавляет товар в корзину текущего пользователя
// @Tags cart
// @Accept  json
// @Produce  json
// @Param cart_item body models.CartItem true "Данные товара"
// @Success 201 {object} models.CartItem
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/cart/add [post]
func (h *Handler) AddToCart(c *gin.Context) {
	var req models.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Пользователь не авторизован"})
		return
	}

	var product models.Product
	if err := h.DB.First(&product, req.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Товар не найден"})
		return
	}

	// Создайте элемент корзины
	cartItem := models.CartItem{
		UserID:    userID.(uint),
		ProductID: req.ProductID,
		Quantity:  uint(req.Quantity), // Преобразование в uint из int
	}

	if err := h.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Ошибка при добавлении в корзину"})
		return
	}

	c.JSON(http.StatusCreated, cartItem)
}

// GetCartItems возвращает товары из корзины пользователя
// @Summary Получить товары из корзины
// @Description Возвращает список товаров, добавленных в корзину текущего пользователя
// @Tags cart
// @Accept  json
// @Produce  json
// @Success 200 {array} models.CartItem
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/cart [get]
func (h *Handler) GetCartItems(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Пользователь не авторизован"})
		return
	}

	var cartItems []models.CartItem
	if err := h.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Ошибка получения корзины"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

// RemoveFromCart удаляет товар из корзины пользователя
// @Summary Удалить товар из корзины
// @Description Удаляет товар из корзины текущего пользователя по ID позиции
// @Tags cart
// @Accept  json
// @Produce  json
// @Param id path int true "ID позиции в корзине"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /api/cart/remove/{id} [delete]
func (h *Handler) RemoveFromCart(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Неверный ID товара"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Пользователь не авторизован"})
		return
	}

	if err := h.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Ошибка удаления из корзины"})
		return
	}

	c.JSON(http.StatusOK, models.MessageResponse{Message: "Товар удален из корзины"})
}
