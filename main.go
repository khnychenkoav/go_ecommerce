// main.go
package main

import (
	"log"
	"os"

	"go-ecommerce/handlers"
	"go-ecommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "go-ecommerce/docs" // Swagger docs

	swaggerFiles "github.com/swaggo/files"     // Swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger middleware
)

// @title Go E-commerce API
// @version 1.0
// @description API для интернет-магазина на Go.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используется окружение системы")
	}

	// Чтение переменных окружения
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")

	// Формирование строки подключения (DSN)
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=UTC"

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Миграция схемы
	if err := models.RunMigrations(db); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	router := gin.Default()

	// Настройка маршрутов с использованием базы данных
	handlers.SetupRoutes(router, db, jwtSecret)

	// Настройка маршрута Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	log.Println("Сервер запущен на порту 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
