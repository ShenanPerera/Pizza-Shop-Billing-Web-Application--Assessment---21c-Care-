package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"main/models"
	"main/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
	log.Println("Database connection set in controllers package")
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/items called")

	var items []models.Item
	if err := db.Where("is_active = ?", true).Find(&items).Error; err != nil {
		log.Printf("Error fetching items: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve items",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Items retrieved successfully",
		Data:    items,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("GET /api/items/%s called", id)

	itemID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid item ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var item models.Item
	if err := db.Where("id = ? AND is_active = ?", uint(itemID), true).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Item not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching item: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve item",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Item retrieved successfully",
		Data:    item,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func GetItemsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemType := vars["type"]
	log.Printf("GET /api/items/type/%s called", itemType)

	// Validate item type
	validTypes := []string{"pizza", "topping", "beverage"}
	isValid := false
	for _, t := range validTypes {
		if strings.ToLower(itemType) == t {
			isValid = true
			break
		}
	}

	if !isValid {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid item type. Valid types: pizza, topping, beverage",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var items []models.Item
	if err := db.Where("type = ? AND is_active = ?", strings.ToLower(itemType), true).Find(&items).Error; err != nil {
		log.Printf("Error fetching items by type: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve items",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Items retrieved successfully",
		Data:    items,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /api/items called")

	// TODO: Parse request body and create item
	response := utils.APIResponse{
		Success: true,
		Message: "Item created successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("PUT /api/items/%s called", id)

	// TODO: Parse request body and update item
	response := utils.APIResponse{
		Success: true,
		Message: "Item updated successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("DELETE /api/items/%s called", id)

	// TODO: Implement soft delete (set is_active = false)
	response := utils.APIResponse{
		Success: true,
		Message: "Item deleted successfully",
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// Pizza specific endpoints
func GetPizzas(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/pizzas called")

	var pizzas []models.Pizza

	// Using GORM to fetch all active pizzas
	if err := db.Where("is_active = ?", true).Find(&pizzas).Error; err != nil {
		log.Printf("Error fetching pizzas: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve pizzas",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Pizzas retrieved successfully",
		Data:    pizzas,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// Pizza by ID endpoint
func GetPizzaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("GET /api/pizzas/%s called", id)

	pizzaID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid pizza ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var pizza models.Pizza
	if err := db.Where("id = ? AND is_active = ?", uint(pizzaID), true).First(&pizza).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Pizza not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching pizza: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve pizza",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Pizza retrieved successfully",
		Data:    pizza,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func CreatePizza(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /api/pizzas called")

	// TODO: Parse request body and create pizza
	response := utils.APIResponse{
		Success: true,
		Message: "Pizza created successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}

func UpdatePizza(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("PUT /api/pizzas/%s called", id)

	// TODO: Parse request body and update pizza
	response := utils.APIResponse{
		Success: true,
		Message: "Pizza updated successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func DeletePizza(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("DELETE /api/pizzas/%s called", id)

	// TODO: Implement soft delete (set is_active = false)
	response := utils.APIResponse{
		Success: true,
		Message: "Pizza deleted successfully",
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// Topping specific endpoints
func GetToppings(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/toppings called")

	var toppings []models.Topping

	// Using GORM to fetch all active toppings
	if err := db.Where("is_active = ?", true).Find(&toppings).Error; err != nil {
		log.Printf("Error fetching toppings: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve toppings",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Toppings retrieved successfully",
		Data:    toppings,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func CreateTopping(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /api/toppings called")

	// TODO: Parse request body and create topping
	response := utils.APIResponse{
		Success: true,
		Message: "Topping created successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}

func UpdateTopping(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("PUT /api/toppings/%s called", id)

	// TODO: Parse request body and update pizza
	response := utils.APIResponse{
		Success: true,
		Message: "Topping updated successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func DeleteTopping(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("DELETE /api/toppings/%s called", id)

	// TODO: Implement soft delete (set is_active = false)
	response := utils.APIResponse{
		Success: true,
		Message: "Pizza deleted successfully",
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// Beverage specific endpoints
func GetBeverages(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/beverages called")

	var beverages []models.Beverage

	// Using GORM to fetch all active beverages
	if err := db.Where("is_active = ?", true).Find(&beverages).Error; err != nil {
		log.Printf("Error fetching beverages: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve beverages",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Beverages retrieved successfully",
		Data:    beverages,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func GetToppingByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("GET /api/toppings/%s called", id)

	toppingID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid topping ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var topping models.Topping
	if err := db.Where("id = ? AND is_active = ?", uint(toppingID), true).First(&topping).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Topping not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching topping: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve topping",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Topping retrieved successfully",
		Data:    topping,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// Beverage by ID endpoint
func GetBeverageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("GET /api/beverages/%s called", id)

	beverageID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid beverage ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var beverage models.Beverage
	if err := db.Where("id = ? AND is_active = ?", uint(beverageID), true).First(&beverage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Beverage not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching beverage: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve beverage",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Beverage retrieved successfully",
		Data:    beverage,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func CreateBeverage(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /api/beverages called")

	// TODO: Parse request body and create beverage
	response := utils.APIResponse{
		Success: true,
		Message: "Beverage created successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}

func UpdateBeverage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("PUT /api/beverages/%s called", id)

	// TODO: Parse request body and update pizza
	response := utils.APIResponse{
		Success: true,
		Message: "Beverage updated successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func DeleteBeverage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("DELETE /api/beverages/%s called", id)

	// TODO: Implement soft delete (set is_active = false)
	response := utils.APIResponse{
		Success: true,
		Message: "Beverage deleted successfully",
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}
