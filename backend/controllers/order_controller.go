package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"main/models"
	"main/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CreateOrderRequest struct {
	CustomerID uint                     `json:"customer_id" binding:"required"`
	Tax        float64                  `json:"tax"`
	Items      []CreateOrderItemRequest `json:"items" binding:"required"`
}

type CreateOrderItemRequest struct {
	ItemID   uint    `json:"item_id" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Price    float64 `json:"price" binding:"required"` // Unit price
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/orders called")

	// Get pagination parameters
	page := 1
	limit := 10

	if p := r.URL.Query().Get("page"); p != "" {
		if pageNum, err := strconv.Atoi(p); err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		if limitNum, err := strconv.Atoi(l); err == nil && limitNum > 0 && limitNum <= 100 {
			limit = limitNum
		}
	}

	offset := (page - 1) * limit

	var orders []models.Order
	var totalCount int64

	// Get total count for pagination
	if err := db.Model(&models.Order{}).Count(&totalCount).Error; err != nil {
		log.Printf("Error counting orders: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve orders",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Fetch orders with pagination, including order items and their associated items
	if err := db.Preload("OrderItems.Item").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {
		log.Printf("Error fetching orders: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve orders",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Prepare pagination info
	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))

	responseData := map[string]interface{}{
		"orders": orders,
		"pagination": map[string]interface{}{
			"current_page": page,
			"total_pages":  totalPages,
			"total_count":  totalCount,
			"limit":        limit,
		},
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Orders retrieved successfully",
		Data:    responseData,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("GET /api/orders/%s called", id)

	orderID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid order ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var order models.Order
	if err := db.Preload("OrderItems").Preload("OrderItems.Item").
		Where("id = ?", uint(orderID)).
		First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Order not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching order: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve order",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}
	response := utils.APIResponse{
		Success: true,
		Message: "Order retrieved successfully",
		Data:    order,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /api/orders called")

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate required fields
	if req.CustomerID == 0 {
		response := utils.APIResponse{
			Success: false,
			Message: "Customer ID is required",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	if len(req.Items) == 0 {
		response := utils.APIResponse{
			Success: false,
			Message: "At least one item is required",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate items and calculate total
	var totalAmount float64
	for i, item := range req.Items {
		if item.ItemID == 0 {
			response := utils.APIResponse{
				Success: false,
				Message: fmt.Sprintf("Item ID is required for item %d", i+1),
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		if item.Quantity <= 0 {
			response := utils.APIResponse{
				Success: false,
				Message: fmt.Sprintf("Quantity must be greater than 0 for item %d", i+1),
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		if item.Price < 0 {
			response := utils.APIResponse{
				Success: false,
				Message: fmt.Sprintf("Price cannot be negative for item %d", i+1),
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		totalAmount += item.Price * float64(item.Quantity)
	}

	// Add tax to total
	totalAmount += req.Tax

	// Start transaction
	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("Error starting transaction: %v", tx.Error)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to create order",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Create order using your model structure
	order := models.Order{
		CustomerID:  req.CustomerID,
		OrderDate:   time.Now(),
		TotalAmount: totalAmount,
		Tax:         req.Tax,
		OrderStatus: "pending",
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		log.Printf("Error creating order: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to create order",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("Order created with ID: %d", order.ID)

	// Create order items
	for i, item := range req.Items {
		// Verify item exists and is active (assuming you have IsActive field)
		var dbItem models.Item
		if err := tx.Where("id = ?", item.ItemID).First(&dbItem).Error; err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				response := utils.APIResponse{
					Success: false,
					Message: fmt.Sprintf("Item with ID %d not found", item.ItemID),
					Data:    nil,
				}
				utils.SendJSONResponse(w, http.StatusBadRequest, response)
				return
			}
			log.Printf("Error verifying item %d: %v", item.ItemID, err)
			response := utils.APIResponse{
				Success: false,
				Message: "Failed to create order",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusInternalServerError, response)
			return
		}

		// Calculate total price for this order item
		totalPrice := item.Price * float64(item.Quantity)

		orderItem := models.OrderItem{
			OrderID:    order.ID,
			ItemID:     item.ItemID,
			Quantity:   item.Quantity,
			TotalPrice: totalPrice, // Using TotalPrice field from your model
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			log.Printf("Error creating order item %d: %v", i+1, err)
			response := utils.APIResponse{
				Success: false,
				Message: "Failed to create order",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusInternalServerError, response)
			return
		}

		log.Printf("Order item %d created successfully", i+1)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to create order",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Fetch the created order with items for response
	var createdOrder models.Order
	if err := db.Preload("OrderItems").Preload("OrderItems.Item").First(&createdOrder, order.ID).Error; err != nil {
		log.Printf("Error fetching created order: %v", err)
		// Order was created successfully, but we couldn't fetch it for response
		response := utils.APIResponse{
			Success: true,
			Message: "Order created successfully",
			Data:    map[string]interface{}{"order_id": order.ID},
		}
		utils.SendJSONResponse(w, http.StatusCreated, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Order created successfully",
		Data:    createdOrder,
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("PUT /api/orders/%s/status called", id)

	orderID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid order ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate status values (adjust as needed for your business logic)
	validStatuses := []string{"pending", "confirmed", "preparing", "ready", "delivered", "cancelled"}
	isValidStatus := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValidStatus = true
			break
		}
	}

	if !isValidStatus {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid order status",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Update order status
	if err := db.Model(&models.Order{}).Where("id = ?", uint(orderID)).Update("order_status", req.Status).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Order not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error updating order status: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to update order status",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Fetch updated order
	var updatedOrder models.Order
	if err := db.Preload("OrderItems").Preload("OrderItems.Item").First(&updatedOrder, uint(orderID)).Error; err != nil {
		log.Printf("Error fetching updated order: %v", err)
		response := utils.APIResponse{
			Success: true,
			Message: "Order status updated successfully",
			Data:    map[string]interface{}{"order_id": orderID, "status": req.Status},
		}
		utils.SendJSONResponse(w, http.StatusOK, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Order status updated successfully",
		Data:    updatedOrder,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}
