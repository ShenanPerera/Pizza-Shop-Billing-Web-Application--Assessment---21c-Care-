package main

import (
	"main/controllers"
	"main/middleware"
	"main/models"
	"main/routes"

	"log"
	"net/http"
)

// type Item struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Name     string `gorm:"size:100"`
// 	Price    float64
// 	Category string `gorm:"size:50"`
// }

func main() {
	ConnectDB()

	controllers.SetDB(DB) // Set the database connection in controllers package

	err := DB.AutoMigrate(
		&models.Customer{},
		&models.Invoice{},
		&models.Item{},
		&models.Order{},
		&models.OrderItem{},
		&models.Pizza{},
		&models.Topping{},
		&models.Beverage{}) // GORM creates the table if not exists
	if err != nil {
		panic("Failed to migrate Customer model: " + err.Error())
	}

	DB.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Item{})
	DB.AutoMigrate(&models.Invoice{})

	router := routes.SetupRoutes()

	// Apply CORS middleware
	handler := middleware.EnableCORS(router)
	log.Println("🍕 Pizza Shop API Server starting on :8080")
	log.Println("📋 Available endpoints:")
	log.Println("   Item Management: /api/items")
	log.Println("   Invoice Management: /api/invoices")
	log.Println("   Customer Management: /api/customers")
	log.Println("   Order Management: /api/orders")
	// log.Println("   Dashboard: /api/dashboard/stats")

	log.Fatal(http.ListenAndServe(":8080", handler))

}
