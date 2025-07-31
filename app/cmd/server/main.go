package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/youruser/order-app/internal/config"
	"github.com/youruser/order-app/internal/handler"
	"github.com/youruser/order-app/internal/repository"
	"github.com/youruser/order-app/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	m, err := migrate.New("file://migrations", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	cityRepo := repository.NewCityRepo(db)
	citySvc := service.NewCityService(cityRepo)
	cityHandler := handler.NewCityHandler(citySvc)

	companyRepo := repository.NewCompanyRepo(db)
	companySvc := service.NewCompanyService(companyRepo)
	companyHandler := handler.NewCompanyHandler(companySvc)

	firmRepo := repository.NewFirmRepo(db)
	firmSvc := service.NewFirmService(firmRepo)
	firmHandler := handler.NewFirmHandler(firmSvc)

	orderRepo := repository.NewOrderRepo(db)
	orderSvc := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderSvc)

	productRepo := repository.NewProductRepo(db)
	productSvc := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productSvc)

	product_cityRepo := repository.NewProductCityRepo(db)
	product_citySvc := service.NewProductCityService(product_cityRepo)
	product_cityHandler := handler.NewProductCityHandler(product_citySvc)

	router := gin.Default()
	api := router.Group("/api")
	{
		cityHandler.Register(api)
		companyHandler.Register(api)
		firmHandler.Register(api)
		orderHandler.Register(api)
		productHandler.Register(api)
		product_cityHandler.Register(api)
	}

	log.Printf("server running on %s", cfg.Port)
	router.Run(cfg.Port)
}
