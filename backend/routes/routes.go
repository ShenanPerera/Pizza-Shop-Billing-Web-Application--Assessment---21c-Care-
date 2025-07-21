// routes/routes.go
package routes

import (
	"main/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// API prefix
	api := r.PathPrefix("/api").Subrouter()

	// Customer routes
	api.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	api.HandleFunc("/customers/telno/{telno}", controllers.GetCustomerByTelNo).Methods("GET")
	api.HandleFunc("/customers/{id:[0-9]+}", controllers.GetCustomerByID).Methods("GET")
	api.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	api.HandleFunc("/customers/{id:[0-9]+}", controllers.UpdateCustomer).Methods("PUT")
	api.HandleFunc("/customers/{id:[0-9]+}", controllers.DeleteCustomer).Methods("DELETE")

	// Item management routes
	api.HandleFunc("/items", controllers.GetItems).Methods("GET")
	api.HandleFunc("/items/{id:[0-9]+}", controllers.GetItemByID).Methods("GET")
	api.HandleFunc("/items/type/{type}", controllers.GetItemsByType).Methods("GET")
	api.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	api.HandleFunc("/items/{id:[0-9]+}", controllers.UpdateItem).Methods("PUT")
	api.HandleFunc("/items/{id:[0-9]+}", controllers.DeleteItem).Methods("DELETE")

	// // Pizza routes
	api.HandleFunc("/pizzas", controllers.GetPizzas).Methods("GET")
	api.HandleFunc("/pizzas/{id:[0-9]+}", controllers.GetPizzaByID).Methods("GET")
	api.HandleFunc("/pizzas", controllers.CreatePizza).Methods("POST")
	api.HandleFunc("/pizzas/{id:[0-9]+}", controllers.UpdatePizza).Methods("PUT")
	api.HandleFunc("/pizzas/{id:[0-9]+}", controllers.DeletePizza).Methods("DELETE")

	// // Topping routes
	api.HandleFunc("/toppings", controllers.GetToppings).Methods("GET")
	api.HandleFunc("/toppings/{id:[0-9]+}", controllers.GetToppingByID).Methods("GET")
	api.HandleFunc("/toppings", controllers.CreateTopping).Methods("POST")
	api.HandleFunc("/toppings/{id:[0-9]+}", controllers.UpdateTopping).Methods("PUT")
	api.HandleFunc("/toppings/{id:[0-9]+}", controllers.DeleteTopping).Methods("DELETE")

	// Beverage routes
	api.HandleFunc("/beverages", controllers.GetBeverages).Methods("GET")
	api.HandleFunc("/beverages/{id:[0-9]+}", controllers.GetBeverageByID).Methods("GET")
	api.HandleFunc("/beverages", controllers.CreateBeverage).Methods("POST")
	api.HandleFunc("/beverages/{id:[0-9]+}", controllers.UpdateBeverage).Methods("PUT")
	api.HandleFunc("/beverages/{id:[0-9]+}", controllers.DeleteBeverage).Methods("DELETE")

	// // Order routes
	api.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	api.HandleFunc("/orders/{id:[0-9]+}", controllers.GetOrderByID).Methods("GET")
	api.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	api.HandleFunc("/orders/{id:[0-9]+}/status", controllers.UpdateOrderStatus).Methods("PUT")

	// Invoice/Bill routes
	api.HandleFunc("/invoices", controllers.GetInvoices).Methods("GET")
	api.HandleFunc("/invoices/{id:[0-9]+}", controllers.GetInvoiceByID).Methods("GET")
	api.HandleFunc("/invoices/order/{orderId:[0-9]+}", controllers.GetInvoiceByOrderID).Methods("GET")
	api.HandleFunc("/invoices", controllers.CreateInvoice).Methods("POST")
	api.HandleFunc("/invoices/{id:[0-9]+}/payment-status", controllers.UpdateInvoicePaymentStatus).Methods("PUT")
	// api.HandleFunc("/invoices/{id:[0-9]+}/print", controllers.GetPrintableInvoice).Methods("GET")

	// // Dashboard and Reports routes
	// api.HandleFunc("/dashboard/stats", controllers.GetDashboardStats).Methods("GET")
	// api.HandleFunc("/reports/sales", controllers.GetSalesReport).Methods("GET")

	return r
}
