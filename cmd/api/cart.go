package main

import (
	"encoding/json"
	"net/http"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

func (app *application) addToCart(w http.ResponseWriter, r *http.Request) {
	var baseResp response.BaseResponse

	var data request.AddToCartRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "Invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = app.store.Cart.AddToCart(r.Context(), data)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success add to cart!"

	resp, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) CartRouter() *http.ServeMux {
	cartRouter := http.NewServeMux()

	cartRouter.HandleFunc("POST /", app.addToCart)

	// Catch-all route for undefined paths
	cartRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return cartRouter
}
