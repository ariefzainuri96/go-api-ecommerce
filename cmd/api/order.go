package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"

	dbData "github.com/ariefzainuri96/go-api-ecommerce/internal/data"
)

func createInvoice(data request.CreateInvoiceRequest, createdInvoiceResp *response.CreatedInvoiceResp) error {
	xenditKeys := os.Getenv("XENDIT_KEYS")

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api.xendit.co/v2/invoices", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	auth := base64.StdEncoding.EncodeToString([]byte(xenditKeys + ":" + ""))
	req.Header.Set("Authorization", "Basic "+auth)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = createdInvoiceResp.Unmarshal(body)

	if err != nil {
		log.Println("Error unmarshalling created invoice response:", err.Error())
		return err
	}

	return nil
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.GetUserFromContext(r)

	if !ok {
		http.Error(w, "Unauthorized, please re login!", http.StatusUnauthorized)
		return
	}

	baseResp := response.BaseResponse{}

	// Checking request
	var data request.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "Invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = app.validator.Struct(data)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	var createdInvoice response.CreatedInvoiceResp

	err = createInvoice(data.CreateInvoiceRequest, &createdInvoice)

	if err != nil {
		baseResp.Status = http.StatusInternalServerError
		baseResp.Message = "Failed to create invoice!"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	userID := user["user_id"].(int64)

	// Insert to orders table
	createOrderData := dbData.CreateOrderStruct{
		UserID:         userID,
		ProductID:      data.ProductId,
		Quantity:       data.Quantity,
		TotalPrice:     data.Amount,
		Status:         createdInvoice.Status,
		InvoiceID:      createdInvoice.ID,
		InvoiceURL:     createdInvoice.InvoiceURL,
		InvoiceExpDate: createdInvoice.ExpiryDate,
	}

	err = app.store.Order.CreateOrder(r.Context(), createOrderData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create response
	baseResp.Message = "Success create order"
	baseResp.Status = http.StatusOK
	createdOrderResp := response.CreatedOrderResp{
		BaseResponse: baseResp,
		Data: response.CreatedOrderData{
			InvoiceUrl: createdInvoice.InvoiceURL,
			InvoiceID:  createdInvoice.ID,
			Status:     createdInvoice.Status,
			ExpiryDate: createdInvoice.ExpiryDate,
		},
	}

	createdOrderRespJson, err := createdOrderResp.Marshal()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(createdOrderRespJson)
}

func (app *application) OrderRouter() *http.ServeMux {
	orderRouter := http.NewServeMux()

	orderRouter.HandleFunc("POST /create-order", app.createOrder)

	// Catch-all route for undefined paths
	orderRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return orderRouter
}
