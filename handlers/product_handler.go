// handlers/product_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"go-ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler структура, содержащая объект базы данных
type Handler struct {
	DB *gorm.DB
}

// SetupRoutes настраивает маршруты и передаёт объект базы данных
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	h := &Handler{DB: db}

	products := router.Group("/products")
	{
		products.GET("", h.GetProducts)
		products.GET("/:id", h.GetProduct)
		products.POST("", h.CreateProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
	}
}

// GetProducts возвращает список всех товаров из базы данных
// @Summary Получить список товаров
// @Description Возвращает список всех товаров
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product
// @Failure 500 {object} models.ErrorResponse
// @Router /products [get]
func (h *Handler) GetProducts(c *gin.Context) {
	var products []models.Product
	ctx := c.Request.Context()
	if err := h.DB.WithContext(ctx).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Не удалось получить товары"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct возвращает товар по заданному ID из базы данных
// @Summary Получить товар по ID
// @Description Возвращает товар по заданному ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "ID товара"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Неверный ID товара"})
		return
	}

	var product models.Product
	ctx := c.Request.Context()
	if err := h.DB.WithContext(ctx).First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Товар не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Ошибка при получении товара"})
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct создает новый товар
// @Summary Создать новый товар
// @Description Создание нового товара
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.CreateProductRequest true "Создать товар"
// @Success 201 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Неверные данные: " + err.Error()})
		return
	}

	product := models.Product{
		Name:  req.Name,
		Price: req.Price,
		Desc:  req.Desc,
	}

	ctx := c.Request.Context()
	if err := h.DB.WithContext(ctx).Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Не удалось создать товар"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct обновляет существующий товар
// @Summary Обновить товар
// @Description Обновление существующего товара по ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "ID товара"
// @Param product body models.CreateProductRequest true "Обновить товар"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Неверный ID товара"})
		return
	}

	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Неверные данные: " + err.Error()})
		return
	}

	var existingProduct models.Product
	ctx := c.Request.Context()
	if err := h.DB.WithContext(ctx).First(&existingProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Товар не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Ошибка при получении товара"})
		}
		return
	}

	existingProduct.Name = req.Name
	existingProduct.Price = req.Price
	existingProduct.Desc = req.Desc

	if err := h.DB.WithContext(ctx).Save(&existingProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Не удалось обновить товар"})
		return
	}

	c.JSON(http.StatusOK, existingProduct)
}

// DeleteProduct удаляет товар по заданному ID
// @Summary Удалить товар
// @Description Удаление товара по заданному ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "ID товара"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Неверный ID товара"})
		return
	}

	ctx := c.Request.Context()
	if err := h.DB.WithContext(ctx).Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Не удалось удалить товар"})
		return
	}

	c.JSON(http.StatusOK, models.MessageResponse{Message: "Товар удален"})
}
