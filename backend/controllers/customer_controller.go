package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"main/models"
	"main/utils"

	"github.com/gorilla/mux"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/customers called")

	var customers []models.Customer
	result := db.Find(&customers)
	if result.Error != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: result.Error.Error(),
		})
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Customers retrieved successfully",
		Data:    customers,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	var customer models.Customer
	result := db.First(&customer, id)
	if result.Error != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, utils.APIResponse{
			Success: false,
			Message: "Customer not found",
		})
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Customer retrieved successfully",
		Data:    customer,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /api/customers called")

	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	if customer.Name == "" || customer.TelNo == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.APIResponse{
			Success: false,
			Message: "Name and TelNo cannot be empty",
		})
		return
	}

	result := db.Create(&customer)
	if result.Error != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, utils.APIResponse{
			Success: false,
			Message: result.Error.Error(),
		})
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Customer created successfully",
		Data:    customer,
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("PUT /api/customers/%s called", id)

	// TODO: Parse request body and update customer
	response := utils.APIResponse{
		Success: true,
		Message: "Customer updated successfully",
		Data:    nil, // Placeholder
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("DELETE /api/customers/%s called", id)

	// TODO: Implement soft delete for customer
	response := utils.APIResponse{
		Success: true,
		Message: "Customer deleted successfully",
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}
