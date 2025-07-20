package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"main/models"
	"main/utils"

	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Request structures
type CreateInvoiceRequest struct {
	OrderID uint   `json:"order_id" binding:"required"`
	Notes   string `json:"notes"`
}

type UpdatePaymentStatusRequest struct {
	PaymentStatus string     `json:"payment_status" binding:"required"`
	PaymentDate   *time.Time `json:"payment_date"`
	Notes         string     `json:"notes"`
}

// Response structures
type InvoiceResponse struct {
	models.Invoice
	Order *models.Order `json:"order,omitempty"`
}

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/invoices called")

	// Parse pagination parameters
	page := 1
	limit := 10

	if pageParam := r.URL.Query().Get("page"); pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}

	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	// Parse filters
	paymentStatus := r.URL.Query().Get("payment_status")

	var invoices []models.Invoice
	var total int64

	// Build query
	query := db.Model(&models.Invoice{}).Preload("Order")

	if paymentStatus != "" {
		query = query.Where("payment_status = ?", paymentStatus)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		log.Printf("Error counting invoices: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve invoices",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&invoices).Error; err != nil {
		log.Printf("Error fetching invoices: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve invoices",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Prepare response data
	responseData := map[string]interface{}{
		"invoices": invoices,
		"pagination": map[string]interface{}{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Invoices retrieved successfully",
		Data:    responseData,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func GetInvoiceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("GET /api/invoices/%s called", id)

	invoiceID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid invoice ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var invoice models.Invoice
	if err := db.Preload("Order").First(&invoice, invoiceID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Invoice not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching invoice: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve invoice",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Invoice retrieved successfully",
		Data:    invoice,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func GetInvoiceByOrderID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	log.Printf("GET /api/invoices/order/%s called", orderId)

	orderID, err := strconv.ParseUint(orderId, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid order ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var invoice models.Invoice
	if err := db.Preload("Order").Where("order_id = ?", orderID).First(&invoice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Invoice not found for this order",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching invoice by order ID: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve invoice",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Invoice retrieved successfully",
		Data:    invoice,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /api/invoices called")

	var req CreateInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate required fields
	if req.OrderID == 0 {
		response := utils.APIResponse{
			Success: false,
			Message: "Order ID is required",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Check if order exists
	var order models.Order
	if err := db.First(&order, req.OrderID).Error; err != nil {
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
			Message: "Failed to validate order",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Check if invoice already exists for this order
	var existingInvoice models.Invoice
	if err := db.Where("order_id = ?", req.OrderID).First(&existingInvoice).Error; err == nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invoice already exists for this order",
			Data:    existingInvoice,
		}
		utils.SendJSONResponse(w, http.StatusConflict, response)
		return
	}

	// Generate invoice number
	invoiceNumber := generateInvoiceNumber()

	// Calculate amounts (assuming order has total_amount field)
	// You may need to adjust this based on your Order model structure
	subtotal := order.TotalAmount // Adjust field name as needed
	taxRate := 0.10               // 10% tax rate - make this configurable
	taxAmount := subtotal * taxRate
	totalAmount := subtotal + taxAmount

	// Create invoice
	invoice := models.Invoice{
		OrderID:        req.OrderID,
		InvoiceNumber:  invoiceNumber,
		InvoiceDate:    time.Now(),
		SubtotalAmount: subtotal,
		TaxAmount:      taxAmount,
		TotalAmount:    totalAmount,
		PaymentStatus:  "pending",
		Notes:          req.Notes,
	}

	if err := db.Create(&invoice).Error; err != nil {
		log.Printf("Error creating invoice: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to create invoice",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Load the created invoice with order details
	if err := db.Preload("Order").First(&invoice, invoice.ID).Error; err != nil {
		log.Printf("Error loading created invoice: %v", err)
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Invoice created successfully",
		Data:    invoice,
	}
	utils.SendJSONResponse(w, http.StatusCreated, response)
}

func UpdateInvoicePaymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("PUT /api/invoices/%s/payment-status called", id)

	invoiceID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid invoice ID",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	var req UpdatePaymentStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate payment status
	validStatuses := map[string]bool{
		"pending":   true,
		"paid":      true,
		"overdue":   true,
		"cancelled": true,
	}

	if !validStatuses[req.PaymentStatus] {
		response := utils.APIResponse{
			Success: false,
			Message: "Invalid payment status. Must be one of: pending, paid, overdue, cancelled",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Find the invoice
	var invoice models.Invoice
	if err := db.First(&invoice, invoiceID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := utils.APIResponse{
				Success: false,
				Message: "Invoice not found",
				Data:    nil,
			}
			utils.SendJSONResponse(w, http.StatusNotFound, response)
			return
		}

		log.Printf("Error fetching invoice: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to retrieve invoice",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Update payment status
	updates := map[string]interface{}{
		"payment_status": req.PaymentStatus,
		"updated_at":     time.Now(),
	}

	if req.PaymentDate != nil {
		updates["payment_date"] = req.PaymentDate
	} else if req.PaymentStatus == "paid" && invoice.PaymentDate == nil {
		// Auto-set payment date when marking as paid
		now := time.Now()
		updates["payment_date"] = &now
	}

	if req.Notes != "" {
		updates["notes"] = req.Notes
	}

	if err := db.Model(&invoice).Updates(updates).Error; err != nil {
		log.Printf("Error updating invoice payment status: %v", err)
		response := utils.APIResponse{
			Success: false,
			Message: "Failed to update payment status",
			Data:    nil,
		}
		utils.SendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	// Reload invoice with updated data
	if err := db.Preload("Order").First(&invoice, invoiceID).Error; err != nil {
		log.Printf("Error reloading invoice: %v", err)
	}

	response := utils.APIResponse{
		Success: true,
		Message: "Payment status updated successfully",
		Data:    invoice,
	}
	utils.SendJSONResponse(w, http.StatusOK, response)
}

// func GetPrintableInvoice(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	log.Printf("GET /api/invoices/%s/print called", id)

// 	// TODO: Generate printable invoice format (HTML/PDF)
// 	response := utils.APIResponse{
// 		Success: true,
// 		Message: "Printable invoice retrieved successfully",
// 		Data:    nil, // Placeholder - should return HTML or PDF data
// 	}
// 	utils.SendJSONResponse(w, http.StatusOK, response)
// }

// Helper function to generate invoice number
func generateInvoiceNumber() string {
	// Generate invoice number with format: INV-YYYY-XXXXXX
	now := time.Now()
	year := now.Year()

	// Get count of invoices this year for sequential numbering
	var count int64
	db.Model(&models.Invoice{}).Where("EXTRACT(YEAR FROM invoice_date) = ?", year).Count(&count)

	return fmt.Sprintf("INV-%d-%06d", year, count+1)
}
